# 错误处理：从 err.Error() 到 errors.Is/As

## 嫁接功能

无新功能。重构 `user/service.go` 和 `user/handler.go` 中的错误处理，引入业务错误类型。

## 涉及项目文件 / 标准库参考

- [`internal/domain/user/service.go`](../../../internal/domain/user/service.go) — 当前直接 `return nil, fmt.Errorf("邮箱已注册")`
- [`internal/domain/user/handler.go`](../../../internal/domain/user/handler.go) — 当前 `response.Fail(c, err.Error())` 暴露出内部错误信息
- [`internal/common/response/response.go`](../../../internal/common/response/response.go) — 查看当前 Fail 实现

标准库关键片段（`errors` 包）：

```go
// errors.Is 沿错误链查找匹配的哨兵错误
func Is(err, target error) bool

// errors.As 沿错误链查找匹配的错误类型
func As(err error, target interface{}) bool

// fmt.Errorf 配合 %w 包装错误
// err := fmt.Errorf("校验失败: %w", someErr)
```

## 常见误用

1. **到处 `errors.New` 而不定义哨兵错误** — 调用方无法用 `errors.Is` 判断具体错误类型，只能靠字符串匹配
2. **无脑 `%w`** — 把实现细节泄漏到 API 层，比如把 `gorm.ErrRecordNotFound` 一路传到 HTTP 响应里
3. **`err.Error()` 比较** — 用字符串比较判断错误类型，Go 1.13 之后应该用 `errors.Is`

## 练习要求

1. 在 [`internal/domain/user/`](../../../internal/domain/user/) 下新建 `errors.go`，定义业务错误：
   - `ErrEmailAlreadyExists`（哨兵错误）
   - `ErrUserNotFound`（哨兵错误）
   - `ErrDepartmentNotFound`（哨兵错误）
   - `ErrRoleNotFound`（哨兵错误）
2. 修改 `service.go`：用定义的哨兵错误替换 `fmt.Errorf` 和 `errors.New`
3. 修改 `handler.go`：根据错误类型返回不同的 HTTP 状态码（409 冲突、404 未找到等），而不是统一返回 `response.Fail`
4. 在 `response.go` 或 handler 层实现一个错误到 HTTP 状态的映射

可能踩的坑：

- 注意区分「需要包装上下文」和「直接返回哨兵错误」的场景
- gorm 的 `ErrRecordNotFound` 要转换成自己的哨兵错误，不要泄漏 ORM 细节
