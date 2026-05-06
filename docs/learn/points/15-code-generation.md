# 代码生成：go generate 与 stringer

## 嫁接功能

给项目添加枚举类型，用 stringer 生成 String() 方法。

## 涉及项目文件 / 标准库参考

- [`internal/domain/user/model.go`](../../../internal/domain/user/model.go) — 用户模型，可以加状态枚举
- 新建类型定义文件

go generate 模式：

```go
//go:generate stringer -type=Status
type Status int

const (
    StatusActive Status = iota
    StatusInactive
    StatusBanned
)
```

## 常见误用

1. **手写 String() 方法** — 枚举多了维护成本高，用 stringer 自动生成
2. **`go generate` 不检查生成文件是否过时** — CI 里加检查
3. **生成的文件提交到 Git** — 应该提交（接收方不需要装工具链）

## 练习要求

1. 安装 stringer：`go install golang.org/x/tools/cmd/stringer@latest`
2. 在 `internal/domain/user/model.go` 添加用户状态枚举：
   - `UserStatus` 类型：Active、Inactive、Banned
3. 添加 `//go:generate` 注释
4. 运行 `go generate ./internal/domain/user/`
5. 检查生成的 `userstatus_string.go` 文件
6. 在代码中使用 `user.Status.String()` 替代手工字符串转换

可能踩的坑：

- stringer 生成的文件名是小写的（如 `userstatus_string.go`）
- 需要确认 `$GOPATH/bin` 在 PATH 中
- iota 的起始值和断点（建议显式赋值）
