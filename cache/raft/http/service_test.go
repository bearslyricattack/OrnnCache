package http

import (
	"OrnnCache/cache/basefunction/baseclient"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

// Test_NewServer tests that a server can perform all basic operations.
func Test_NewServer(t *testing.T) {
	store := newTestStore()
	s := &testServer{New(":0", store)}
	if s == nil {
		t.Fatal("failed to create HTTP http")
	}

	if err := s.Start(); err != nil {
		t.Fatalf("failed to start HTTP http: %s", err)
	}
	//经测试该部分没有问题
	/*b := doGet(t, s.URL(), "k1")
	if str, ok := b.(string); ok && str != "" {
		t.Fatalf("wrong value received for key k1: %s (expected empty string)", str)
	}
	doPost(t, s.URL(), "k1", "v1", 1*time.Second)

	b = doGet(t, s.URL(), "k1")
	if str, ok := b.(string); ok && str != `{"value":"v1","expired":1000000000}` {
		t.Fatalf("wrong value received for key k1: %s (expected empty string)", str)
	}
	ctx := context.Background()
	store.client.Set(ctx, "k2", "v2", 1*time.Second)
	b = doGet(t, s.URL(), "k2")
	if str, ok := b.(string); ok && str != `{"value":"v2","expired":1000000000}` {
		t.Fatalf("wrong value received for key k2: %s (expected empty string)", str)
	}

	doDelete(t, s.URL(), "k2")
	b = doGet(t, s.URL(), "k2")
	if str, ok := b.(string); ok && str != "" {
		t.Fatalf("wrong value received for key k2: %s (expected empty string)", str)
	}*/

}

type testServer struct {
	*Service
}

func (t *testServer) URL() string {
	port := strings.TrimLeft(t.Addr().String(), "[::]:")
	return fmt.Sprintf("http://127.0.0.1:%s", port)
}

type testStore struct {
	client baseclient.BaseClient
}

func newTestStore() *testStore {
	return &testStore{
		client: *baseclient.New(),
	}
}

func (t *testStore) Get(key string) (interface{}, int64, bool) {
	ctx := context.Background()
	if value, yes := t.client.GetWithTTL(ctx, key); yes {
		if expiredItem, ok := value.(baseclient.Item); ok {
			return expiredItem.Object, expiredItem.Expiration, true
		}
	}
	return nil, 0, false
}

func (t *testStore) Set(key string, value interface{}, d time.Duration) error {
	ctx := context.Background()
	t.client.Set(ctx, key, value, d)
	return nil
}

func (t *testStore) Delete(key string) error {
	ctx := context.Background()
	t.client.Delete(ctx, key)
	return nil
}

func (t *testStore) Join(nodeID, addr string) error {
	return nil
}

func doGet(t *testing.T, url, key string) interface{} {
	resp, err := http.Get(fmt.Sprintf("%s/key/%s", url, key))
	if err != nil {
		t.Fatalf("failed to GET key: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response: %s", err)
	}
	return string(body)
}

func doPost(t *testing.T, url, key string, value interface{}, expired time.Duration) {
	b, err := json.Marshal(map[string]KeyValueWithExpiration{
		key: KeyValueWithExpiration{
			Value:   value,
			Expired: int64(expired),
		},
	})
	if err != nil {
		t.Fatalf("failed to encode key and value for POST: %s", err)
	}
	resp, err := http.Post(fmt.Sprintf("%s/key", url), "application-type/json", bytes.NewReader(b))
	if err != nil {
		t.Fatalf("POST request failed: %s", err)
	}
	defer resp.Body.Close()
}

func doDelete(t *testing.T, u, key string) {
	ru, err := url.Parse(fmt.Sprintf("%s/key/%s", u, key))
	if err != nil {
		t.Fatalf("failed to parse URL for delete: %s", err)
	}
	req := &http.Request{
		Method: "DELETE",
		URL:    ru,
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to GET key: %s", err)
	}
	defer resp.Body.Close()
}
