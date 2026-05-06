# goroutine 生命周期：启动、退出、泄漏检测

## 嫁接功能

给项目添加一个后台 worker，定期清理过期 token（或定期打印服务状态）。

## 涉及项目文件 / 标准库参考

- [`cmd/server/main.go`](../../../cmd/server/main.go) — 启动 worker 的位置
- [`internal/infra/`](../../../internal/infra/) — worker 实现位置

标准库关键片段：

```go
// net/http.Server 的 ListenAndServe 模式
go func() {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        slog.Error("Listen Error: ", "error", err)
    }
}()
```

## 常见误用

1. **启动 goroutine 不管退出** — goroutine 泄漏是最常见的内存泄漏，每个 goroutine 至少占 2KB
2. **用 `time.Sleep` 做定时器** — 无法被取消，优雅关闭时等它自然醒
3. **没有 recover** — 一个 goroutine panic 会导致整个进程崩溃

## 练习要求

1. 新建 `internal/infra/worker/cleanup.go`，实现一个定时 worker：
   - 接受 `ctx context.Context` 控制生命周期
   - 使用 `time.Ticker`（不是 `time.Sleep`）
   - 在收到 ctx.Done() 时退出
   - 用 `defer` + `recover` 防止 panic 扩散
2. 在 `cmd/server/main.go` 中启动这个 worker，在优雅关闭时让它退出
3. 运行 `go run -race cmd/server/main.go`，确保没有 data race

可能踩的坑：

- Ticker 的 Stop 方法和 channel 的 drain
- 多个 goroutine 的启动顺序和退出顺序
- context 取消后 Ticker 不会自动停
