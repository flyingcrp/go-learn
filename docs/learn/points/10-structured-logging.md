# 结构化日志与上下文传播

## 嫁接功能

增强现有日志系统，实现从 context 中自动提取 traceID 等字段，避免每次手动传参。

## 涉及项目文件 / 标准库参考

- [`internal/infra/logger/logger.go`](../../../internal/infra/logger/logger.go) — 当前日志实现
- [`internal/infra/middleware/trace.go`](../../../internal/infra/middleware/trace.go) — trace 中间件
- [`internal/domain/user/handler.go`](../../../internal/domain/user/handler.go) — `slog.WarnContext(c.Request.Context(), "测试")`

log/slog 关键 API：

```go
// 从 context 提取属性
func (h *Handler) Handle(ctx context.Context, r Record) error

// 自定义 Handler 包装
type ContextHandler struct {
    slog.Handler
}
```

## 常见误用

1. **日志和参数分离** — `slog.Info("用户注册:" + email)` 是字符串拼接，不是结构化
2. **关键信息不在 context 里** — traceID、userID 存在 context 里，但日志没有自动提取
3. **无脑 `slog.ErrorContext` 但没检查 context 里有什么** — 需要自定义 Handler

## 练习要求

1. 修改 `internal/infra/logger/logger.go`：
   - 实现自定义 `slog.Handler`，从 context 中自动提取 traceID（已由中间件注入）
   - 让日志自动带上 traceID，而不是每次手动拼接
   - 同时支持默认字段（如 service name、version）
2. 修改 `internal/infra/middleware/trace.go`：在 context 中注入 traceID 的同时，也注入 userID（从 JWT 中解析）
3. 验证：发起一个请求，观察日志输出是否自动带上了 traceID

可能踩的坑：

- `slog.Handler` 接口的方法签名（Handle、Enabled、WithAttrs、WithGroup）
- 需要正确处理 `r.With()` 和属性合并
