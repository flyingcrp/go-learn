package cache

import (
	"sync"
	"time"
)

// entry 内部缓存项，包含值和过期时间
type entry[V any] struct {
	value     V
	expiresAt time.Time
}

// LocalCache 泛型内存缓存，使用 sync.RWMutex 保护并发访问。
// K 必须是 comparable，V 任意类型。
type LocalCache[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]entry[V]
	done  chan struct{} // 通知后台清理 goroutine 退出
}

// NewLocalCache 创建缓存实例并启动后台过期清理。
// cleanupInterval 决定清理检查频率。
func NewLocalCache[K comparable, V any](cleanupInterval time.Duration) *LocalCache[K, V] {
	c := &LocalCache[K, V]{
		items: make(map[K]entry[V]),
		done:  make(chan struct{}),
	}
	go c.cleanup(cleanupInterval)
	return c
}

// Get 返回 key 对应的值。如果 key 不存在或已过期，返回零值和 false。
func (c *LocalCache[K, V]) Get(key K) (V, bool) {
	// TODO: 使用读锁，检查 key 存在且未过期
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.items[key]
	if ok && entry.expiresAt.After(time.Now()) {
		return entry.value, true
	}
	var zero V
	return zero, false
}

// Set 存储 key-value，ttl 后过期。
func (c *LocalCache[K, V]) Set(key K, value V, ttl time.Duration) {
	// TODO: 使用写锁，存储 entry{value, expiresAt}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = entry[V]{
		expiresAt: time.Now().Add(ttl),
		value:     value,
	}
}

// Delete 删除指定 key。
func (c *LocalCache[K, V]) Delete(key K) {
	// TODO: 使用写锁，delete(c.items, key)
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// Close 停止后台清理 goroutine。调用后不应再使用缓存。
func (c *LocalCache[K, V]) Close() {
	close(c.done)
}

// cleanup 后台清理过期项，由 NewLocalCache 启动。
func (c *LocalCache[K, V]) cleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			defer c.mu.Unlock()
			for key, item := range c.items {
				if item.expiresAt.Before(time.Now()) {
					delete(c.items, key)
				}
			}
			// TODO: 使用写锁，遍历并删除过期项
		case <-c.done:
			return
		}
	}
}
