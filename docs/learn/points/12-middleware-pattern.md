# 中间件模式：责任链与注入

## 嫁接功能

实现自定义中间件：请求限流（token bucket）、响应时间记录、panic recovery 增强。

## 涉及项目文件 / 标准库参考

- [`internal/infra/middleware/`](../../../internal/infra/middleware/) — 现有中间件目录
- [`internal/infra/middleware/trace.go`](../../../internal/infra/middleware/trace.go) — 参考现有中间件写法
- [`internal/infra/middleware/auth.go`](../../../internal/infra/middleware/auth.go) — 参考现有中间件写法

net/http 中间件签名：

```go
// 标准库风格
func(next http.Handler) http.Handler

// Gin 风格
func(c *gin.Context)
```

## 常见误用

1. **在中间件里写业务逻辑** — 中间件是横切关注点（日志、鉴权、限流），不是业务
2. **中间件顺序不敏感** — auth 必须在业务 handler 前执行；trace 必须在最外层
3. **`c.Next()` 后的代码** — 中间件的洋葱模型：Next 之前是 request 阶段，之后是 response 阶段

## 练习要求

1. 实现限流中间件 `internal/infra/middleware/ratelimit.go`：
   - 使用 `golang.org/x/time/rate` 包
   - 每个 IP 一个 rate limiter（用 map + sync.RWMutex）
   - 超限返回 429 Too Many Requests
2. 实现请求计时中间件 `internal/infra/middleware/latency.go`：
   - 记录每个请求的耗时
   - 超过阈值（如 500ms）用 slog.Warn 输出慢请求日志
3. 在路由中正确编排中间件顺序：trace → latency → ratelimit → auth → handler

可能踩的坑：

- `x/time/rate` 的 `Wait` vs `Allow` 选择
- IP 提取在反向代理后面（`X-Forwarded-For` vs `X-Real-IP`）
- 限流器 map 的内存泄漏（不活跃的 IP 不清理）
