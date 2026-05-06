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
	// TODO: 用 100 个 goroutine 同时读写，go test -race 验证
	c := NewLocalCache[int, int](time.Minute)
	defer c.Close()

	var wg sync.WaitGroup
	const numGoroutines = 100
	const numOps = 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// TODO: 交替执行 Set/Get/Delete，模拟真实并发场景
			for j := 0; j < numOps; j++ {
				switch j % 3 {
				case 0:
					c.Set(id, j, time.Minute)
				case 1:
					c.Get(id)
				default:
					c.Delete(id)
				}
			}
		}(i)
	}

	wg.Wait()
}
