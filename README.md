# Go Learn 项目

这是一个基于Gin框架的Go Web应用项目。

## 快速开始

### 开发环境启动

根据不同操作系统使用对应的启动脚本：

**Windows:**

```bash
dev.bat
```

**macOS/Linux:**

```bash
chmod +x dev.sh
./dev.sh
```

**或者手动指定配置文件:**

Windows:

```bash
air -c .air.windows.toml
```

macOS/Linux:

```bash
air -c .air.unix.toml
```

### 配置文件说明

- `.air.toml` - 默认跨平台配置（推荐）
- `.air.windows.toml` - Windows专用配置
- `.air.unix.toml` - Unix/Linux/macOS专用配置

### 手动构建

```bash
# 构建当前平台
go build -o tmp/main ./app

# 运行
./tmp/main  # Unix/Linux/macOS
tmp/main.exe  # Windows
```

## 项目结构

```Text
app/
├── common/
│   ├── middleware/     # 中间件
│   ├── response/       # 响应封装
│   ├── router/         # 路由管理
│   ├── storage/        # 数据存储
│   └── validation/     # 参数验证
├── user/               # 用户模块
│   ├── handler.go      # 处理器
│   ├── request.go      # 请求结构体
│   ├── router.go       # 路由定义
│   └── service.go      # 业务逻辑
└── main.go             # 程序入口
```

## 环境要求

- Go 1.21+
- Air (热重载工具)

## 安装Air

```bash
go install github.com/air-verse/air@latest
```
