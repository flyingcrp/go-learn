# Day 1 — 工程初始化与环境搭建

## 今日学习主题
工程初始化、Go Modules 包管理、air 热重载配置、MySQL 驱动引入与依赖注入实践。

## 关键点

### 1. 工程初始化
- 从 GitHub 克隆项目后使用 `go mod init <模块名>` 初始化，模块名与仓库路径保持一致（如 `github.com/username/project`）
- Go 的模块名直接影响后续 `import` 路径，命名规范很重要

### 2. Go CLI 与包管理
- `go get`：添加/更新依赖；`go mod tidy`：清理未使用的依赖并补全缺失项
- 版本升级策略：`go get -u=patch` 只升 patch，`go get -u=minor` 只升 minor，`go get -u` 升最新
- Major 版本升级（v1 → v2）需在 import 路径中显式携带 `/v2`
- 模块缓存在 `~/go/pkg/mod`，可用 `go clean -modcache` 清理

### 3. air 热重载
- 配置 `.air.toml` 实现文件变更后自动编译重启
- 编译输出目录通过 `cmd` 的 `-o` 参数和 `bin` 字段指定，避免污染源码
- 支持 `env_files` 加载本地 `.env`，可配置跨平台编译参数
- `send_interrupt = true` 实现优雅关闭，`exclude_regex` 可排除测试文件避免误触发

### 4. MySQL 与依赖注入
- 引入 `go-sql-driver/mysql` + `gorm.io/gorm` 组合
- 摒弃全局单例模式，改为 `main` 中初始化 `Infra` 对象，通过构造函数显式注入各模块
- `NewInfra()` 返回 `(infra, cleanup, error)` 三元组，生命周期管理清晰

## 常见用途

| 内容 | 用途 |
|------|------|
| Go Modules | 管理项目依赖版本，确保构建可复现 |
| air 热重载 | 本地开发时自动重编译，提升开发效率 |
| 构造函数注入 | 模块间解耦，便于单元测试和后续替换实现 |
| Infra 聚合层 | 统一管理 DB/Cache/MQ 等基础设施，扩展时无需改动模块接口 |

## 总结
第一天完成了工程脚手架搭建，掌握了 Go 的包管理机制和本地开发工具链。最关键的认知转变是从 Node.js 框架的自动 DI 转向 Go 的显式依赖管理——虽然繁琐，但代码路径清晰、测试友好。
