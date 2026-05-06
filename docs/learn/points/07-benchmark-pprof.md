# benchmark 与 pprof 入门

## 嫁接功能

无新功能。为现有代码写 benchmark，用 pprof 分析。

## 涉及项目文件 / 标准库参考

- [`internal/domain/user/service_test.go`](../../../internal/domain/user/service_test.go) — benchmark 函数
- [`cmd/server/main.go`](../../../cmd/server/main.go) — 需要引入 `net/http/pprof`

benchmark 基本格式：

```go
// 函数名以 Benchmark 开头，参数 *testing.B
func BenchmarkRegister(b *testing.B) {
    // 准备数据
    for i := 0; i < b.N; i++ {
        // 被测代码
    }
}
```

pprof 接入：

```go
import _ "net/http/pprof"
// 启动 pprof HTTP 服务
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

## 常见误用

1. **benchmark 里的 b.N 循环包含准备代码** — 准备数据应该在循环外
2. **只看 wall time 不看 allocs** — `go test -benchmem` 查看内存分配
3. **pprof 只在线上开** — 应该是默认启用的，只是换个端口

## 练习要求

1. 为 `user/service.go` 的 `Register` 和 `List` 方法写 benchmark
2. 使用 `go test -bench=. -benchmem ./internal/domain/user/` 运行
3. 阅读输出，找到分配最多的代码路径
4. 在 `cmd/server/main.go` 中添加 pprof 端点
5. 用 `go tool pprof` 采样 heap profile

可能踩的坑：

- mock 的时间如果包含在 benchmark 里，结果没有意义
- 启动 pprof 后要用 HTTP 请求触发，空闲进程没有数据
