# llm-2-golang 项目

这是一个基于Gin框架的Go Web应用项目。

## 快速开始

### 开发环境启动

根据不同操作系统使用对应的启动脚本：

**启动命令:**

```bash
air
```

或指定配置文件（如有需要）：

```bash
air -c .air.toml
```

### 手动构建

```bash
# 构建当前平台
go build -o tmp/main ./cmd/server

# 运行
./tmp/main  # Unix/Linux/macOS
tmp/main.exe  # Windows
```

## 环境要求

- Go 1.25.0
- Air (热重载工具)

## 安装Air

```bash
go install github.com/air-verse/air@latest
```