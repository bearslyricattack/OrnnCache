package store

import (
	"OrnnCache/cache/basefunction/baseclient"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb/v2"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	retainSnapshotCount = 2
	raftTimeout         = 10 * time.Second
)

type command struct {
	Op          string      `json:"op,omitempty"`
	Key         string      `json:"key,omitempty"`
	Value       interface{} `json:"value,omitempty"`
	ExpiredTime int64       `json:"expiredTime,omitempty"`
}

type Store struct {
	RaftDir  string
	RaftBind string

	m      baseclient.BaseClient
	mu     sync.Mutex
	raft   *raft.Raft
	logger *log.Logger
}

func New() *Store {
	return &Store{
		m:      *baseclient.New(),
		logger: log.New(os.Stderr, "[store] ", log.LstdFlags),
	}
}

func (s *Store) Open(enableSingle bool, localID string) error {
	// Setup Raft configuration.
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(localID)

	// Setup Raft communication.
	addr, err := net.ResolveTCPAddr("tcp", s.RaftBind)
	if err != nil {
		return err
	}
	transport, err := raft.NewTCPTransport(s.RaftBind, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshots, err := raft.NewFileSnapshotStore(s.RaftDir, retainSnapshotCount, os.Stderr)
	if err != nil {
		return fmt.Errorf("file snapshot store: %s", err)
	}
	// Create the log store and stable store.
	var logStore raft.LogStore
	var stableStore raft.StableStore
	boltDB, err := raftboltdb.New(raftboltdb.Options{
		Path: filepath.Join(s.RaftDir, "raft.db"),
	})
	if err != nil {
		return fmt.Errorf("new bbolt store: %s", err)
	}
	logStore = boltDB
	stableStore = boltDB

	// Instantiate the Raft systems.
	ra, err := raft.NewRaft(config, (*fsm)(s), logStore, stableStore, snapshots, transport)
	if err != nil {
		return fmt.Errorf("new raft: %s", err)
	}
	s.raft = ra

	if enableSingle {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		ra.BootstrapCluster(configuration)
	}

	return nil
}

// Get returns the value for the given key.
func (s *Store) Get(key string) (interface{}, int64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ctx := context.Background()
	if value, yes := s.m.GetWithTTL(ctx, key); yes {
		if expiredItem, ok := value.(baseclient.Item); ok {
			return expiredItem.Object, expiredItem.Expiration, true
		}
	}
	return nil, 0, false
}

// Set sets the value for the given key.
func (s *Store) Set(key string, value interface{}, d time.Duration) error {
	if s.raft.State() != raft.Leader {
		return fmt.Errorf("not leader")
	}

	c := &command{
		Op:          "set",
		Key:         key,
		Value:       value,
		ExpiredTime: int64(d) + time.Now().Unix(),
	}
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	f := s.raft.Apply(b, raftTimeout)
	return f.Error()
}

// Delete deletes the given key.
func (s *Store) Delete(key string) error {
	if s.raft.State() != raft.Leader {
		return fmt.Errorf("not leader")
	}

	c := &command{
		Op:  "delete",
		Key: key,
	}
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	f := s.raft.Apply(b, raftTimeout)
	return f.Error()
}

// Restore stores the key-value store to a previous state.
func (f *fsm) Restore(rc io.ReadCloser) error {
	o := baseclient.New()
	if err := json.NewDecoder(rc).Decode(&o); err != nil {
		return err
	}
	ctx := context.Background()
	// Set the state from the snapshot, no lock required according to
	// Hashicorp docs.
	f.m.Flush(ctx)
	return nil
}

func (s *Store) Join(nodeID, addr string) error {
	s.logger.Printf("received join request for remote node %s at %s", nodeID, addr)

	configFuture := s.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		s.logger.Printf("failed to get raft configuration: %v", err)
		return err
	}

	for _, srv := range configFuture.Configuration().Servers {
		// If a node already exists with either the joining node's ID or address,
		// that node may need to be removed from the config first.
		if srv.ID == raft.ServerID(nodeID) || srv.Address == raft.ServerAddress(addr) {
			// However if *both* the ID and the address are the same, then nothing -- not even
			// a join operation -- is needed.
			if srv.Address == raft.ServerAddress(addr) && srv.ID == raft.ServerID(nodeID) {
				s.logger.Printf("node %s at %s already member of cluster, ignoring join request", nodeID, addr)
				return nil
			}

			future := s.raft.RemoveServer(srv.ID, 0, 0)
			if err := future.Error(); err != nil {
				return fmt.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
			}
		}
	}

	f := s.raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if f.Error() != nil {
		return f.Error()
	}
	s.logger.Printf("node %s at %s joined successfully", nodeID, addr)
	return nil
}

type fsm Store

// Apply applies a Raft log entry to the key-value store.
func (f *fsm) Apply(l *raft.Log) interface{} {
	var c command
	if err := json.Unmarshal(l.Data, &c); err != nil {
		panic(fmt.Sprintf("failed to unmarshal command: %s", err.Error()))
	}
	ctx := context.Background()
	switch c.Op {
	case "set":
		return f.applySet(ctx, c.Key, c.Value, time.Duration(c.ExpiredTime))
	case "delete":
		return f.applyDelete(ctx, c.Key)
	default:
		panic(fmt.Sprintf("unrecognized command op: %s", c.Op))
	}
}

func (f *fsm) applySet(ctx context.Context, key string, value interface{}, d time.Duration) interface{} {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.m.Set(ctx, key, value, d)
	return nil
}

func (f *fsm) applyDelete(ctx context.Context, key string) interface{} {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.m.Delete(ctx, key)
	return nil
}

// Snapshot returns a snapshot of the key-value store.
func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	// Clone the map.
	client := baseclient.New()
	ctx := context.Background()
	keys := f.m.Keys(ctx)
	for _, key := range keys {
		if item, ok := f.m.GetWithTTL(ctx, key); ok {
			if expiredItem, ok := item.(baseclient.Item); ok {
				client.Set(ctx, key, expiredItem.Object, time.Duration(expiredItem.Expiration))
			}
		}
	}
	return &fsmSnapshot{client: *client}, nil
}

type fsmSnapshot struct {
	client baseclient.BaseClient
}

func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	err := func() error {
		// Encode data.
		b, err := json.Marshal(f.client)
		if err != nil {
			return err
		}

		// Write data to sink.
		if _, err := sink.Write(b); err != nil {
			return err
		}

		// Close the sink.
		return sink.Close()
	}()

	if err != nil {
		sink.Cancel()
	}

	return err
}

func (f *fsmSnapshot) Release() {}
