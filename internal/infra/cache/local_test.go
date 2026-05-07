package cache

import (
	"sync"
	"testing"
	"time"
)

func TestLocalCache_SetGet(t *testing.T) {
	c := NewLocalCache[string, int](10 * time.Minute)
	defer c.Close()

	c.Set("answer", 42, time.Minute)
	v, ok := c.Get("answer")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if v != 42 {
		t.Fatalf("expected 42, got %d", v)
	}
}

func TestLocalCache_GetMiss(t *testing.T) {
	c := NewLocalCache[string, int](10 * time.Minute)
	defer c.Close()

	_, ok := c.Get("nonexistent")
	if ok {
		t.Fatal("expected false for missing key")
	}
}

func TestLocalCache_Delete(t *testing.T) {
	c := NewLocalCache[string, int](10 * time.Minute)
	defer c.Close()

	c.Set("x", 1, time.Minute)
	c.Delete("x")
	_, ok := c.Get("x")
	if ok {
		t.Fatal("expected false after delete")
	}
}

func TestLocalCache_Expire(t *testing.T) {
	c := NewLocalCache[string, int](50 * time.Millisecond)
	defer c.Close()

	c.Set("short", 99, 30*time.Millisecond)
	time.Sleep(80 * time.Millisecond)
	_, ok := c.Get("short")
	if ok {
		t.Fatal("expected key to expire")
	}
}

func TestLocalCache_Concurrent(t *testing.T) {
	c := NewLocalCache[int, int](time.Minute)
	defer c.Close()

	var wg sync.WaitGroup
	const numGoroutines = 100
	const numOps = 1000

	// 使用少量共享 key，让 goroutine 之间产生真正的竞争
	sharedKeys := []int{1, 2, 3}

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				key := sharedKeys[j%len(sharedKeys)]
				switch j % 3 {
				case 0:
					c.Set(key, j, time.Minute)
				case 1:
					c.Get(key)
				default:
					c.Delete(key)
				}
			}
		}()
	}

	wg.Wait()
}
