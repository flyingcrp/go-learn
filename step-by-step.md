# Go 快速学习路线：从 Node.js 到 Go 实战

目标：不是“把 Go 语法背完”，而是尽快能维护和开发一个真实的 Gin/GORM 后端项目，并能在面试中讲清楚自己的设计。

你已经有 Node.js 开发经验，所以重点放在 Go 和 Node.js 不一样的地方：类型系统、错误处理、接口、上下文、并发、测试、工程组织。

---

## 总体判断

进度：进行中

按这份路线完成后，目标不是“精通 Go”，而是达到可以进入普通 Go Web 后端项目的水平：

- [ ] 能读懂并修改 Gin/GORM 分层项目。
- [ ] 能独立开发 CRUD、分页、鉴权、配置、日志、测试。
- [ ] 能处理常见后端问题：参数校验、错误响应、事务、密码安全、请求超时。
- [ ] 能在面试里讲清楚 handler/service/repo、context、interface、测试和并发的取舍。

仍然要诚实一点：真实工作里还会遇到缓存、消息队列、微服务、部署、监控、CI/CD。这些可以后续补，但不是你当前从 Node.js 转 Go 找第一份 Go 岗位的第一优先级。

---

## 阶段一：读懂并改动当前项目

进度：进行中

目标：先能看懂项目，能做小功能，能定位 bug。

当前状态：

- [x] `User` 补充 `CreatedAt`、`UpdatedAt`。
- [x] 修复 `department/model.go` 的 GORM tag。
- [x] `context.Context` 已从 handler 透传到 service/repo。
- [x] user、department、role 的 repo 查询已接入 `db.WithContext(ctx)`。
- [ ] 下一步：新增用户详情接口。

### 1. Go 基础语法迁移

进度：进行中

重点学习：

- [ ] `package`、导入、导出规则：大写开头才对包外可见。
- [ ] `struct`、方法、指针接收者。
- [ ] `error` 返回值：Go 常用显式错误处理，不像 Node.js 常用异常或 Promise reject。
- [ ] `slice`、`map`、`for range`。
- [ ] `defer`：资源释放、关闭连接、延迟执行。

项目练习：

- [ ] 通读 `internal/user`、`internal/department`、`internal/role` 三个模块。
- [x] 给 `User` 补 `CreatedAt`、`UpdatedAt` 字段。
- [x] 修复 `department/model.go` 里的 GORM tag：`,` 应为 `;`。
- [ ] 统一 department/role repo 中的 `gorm.ErrRecordNotFound` 判断，使用 `errors.Is`。

### 2. 理解三层结构

进度：进行中

项目规则：

- [ ] `handler` 只做 HTTP 收发：参数绑定、状态码、响应。
- [ ] `service` 只放业务逻辑。
- [ ] `repo` 只做数据访问。
- [ ] 能复用就复用 `internal/common/*`。

项目练习：

- [ ] 给用户模块加一个详情接口：`GET /v1/user/:id`。
- [ ] handler 只取 `id` 和返回响应。
- [ ] service 判断用户是否存在。
- [ ] repo 查询数据库。
- [ ] 复用 `internal/common/response` 返回统一响应。

### 3. context.Context 链路

进度：已完成

为什么重要：

Go 后端里，`context.Context` 用来传递请求取消、超时、trace 信息。它很像 Node.js 里的 request scope，但 Go 会显式传参。

项目练习：

- [x] `handler` 从 `c.Request.Context()` 获取 ctx。
- [x] `service` 方法第一参数加 `ctx context.Context`。
- [x] `repo` 使用 `db.WithContext(ctx)`。
- [x] 已推广到 user、department、role 模块。
- [ ] 后续学习 `context.WithTimeout` 和 `context cancellation` 在并发里的用法。

---

## 阶段二：具备真实接口开发能力

进度：未开始

目标：能像入职项目一样开发、调试、交付接口。

### 4. 配置管理

进度：未开始

当前问题：

- [ ] `cmd/server/main.go` 里直接读 `os.Getenv`。
- [ ] 环境变量缺少默认值。
- [ ] 环境变量缺少启动时校验。
- [ ] JWT secret、server addr、shutdown timeout 还没有统一配置入口。

项目练习：

- [ ] 新建 `internal/common/config/config.go`。
- [ ] 定义 `Config` struct。
- [ ] 启动时一次性加载并校验环境变量。
- [ ] 把 DB、JWT secret、server addr、shutdown timeout 都放进去。
- [ ] main 函数只负责组装依赖，不到处读环境变量。

### 5. JWT 中间件

