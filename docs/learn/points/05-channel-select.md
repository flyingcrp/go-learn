# channel 与 select：生产者-消费者模式

## 嫁接功能

给项目添加一个内存事件总线，用于模块间解耦通信（例如用户注册后发送欢迎通知）。

## 涉及项目文件 / 标准库参考

- [`internal/infra/event/`](../../../internal/infra/event/) — 事件总线实现
- [`internal/domain/user/service.go`](../../../internal/domain/user/service.go) — 注册完成后发布事件

channel 关键语义：

```go
// 无缓冲 channel — 同步，发送方阻塞直到接收方就绪
ch := make(chan Event)

// 有缓冲 channel — 异步，缓冲满之前不阻塞
ch := make(chan Event, 100)

// select 多路复用
select {
case e := <-ch:
    // 处理事件
case <-ctx.Done():
    // 取消
}
```

## 常见误用

1. **向已关闭的 channel 发送** — panic。发送方绝不能 close channel
2. **关闭 channel 的一方不确定** — 一般是发送方 close，接收方用 `v, ok := <-ch` 判断
3. **select 不加 default 或不加 ctx.Done()** — 死锁或无法取消

## 练习要求

1. 新建 `internal/infra/event/bus.go`：
   - 定义 `Event` 结构体（Type + Payload）
   - 实现带缓冲的 channel 作为事件队列
   - `Publish(ctx, event)` 方法
   - `Subscribe(ctx)` 返回 `<-chan Event`
   - 使用 select 正确处理 ctx 取消
2. 在 `user/service.go` 的 `Register` 方法中，成功后发布一个 `UserRegistered` 事件
3. 新建一个简单的消费者（比如打印日志），订阅并消费事件

可能踩的坑：

- 多个订阅者时 channel 的广播语义（一个事件只能被一个消费者拿走）
- 消费者慢于生产者时 channel 满了怎么办
