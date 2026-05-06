# 数据库查询优化：N+1 与批量操作

## 嫁接功能

无新功能。分析并优化现有数据库查询模式，添加批量查询接口。

## 涉及项目文件 / 标准库参考

- [`internal/domain/user/repo.go`](../../../internal/domain/user/repo.go) — 当前查询实现
- [`internal/domain/user/service.go`](../../../internal/domain/user/service.go) — `Register` 中的两次并发校验

GORM 关键模式：

```go
// Preload 避免 N+1
db.Preload("Department").Find(&users)

// 批量插入
db.CreateInBatches(users, 100)

// SELECT 特定列
db.Select("id", "name", "email").Find(&users)
```

## 常见误用

1. **循环里查数据库** — 典型的 N+1，比如查用户列表后逐个查部门
2. **不做分页** — `Find(&users)` 不加 Limit，全表扫
3. **`Omit("Token")` 但 Token 是大字段** — 用 `Select` 明确要哪些列比 Omit 更安全

## 练习要求

1. 检查 `user/repo.go` 的 `List` 方法：当前只 Omit 了 Token，分析是否有不必要的列被查出来
2. 给 `user/repo.go` 添加 `ListByIDs(ctx, ids []string)` 批量查询方法
3. 在 `service.go` 中考虑：如果 `Register` 里的 `depChecker.CheckID` 和 `roleChecker.CheckID` 各查一次 DB，加上 `ExistsByEmail`，一次注册有 3 次 DB 查询。如果有批量注册的需求，如何优化？
4. 实现 `BatchCreate(ctx, users []User) error`

可能踩的坑：

- GORM 的 `Updates` 对零值字段的处理（`UpdateColumns` vs `Updates`）
- 批量查询 IN 子句的元素数量限制