进度：未开始

当前问题：

- [ ] `AuthGuard` 写了但没挂到路由。
- [ ] JWT secret 硬编码。
- [ ] `strings.Split(...)[1]` 在 Authorization 为空或格式错误时会 panic。
- [ ] 鉴权成功后还没有把用户信息写入请求上下文。

项目练习：

- [ ] 修复 Authorization 解析。
- [ ] 从配置读取 JWT secret。
- [ ] 校验 token 签名、过期时间、签名算法。
- [ ] 给需要鉴权的路由组挂上中间件。
- [ ] 中间件解析成功后用 `c.Set` 写入用户信息。
- [ ] handler/service 不直接解析 JWT，只读取中间件产物。

### 6. 密码哈希与登录注册

进度：未开始

为什么提前：

用户系统里，明文密码是高风险问题。这个比很多高级语法都更接近真实工作。

项目练习：

- [ ] 注册请求增加 password。
- [ ] User model 增加 password hash 字段，不存明文密码。
- [ ] 使用 `bcrypt.GenerateFromPassword` 存储密码哈希。
- [ ] 增加登录接口。
- [ ] 登录时使用 `bcrypt.CompareHashAndPassword` 校验。
- [ ] 登录成功后签发 JWT。
- [ ] 不要在响应里返回密码、password hash、JWT secret。

### 7. 统一错误和响应

进度：未开始

为什么重要：

真实项目里，不能让 service 返回的任意字符串直接决定 HTTP 行为。你需要能区分参数错误、未登录、无权限、资源不存在、业务冲突、服务器错误。

项目练习：

- [ ] 设计常见业务错误，例如 `ErrNotFound`、`ErrConflict`、`ErrUnauthorized`。
- [ ] service 返回业务错误，不直接关心 HTTP 状态码。
- [ ] handler 把业务错误转换成统一响应。
- [ ] 复用或完善 `internal/common/response`。
- [ ] 错误响应不泄露数据库细节。

### 8. 接口调试与文档

进度：未开始

为什么重要：

入职后你不只写代码，还要让别人能调用、能复现、能联调。

项目练习：

- [ ] 在 README 里写清楚启动方式。
- [ ] 写清楚必要环境变量。
- [ ] 记录用户、部门、角色接口的请求/响应示例。
- [ ] 使用 curl、Postman 或 Apifox 验证主要接口。
- [ ] 给每个新接口保留最小可复现调用示例。

---

## 阶段三：测试、数据库和工程质量

进度：未开始

目标：能写出可维护、可验证的 Go 后端代码。

### 9. 单元测试 + interface

进度：未开始

学习重点：

- [ ] `_test.go`。
- [ ] table-driven tests。
- [ ] `testing.T`。
- [ ] 为了测试而抽接口，不要为了抽象而抽象。

项目练习：

- [ ] 给 `UserService.Register` 写测试。
- [ ] 给 `UserService` 的 repo 依赖抽一个小接口。
- [ ] 用 fake repo 测试邮箱已存在。
- [ ] 用 fake repo 测试部门不存在。
- [ ] 用 fake repo 测试角色不存在。
- [ ] 用 fake repo 测试注册成功。
- [ ] 给 handler 写至少一个 `httptest` 测试。

### 10. GORM 查询、分页、事务

进度：未开始

学习重点：

- [ ] `Where`、`Take`、`Find`。
- [ ] `Limit`、`Offset`。
- [ ] `Order`。
- [ ] `Transaction(func(tx *gorm.DB) error { ... })`。
- [ ] 区分“业务分页”和“数据库分页”。

项目练习：

- [ ] 加用户列表接口：`GET /v1/user?page=1&page_size=20`。
- [ ] handler 解析分页参数。
- [ ] service 处理默认值和边界。
- [ ] repo 使用 `Limit`、`Offset` 查询。
- [ ] 响应里返回列表和分页信息。
- [ ] 如果注册要同时写多张表，再使用事务。

### 11. 数据库迁移和数据初始化

进度：未开始

为什么重要：

真实项目不会只靠手动建表。即使用 GORM AutoMigrate，也需要知道表结构如何随代码演进。

项目练习：

- [ ] 明确当前项目是使用 AutoMigrate、SQL 文件，还是手动建表。
- [ ] 在 README 写清楚建库建表方式。
- [ ] 准备最小初始化数据：部门、角色。
- [ ] 如果暂时不用迁移工具，也要保留 SQL 说明。

### 12. 结构化日志 slog

进度：未开始

项目练习：

