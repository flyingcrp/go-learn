# 优雅关闭：组件注册与依赖顺序

## 嫁接功能

重构 `cmd/server/main.go` 的启动和关闭逻辑，将组件注册到统一的生命周期管理器。

## 涉及项目文件 / 标准库参考

- [`cmd/server/main.go`](../../../cmd/server/main.go) — 当前主函数，手动管理各组件的启动关闭
- [`internal/infra/storage/infra.go`](../../../internal/infra/storage/infra.go) — 基础设施

os/signal 关键模式：

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()
```

## 常见误用

1. **组件启动顺序写死在 main 里** — 组件多了之后 main 函数变成意大利面
2. **关闭顺序随便** — 应该先关依赖方再关被依赖方（HTTP server → DB → 其他）
3. **关闭超时不合理** — 生产环境可能需要更长；开发环境可以短但至少给 inflight 请求一个窗口

## 练习要求

1. 新建 `internal/infra/lifecycle/lifecycle.go`：
   - 定义 `Component` 接口：`Name() string`、`Start(ctx) error`、`Stop(ctx) error`
   - 实现 `Manager`，管理多个 Component：注册、启动、优雅关闭
   - 启动按注册顺序，关闭按注册逆序
   - 每个组件的 Stop 有独立超时
2. 把 HTTP server、MySQL 连接、worker（来自阶段 2）改造为 Component
3. 重构 `cmd/server/main.go`，使用 Manager 管理所有组件
4. main 函数不应超过 30 行

可能踩的坑：

- Start 方法是否需要阻塞（HTTP server 的 ListenAndServe 是阻塞的）
- 某个组件 Start 失败，已启动的组件需要回滚
