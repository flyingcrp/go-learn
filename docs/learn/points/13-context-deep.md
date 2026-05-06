# context 深入：取消传播与超时策略

## 嫁接功能

审查项目中所有 context 使用，统一超时策略，修复潜在问题。

## 涉及项目文件 / 标准库参考

- [`internal/domain/user/service.go`](../../../internal/domain/user/service.go) — `errgroup.WithContext`、`context.WithTimeout`
- [`internal/domain/user/repo.go`](../../../internal/domain/user/repo.go) — `db.WithContext(ctx)`

context 包关键：

```go
// 派生 context
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()

// 不要存 context 在结构体里
// Bad:
type Service struct { ctx context.Context }
// Good:
func (s *Service) Do(ctx context.Context, ...) {}
```

## 常见误用

1. **context 存在结构体字段里** — context 应该通过参数传递，不该被存储
2. **`defer cancel()` 之后还用 ctx** — cancel 之后 ctx 就失效了
3. **超时设太长或太短** — 没有考虑上游的超时 / 用户的等待耐心
4. **`context.Background()` 在请求处理路径里** — 应该用请求的 ctx

## 练习要求

1. 审查项目中所有使用 context 的地方，找出问题
2. 修复 `user/service.go` 中 `Register` 的 context 使用：
   - `errgroup.WithContext` 的返回值 `egCtx` 被 `context.WithTimeout` shadow 了，分析是否需要单独命名
3. 为 repo 操作设置合理的 DB 超时（用 `context.WithTimeout` 包裹）
4. 实现一个 `middleware.Timeout(duration)` 中间件，给每个请求设置整体超时

可能踩的坑：

- HTTP handler 里 gin.Context 和 request.Context() 的区别
- context 取消后 gorm 操作会返回什么错误
