package store

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func Test_StoreOpen(t *testing.T) {
	s := New()
	tmpDir, _ := ioutil.TempDir("", "store_test")
	defer os.RemoveAll(tmpDir)

	s.RaftBind = "127.0.0.1:0"
	s.RaftDir = tmpDir
	if s == nil {
		t.Fatalf("failed to create store")
	}

	if err := s.Open(false, "node0"); err != nil {
		t.Fatalf("failed to open store: %s", err)
	}
}

func Test_StoreOpenSingleNode(t *testing.T) {
	s := New()
	tmpDir, _ := ioutil.TempDir("", "store_test")
	defer os.RemoveAll(tmpDir)

	s.RaftBind = "127.0.0.1:0"
	s.RaftDir = tmpDir
	if s == nil {
		t.Fatalf("failed to create store")
	}

	if err := s.Open(true, "node0"); err != nil {
		t.Fatalf("failed to open store: %s", err)
	}

	// Simple way to ensure there is a leader.
	time.Sleep(3 * time.Second)

	if err := s.Set("foo", "bar", 1*time.Second); err != nil {
		t.Fatalf("failed to set key: %s", err.Error())
	}

	// Wait for committed log entry to be applied.
	time.Sleep(500 * time.Millisecond)
	value, expired, ok := s.Get("foo")
	if !ok {
		t.Fatalf("failed to get key: 'foo'")
	}
	if value != "bar" {
		t.Fatalf("key has wrong value: %s", value)
	}
	if expired == 0 && time.Now().UnixNano() > expired {
		t.Fatalf("key has expired,please delete it")
	}
	if err := s.Delete("foo"); err != nil {
		t.Fatalf("failed to delete key: %s", err.Error())
	}

	// Wait for committed log entry to be applied.
	time.Sleep(500 * time.Millisecond)
	//time.Sleep(300 * time.Millisecond)
	value, expired, ok = s.Get("foo")

	assert.Equal(t, ok, false, "key should not exist")
	assert.Nil(t, value, "key should not exist")
	assert.Equal(t, expired, int64(0))
}
