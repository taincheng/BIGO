# BI 系统

企业级商业智能系统，提供数据可视化、仪表盘、多数据源连接和完整权限管理功能。

## 技术栈

### 前端
- Vue 3 + TypeScript
- Vite
- Element Plus
- ECharts
- Pinia
- Vue Router

### 后端
- Go 1.21+
- Gin
- GORM
- PostgreSQL
- Redis

## 项目结构

```
bi-system/
├── frontend/                    # 前端应用 (Vue 3)
│   └── src/
│       ├── api/                # API 服务
│       ├── assets/             # 静态资源
│       ├── components/         # 通用组件
│       ├── router/             # 路由配置
│       ├── stores/             # Pinia 状态管理
│       ├── types/              # TypeScript 类型
│       ├── utils/             # 工具函数
│       └── views/              # 页面视图
│
├── backend/                     # 后端应用 (Go)
│   ├── cmd/server/             # 主入口
│   ├── internal/               # 内部代码 (仅本项目可用)
│   │   ├── config/             # 配置管理
│   │   ├── handlers/           # HTTP 处理器 (路由逻辑)
│   │   ├── middleware/         # 中间件 (认证、日志、限流)
│   │   ├── models/             # 数据模型 (GORM struct)
│   │   ├── services/           # 业务逻辑层
│   │   └── repositories/        # 数据访问层 (数据库操作)
│   └── pkg/                    # 公共代码 (可被外部项目导入)
│       ├── utils/              # 工具函数
│       ├── response/           # 统一响应结构
│       └── errors/             # 错误定义
│
└── docker-compose.yml          # Docker 配置
```

## 目录说明

### `internal/` vs `pkg/`

| 目录 | 访问权限 | 用途 |
|------|---------|------|
| `internal` | 仅项目内 | 业务逻辑、API、数据库 |
| `pkg` | 公开 | 通用工具、可复用组件 |

- **`internal`**: 只能被本项目内部模块引用，**不能**被外部项目导入
- **`pkg`**: 可以被外部项目导入，用于存放可复用的工具/库

## 快速开始

### 后端

```bash
cd backend
go mod init bi-system
go mod tidy
go run cmd/server/main.go
```

### 前端

```bash
cd frontend
npm install
npm run dev
```

## 开发任务

详见 [CHECKLIST.md](./CHECKLIST.md)
