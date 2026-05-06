# 事务管理：数据库事务的惯用模式

## 嫁接功能

给用户注册添加事务支持，确保用户创建和关联数据写入的原子性。

## 涉及项目文件 / 标准库参考

- [`internal/domain/user/repo.go`](../../../internal/domain/user/repo.go) — 当前没有事务支持
- [`internal/infra/storage/infra.go`](../../../internal/infra/storage/infra.go) — DB 实例

GORM 事务模式：

```go
// 手动事务
tx := db.Begin()
// ...
tx.Commit()
tx.Rollback()

// 闭包事务（推荐）
db.Transaction(func(tx *gorm.DB) error {
    // 返回 error 自动回滚
    return nil
})
```

## 常见误用

1. **忘了 `defer tx.Rollback()`** — Commit 之后 Rollback 是 no-op，但 Commit 失败后需要手动 Rollback
2. **长事务** — 事务里做 RPC 调用或 HTTP 请求
3. **事务传播用 ctx 代替** — ctx 不传递事务信息，需要显式传 `*gorm.DB`

## 练习要求

1. 修改 `UserRepo` 接口，添加事务相关方法：
   - `WithTx(tx *gorm.DB) UserRepo` — 返回使用事务的 repo
2. 在 `user/service.go` 的 `Register` 中添加事务支持：
   - 创建用户应该在事务中执行
   - 如果后续操作失败，用户创建也应该回滚
3. 实现时用 GORM 的 `Transaction` 闭包模式
4. 考虑：事务的边界应该在哪一层？service 还是 repo？

可能踩的坑：

- GORM Transaction 闭包里的 panic 会触发 rollback
- `db.Transaction` 返回的 error 如果是 `gorm.ErrRecordNotFound`，事务会自动回滚
- 事务隔离级别默认是 REPEATABLE READ
