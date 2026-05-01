# Go 中高级开发者学习路线

## 目的

本路线图的目标：以 Node.js（10+ 年）全栈开发经验为基础，系统学习 Go 语言特性与生态，最终能将 Go 用于日常开发，尤其是高性能 AI Agent 方向。

后续新 session 中，可直接告知 Claude 参考此文件来延续学习计划。

## 核心主题学习顺序

1. Go 惯用法与思维转变
   - [x] 接受 interface 返回 struct：接口由消费方（调用方）定义，不是由实现方导出。实现方只需返回具体 struct
   - [x] 显式 > 隐式：类型断言必须有 comma-ok；error 必须显式处理、逐层传递；不依赖 init() 的执行顺序
   - [x] 声明错误而非抛出：错误是值，通过返回值流动。找不到记录用 sentinel error（如 ErrNotFound），不用 (nil, nil) 或 panic
   - [x] Go 的 REST：按语义使用 HTTP 状态码，不要让错误响应返回 HTTP 200
   - [x] 包按职责拆分，不按类型拆分：消除 common / utils / base 这类万能包；包名简短且表达明确功能
   - [x] 少即是多：不写省一个参数的 convenience wrapper；一个 struct 能满足的场景不抽象 interface
   - [x] 一致性优先于灵活性：同一个 struct 的 receiver 名全文件统一；命名风格遵循 Go 社区惯例

2. Go 类型系统与语义
3. 标准库
4. 测试
5. 并发编程
6. 网络与流式通信
7. 内存与性能
