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
   - [x] 值语义 vs 指针语义：Go 一切传值，slice/map 的「引用感」只是 header 拷贝
   - [x] Struct 嵌入（组合 > 继承）：方法提升、字段提升，不是继承而是语法糖
   - [x] Slice 内部原理：ptr + len + cap 三元组，append 扩容与底层数组共享的坑
   - [x] Interface 底层表示：eface vs iface，nil interface ≠ nil concrete type
   - [x] 零值设计：一切都有零值，利用零值减少 nil 检查
   - [x] 类型定义 vs 类型别名：type T S vs type T = S
   - [x] 泛型基础：类型参数、约束（constraints）、~ 近似约束
3. 标准库

   **基础 — 写 Go 绕不开的**
   - [x] `fmt`：格式化输入输出，Stringer 接口
   - [x] `io` / `bufio`：Reader/Writer 是 I/O 基石；缓冲读写
   - [x] `strings` / `bytes`：Builder 避免拼接分配；bytes 是 []byte 版的 strings
   - [x] `strconv`：字符串与数值互转
   - [x] `os` / `path/filepath`：文件操作、环境变量；跨平台路径
   - [x] `time`：Time 值类型、Duration / Ticker / Timer、时区
   - [x] `errors`：wrap / unwrap、Is / As、Join（1.20+）

   **中级 — 项目实战常用**
   - [x] `context`：取消传播、超时控制，请求生命周期的脊梁
   - [x] `encoding/json`：struct tag、自定义 Marshal/Unmarshal、RawMessage
   - [x] `sync`：Mutex / RWMutex、WaitGroup、Once、Pool
   - [x] `flag`：CLI 参数解析
   - [ ] `net/http` / `net/url`：Server / Handler / 中间件模式；Client 连接复用
   - [ ] `log/slog`：结构化日志（1.21+，替代 log 包）

   **高级 — 性能 & 细节**
   - [ ] `sync/atomic`：Go 1.19+ 推荐 atomic.Int64 类型
   - [ ] `slices` / `sort`：泛型切片操作（1.21+，逐步替代 sort）
   - [ ] `net`：TCP / UDP 底层（后续并发编程会深入）
   - [ ] `regexp`：线性时间保证，不支持回溯/前瞻
4. 测试
5. 并发编程
6. 网络与流式通信
7. 内存与性能