- [ ] 用 `log/slog` 替换启动流程里的 `log.Fatalf` 和 `color` 日志。
- [ ] 错误日志带上关键字段，例如 `addr`、`db_name`、`user_id`。
- [ ] 请求失败时记录必要上下文。
- [ ] 不要在日志里输出密码、token、JWT secret。

---

## 阶段四：Go 特性和进阶能力

进度：未开始

目标：不是炫技，而是知道什么时候该用、什么时候不该用。

### 13. switch、iota、业务状态

进度：未开始

学习重点：

- [ ] `iota` 适合定义有限状态。
- [ ] `switch` 适合表达清晰的状态分支。
- [ ] 状态流转要放在 service，不放在 handler。

项目练习：

- [ ] 给用户加 `Status`：正常、禁用、待验证。
- [ ] 状态判断放在 service。
- [ ] handler 只负责把业务错误转换成 HTTP 响应。
- [ ] 禁用用户不能登录。

### 14. goroutine、channel、select

进度：进行中

当前状态：

- [x] `UserService.Register` 已经用了 goroutine/channel 做部门和角色并发校验。
- [ ] 还没有使用 `select`。
- [ ] 还没有处理并发取消。
- [ ] 还没有检查数据竞争风险。

后续学习：

- [ ] channel 关闭。
- [ ] `select` 等待多个 channel。
- [ ] `context` 超时取消。
- [ ] 并发写共享变量时的数据竞争问题。
- [ ] 使用 `go test -race ./...` 检查数据竞争。

提醒：并发不是越多越好，真实项目里优先保证代码简单、正确、可测试。

### 15. sync.Once、sync.Mutex、sync.RWMutex

进度：未开始

适合学习：

- [ ] validator 翻译器只初始化一次：`sync.Once`。
- [ ] 内存缓存的并发保护：`map + sync.RWMutex`。
- [ ] 理解什么时候不应该加缓存。

不建议现在就给 UserService 加缓存。缓存会带来一致性、失效、并发安全问题，容易偏离主线。

### 16. embed、time.Ticker、encoding/json

进度：未开始

放到最后：

- [ ] `embed`：嵌入 SQL、模板、静态文件。
- [ ] `time.Ticker`：定时清理过期 token 或临时数据。
- [ ] `encoding/json`：理解底层 JSON 编解码即可，Gin 项目正常使用 `c.JSON`。

---

## 四周节奏

### 第 1 周：能读懂并改小功能

进度：进行中

- [ ] Go 基础语法迁移。
- [ ] 理解 handler/service/repo。
- [x] 修复模型小问题。
- [x] 打通 context。
- [ ] 新增用户详情接口。
- [ ] 统一 repo 的 not found 判断。

### 第 2 周：能写真实接口

进度：未开始

- [ ] 配置管理。
- [ ] JWT 中间件。
- [ ] 用户登录/注册密码哈希。
- [ ] 统一错误和响应。
- [ ] 接口调试与文档。

### 第 3 周：能测试和重构

进度：未开始

- [ ] 为 service 抽小接口。
- [ ] 写 table-driven tests。
- [ ] 学会 fake repo。
- [ ] 写 handler 的 `httptest`。
- [ ] 理解 Go 的错误处理和边界测试。

### 第 4 周：能准备面试项目

进度：未开始

- [ ] 用户列表分页。
- [ ] 事务。
- [ ] 数据库迁移或建表说明。
- [ ] slog 日志。
- [ ] 状态枚举。
- [ ] 整理 README：项目结构、启动方式、接口说明、技术点。

---

## 每次练习的完成标准

进度：长期执行

1. [ ] 改动保持小而集中。
2. [ ] 不为学语法硬改业务。
3. [ ] 遵守项目分层：handler/service/repo。
4. [ ] 执行 `gofmt -w cmd internal`。
5. [ ] 执行 `go test ./...`。
6. [ ] 如果涉及并发，执行 `go test -race ./...`。
7. [ ] 提交信息使用 Conventional Commits，例如 `feat(user): add user detail api`。

---

## 面试优先级

进度：进行中

如果时间紧，优先掌握这些：

1. [ ] Go 的 `struct`、interface、error、context。
2. [ ] Gin 路由、中间件、参数绑定、统一响应。
3. [ ] GORM 基础 CRUD、分页、事务。
4. [ ] 单元测试和 table-driven tests。
5. [ ] JWT、bcrypt、配置管理、结构化日志。
6. [ ] goroutine/channel/select 的基本使用和常见风险。
7. [ ] 能讲清楚这个项目的目录结构、请求链路、错误处理和安全设计。
