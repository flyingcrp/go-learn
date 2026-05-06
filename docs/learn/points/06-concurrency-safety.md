# 并发安全：data race 与 sync 原语

## 嫁接功能

给项目添加一个内存缓存层（例如缓存部门信息，减少 DB 查询），并保证并发安全。

## 涉及项目文件 / 标准库参考

- [`internal/infra/cache/`](../../../internal/infra/cache/) — 内存缓存实现
- [`internal/domain/user/service.go`](../../../internal/domain/user/service.go) — 使用缓存

sync 包关键类型：

```go
// sync.Mutex — 互斥锁
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()

// sync.RWMutex — 读写锁，读多写少场景
var mu sync.RWMutex
mu.RLock()  // 多个 goroutine 可同时持有读锁
mu.RUnlock()

// sync.Map — 并发安全 map，适合读多写少
var m sync.Map
m.Store("key", value)
v, ok := m.Load("key")
```

## 常见误用

1. **复制含锁的结构体** — `sync.Mutex` 值复制后锁失效，go vet 能检测
2. **map 并发写** — Go 的 map 不是并发安全的，并发写会 panic
3. **读多写少场景用 Mutex 而不是 RWMutex** — 白白降低并发度

## 练习要求

1. 新建 `internal/infra/cache/local.go`：
   - 实现一个泛型 `LocalCache[K comparable, V any]`
   - 使用 `sync.RWMutex` 保护内部 map
   - `Get(key K) (V, bool)` — 读锁
   - `Set(key K, value V)` — 写锁
   - `Delete(key K)` — 写锁
2. 添加带超时的缓存项：存值时记录 `expiresAt`，`Get` 时检查过期
3. 写并发测试（起 100 个 goroutine 同时读写），用 `go test -race` 验证

可能踩的坑：

- 锁的粒度：不要把耗时操作放在锁内
- `sync.RWMutex` 的写锁饥饿问题
- 泛型语法（Go 1.18+）
