# Go 学习路线
>
> 更新时间：2026-05-08

## 阶段 1：惯用 Go — 从能跑到写对

- [x] 1. [错误处理：从 err.Error() 到 errors.Is/As](points/01-errors.md)
- [x] 2. [接口设计：消费者侧接口与依赖反转](points/02-interfaces.md)
- [x] 3. [测试基础：表驱动测试与 mock](points/03-testing.md)

## 阶段 2：并发编程 — 从串行到并行

- [x] 4. [goroutine 生命周期：启动、退出、泄漏检测](points/04-goroutine-lifecycle.md)
- [x] 5. [channel 与 select：生产者-消费者模式](points/05-channel-select.md)
- [x] 6. [并发安全：data race 与 sync 原语](points/06-concurrency-safety.md)

## 阶段 3：性能工程 — 从够快到更快

- [x] 7. [benchmark 与 pprof 入门](points/07-benchmark-pprof.md)
- [ ] 8. [内存分配优化与逃逸分析](points/08-allocation-escape.md)
- [ ] 9. [数据库查询优化：N+1 与批量操作](points/09-db-optimization.md)

## 阶段 4：生产就绪 — 从 demo 到服务

- [ ] 10. [结构化日志与上下文传播](points/10-structured-logging.md)
- [ ] 11. [优雅关闭：组件注册与依赖顺序](points/11-graceful-shutdown.md)
- [ ] 12. [中间件模式：责任链与注入](points/12-middleware-pattern.md)

## 阶段 5：进阶专题

- [ ] 13. [context 深入：取消传播与超时策略](points/13-context-deep.md)
- [ ] 14. [事务管理：数据库事务的惯用模式](points/14-transactions.md)
- [ ] 15. [代码生成：go generate 与 stringer](points/15-code-generation.md)
