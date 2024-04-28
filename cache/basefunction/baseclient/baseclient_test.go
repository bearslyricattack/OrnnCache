package baseclient

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestItem_Expired(t *testing.T) {
	ctx := context.Background()
	client := New()
	client.Set(ctx, "key", "value", 1*time.Second)
	time.Sleep(1 * time.Second)
	if value, ok := client.Get(ctx, "key"); ok {
		fmt.Println(value)
		assert.Error(t, fmt.Errorf("get wrong value"))
	}
	assert.Equal(t, client.ItemCount(ctx), 0)
	client.Set(ctx, "key1", "value1", 1*time.Second)
	time.Sleep(700 * time.Millisecond)
	if value, ok := client.Get(ctx, "key1"); ok {
		assert.Equal(t, value, "value1")
	}
	client.Set(ctx, "key2", "value2", 1*time.Second)
	time.Sleep(1 * time.Second)
	if value, ok := client.GetWithTTL(ctx, "key2"); ok {
		if expiredItem, ok := value.(Item); ok {
			assert.Equal(t, ok, expiredItem.Expired())
		}
	}
}

func TestItem_ExpiredWithTTL(t *testing.T) {
	ctx := context.Background()
	client := New()
	client.Set(ctx, "key", "value", 1*time.Second)
	//time.Sleep(1 * time.Second)
	time.Sleep(700 * time.Millisecond)
	if value, ok := client.GetWithTTL(ctx, "key"); ok {
		if expiredItem, ok := value.(Item); ok {
			assert.Equal(t, false, time.Now().UnixNano() > expiredItem.Expiration)
		}
	}
}
