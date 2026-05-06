# 接口设计：消费者侧接口与依赖反转

## 嫁接功能

无新功能。审视并重构现有接口定义，理解「谁用谁定义」原则。

## 涉及项目文件 / 标准库参考

- [`internal/domain/user/service.go`](../../../internal/domain/user/service.go) — `UserRepo`、`departmentChecker`、`roleChecker` 接口
- [`internal/domain/role/service.go`](../../../internal/domain/role/service.go) — 角色服务接口
- [`internal/domain/department/service.go`](../../../internal/domain/department/service.go) — 部门服务接口

Go 标准库中的经典小接口：

```go
// io.Reader — 只有一个方法
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

## 常见误用

1. **接口定义在生产方** — Java/C# 习惯：定义 `UserRepository interface` 然后 `UserRepositoryImpl` 实现。Go 里接口定义在使用方（消费者侧）
2. **大而全的接口** — 一个接口包含所有 CRUD 方法，mock 时必须实现全部
3. **提前抽象** — 只有一个实现时就定义接口，增加不必要的间接层

## 练习要求

1. 审视 `internal/domain/user/service.go` 中 `UserRepo` 接口是否太大
2. 按方法职责拆分为小接口：`userCreator`、`userFinder`、`userUpdater` 等（或更好的命名）
3. 观察 `departmentChecker` 和 `roleChecker` 接口 — 它们已经符合「消费者侧定义」原则，思考为什么这样设计
4. 修改 `user/service.go` 的 `NewUserService` 构造函数，接受小接口而非大接口

可能踩的坑：

- 拆分后需要改 `user/module.go` 的组装代码
- 多个小接口指向同一个实现（`gormUserRepo`）是正常的，Go 里这叫「满足多个接口」
