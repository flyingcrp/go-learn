# 内存分配优化与逃逸分析

## 嫁接功能

无新功能。分析现有代码的内存分配，对比优化前后。

## 涉及项目文件 / 标准库参考

- [`internal/domain/user/service.go`](../../../internal/domain/user/service.go) — `toUserDetailResp`、`toLoginResp`
- [`internal/domain/user/dto.go`](../../../internal/domain/user/dto.go) — DTO 转换函数
- [`internal/common/response/response.go`](../../../internal/common/response/response.go) — 响应构造

逃逸分析命令：

```bash
go build -gcflags="-m" ./...
# 关注 "escapes to heap" 和 "does not escape"
```

sync.Pool 模式：

```go
var bufPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}
buf := bufPool.Get().(*bytes.Buffer)
defer bufPool.Put(buf)
```

## 常见误用

1. **返回函数内局部变量的指针** — 会导致逃逸到堆，但有时候这是必要的（比如返回结构体指针）
2. **盲目用 sync.Pool** — 只有高频分配 + 短生命周期的对象才值得
3. **在 interface{} 里装值类型** — 会装箱逃逸

## 练习要求

1. 对项目运行逃逸分析，记录到 heap 的分配位置
2. 挑 2-3 个可以优化的点（例如 DTO 转换中的 `make` 预分配）
3. 写 benchmark 对比优化前后的分配次数
4. 判断是否值得引入 `sync.Pool` 到某个热点路径

可能踩的坑：

- 逃逸分析的结果依赖编译器版本
- `sync.Pool` 里的对象随时会被 GC 清掉，不能假设它一直存在
