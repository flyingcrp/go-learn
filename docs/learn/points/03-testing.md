# 测试基础：表驱动测试与 mock

## 嫁接功能

为 `user/service.go` 编写单元测试。

## 涉及项目文件 / 标准库参考

- [`internal/domain/user/service.go`](../../../internal/domain/user/service.go) — 被测试对象
- `internal/domain/user/service_test.go` — 新建测试文件（你需要创建）

Go 标准库测试模式：

```go
func TestRegister(t *testing.T) {
    tests := []struct {
        name    string
        input   *UserRegisterReq
        wantErr bool
    }{
        {name: "正常注册", input: &UserRegisterReq{...}, wantErr: false},
        {name: "邮箱已存在", input: &UserRegisterReq{...}, wantErr: true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // 调用被测函数
        })
    }
}
```

## 常见误用

1. **在测试里写复杂逻辑** — 测试代码不需要 DRY，重复就是可读性
2. **mock 数据库而不是 mock 接口** — 不需要真实 MySQL，mock 你定义的接口即可
3. **只测 happy path** — 边界条件（邮箱重复、部门不存在）才是 bug 高发区

## 练习要求

1. 新建 `internal/domain/user/service_test.go`
2. 为 `UserRepo`、`departmentChecker`、`roleChecker` 三个接口各写一个 mock struct（手动 mock，不用 mockgen）
3. 为 `Register` 方法写表驱动测试，至少覆盖：
   - 正常注册成功
   - 邮箱已存在
   - 部门校验失败
   - 角色校验失败
4. 运行 `go test -v ./internal/domain/user/` 通过

可能踩的坑：

- `uuid.NewV7()` 会产生随机值，测试里需要处理这个非确定性
- errgroup 的 goroutine 调度顺序不确定，mock 期望要对应
