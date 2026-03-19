# BI 系统开发 Checklist

> 按顺序开发，遵循依赖关系。点击复选框即可标记任务完成。

---

## 阶段 1: 基础架构 (第 1-2 周)

### 1.1 项目初始化

- [ ] **T001** 初始化 Go 项目 (go mod init)
  - [ ] 创建项目目录结构
  - [ ] 执行 `go mod init github.com/xxx/bi-system`
  - [ ] 创建 config, internal, pkg, cmd 等目录
  - 前置: 无

- [ ] **T002** 搭建 Gin 框架基础结构
  - [ ] 引入 Gin 依赖 (github.com/gin-gonic/gin)
  - [ ] 创建 main.go 入口文件
  - [ ] 配置路由组 (api, web, ws)
  - [ ] 添加基础中间件 (Logger, Recovery, CORS)
  - 前置: T001

- [ ] **T003** 初始化 Vue 3 前端项目 (npm create vite@latest)
  - [ ] 执行 `npm create vite@latest bi-web -- --template vue`
  - [ ] 安装依赖 `npm install`
  - [ ] 配置项目名称和基础信息
  - 前置: 无

- [ ] **T004** 配置前端路由和基础布局
  - [ ] 安装 vue-router: `npm install vue-router`
  - [ ] 创建 router/index.js 路由配置
  - [ ] 创建 layouts 基础布局组件
  - [ ] 配置登录页、注册页、首页路由
  - 前置: T003

- [ ] **T005** 配置 Element Plus 和 ECharts
  - [ ] 安装 Element Plus: `npm install element-plus`
  - [ ] 安装 ECharts: `npm install echarts vue-echarts`
  - [ ] 按需引入或全量引入组件库
  - [ ] 创建全局样式配置
  - 前置: T003

### 1.2 数据库设计

- [ ] **T006** 设计并创建 PostgreSQL 数据库表结构
  - [ ] 设计 users 用户表 (id, username, email, password_hash, created_at, updated_at)
  - [ ] 设计 roles 角色表 (id, name, description, created_at)
  - [ ] 设计 permissions 权限表 (id, name, code, description)
  - [ ] 设计 column_permissions 列权限表 (id, role_id, datasource_id, table_name, columns)
  - [ ] 设计 datasources 数据源表 (id, name, type, host, port, database, username, password)
  - [ ] 设计 charts 图表表 (id, name, type, config, dashboard_id)
  - [ ] 设计 dashboards 仪表盘表 (id, name, description, layout, created_by, published)
  - [ ] 编写建表 SQL 脚本
  - 前置: T002

- [ ] **T007** 配置 GORM 数据库连接
  - [ ] 安装 GORM: `go get gorm.io/gorm`
  - [ ] 安装 PostgreSQL 驱动: `go get gorm.io/driver/postgres`
  - [ ] 创建 database/config.go 配置文件
  - [ ] 创建 database/connection.go 连接管理
  - [ ] 实现数据库自动迁移
  - 前置: T006

- [ ] **T008** 配置 Redis 连接
  - [ ] 安装 Redis 客户端: `go get github.com/redis/go-redis/v9`
  - [ ] 创建 cache/redis.go 连接配置
  - [ ] 配置连接池参数
  - [ ] 添加连接测试代码
  - 前置: T002

### 1.3 认证模块

- [ ] **T009** 实现用户注册 API
  - [ ] 创建 user/model.go 用户模型
  - [ ] 创建 user/dto.go 数据传输对象
  - [ ] 实现 Register handler
  - [ ] 实现密码哈希加密 (bcrypt)
  - [ ] 实现邮箱/用户名重复校验
  - [ ] 编写单元测试
  - 前置: T007

- [ ] **T010** 实现用户登录 API (JWT)
  - [ ] 实现 Login handler
  - [ ] 安装 JWT 库: `go get github.com/golang-jwt/jwt/v5`
  - [ ] 实现密码验证逻辑
  - [ ] 生成 JWT token
  - [ ] 实现 refresh token 机制
  - 前置: T007, T009

- [ ] **T011** 前端登录页面开发
  - [ ] 创建 views/auth/Login.vue
  - [ ] 实现表单验证 (element-plus form)
  - [ ] 调用登录 API
  - [ ] 保存 JWT 到 localStorage
  - [ ] 实现登录后跳转
  - 前置: T004, T010

- [ ] **T012** 前端注册页面开发
  - [ ] 创建 views/auth/Register.vue
  - [ ] 实现表单验证
  - [ ] 调用注册 API
  - [ ] 实现注册成功跳转登录
  - 前置: T004, T009

- [ ] **T013** 实现 JWT 认证中间件
  - [ ] 创建 middleware/auth.go
  - [ ] 解析 Authorization header
  - [ ] 验证 JWT token 有效性
  - [ ] 提取用户信息注入 context
  - [ ] 实现 token 过期处理
  - 前置: T010

- [ ] **T014** 实现路由守卫 (前端)
  - [ ] 创建 router/guard.js
  - [ ] 实现未登录跳转登录页
  - [ ] 实现权限校验逻辑
  - [ ] 处理 token 过期情况
  - 前置: T011, T013

---

## 阶段 2: 核心功能 (第 3-5 周)

### 2.1 用户管理

- [ ] **T015** 用户列表 API
  - [ ] 实现 ListUsers handler
  - [ ] 支持分页参数 (page, page_size)
  - [ ] 支持搜索过滤 (username, email)
  - [ ] 返回用户列表和总数
  - 前置: T013

- [ ] **T016** 用户编辑/删除 API
  - [ ] 实现 UpdateUser handler
  - [ ] 实现 DeleteUser handler
  - [ ] 实现批量删除
  - [ ] 软删除与硬删除选择
  - 前置: T015

- [ ] **T017** 前端用户管理页面
  - [ ] 创建 views/admin/UserManage.vue
  - [ ] 实现用户列表展示 (element-plus table)
  - [ ] 实现分页组件
  - [ ] 实现搜索过滤功能
  - [ ] 添加编辑/删除操作按钮
  - 前置: T015

### 2.2 角色与权限

- [ ] **T018** 角色 CRUD API
  - [ ] 创建 role/model.go
  - [ ] 实现 CreateRole handler
  - [ ] 实现 UpdateRole handler
  - [ ] 实现 DeleteRole handler
  - [ ] 实现 ListRoles handler
  - 前置: T013

- [ ] **T019** 权限配置 API
  - [ ] 创建 permission/model.go
  - [ ] 实现权限列表 API
  - [ ] 实现角色权限分配 API
  - [ ] 实现获取角色权限 API
  - 前置: T018

- [ ] **T020** 数据列级别权限 API
  - [ ] 创建 column_permission/model.go
  - [ ] 实现列权限配置 API
  - [ ] 实现数据源/表/列权限设置
  - [ ] 实现获取用户可访问列 API
  - 前置: T019

- [ ] **T021** 前端角色管理页面
  - [ ] 创建 views/admin/RoleManage.vue
  - [ ] 实现角色列表展示
  - [ ] 实现角色创建/编辑弹窗
  - [ ] 实现角色删除确认
  - 前置: T018

- [ ] **T022** 前端权限配置界面
  - [ ] 创建 views/admin/PermissionConfig.vue
  - [ ] 实现权限树形展示
  - [ ] 实现角色权限分配勾选
  - [ ] 实现数据列权限配置表格
  - 前置: T019, T021

### 2.3 数据源管理

- [ ] **T023** 数据源添加/编辑/删除 API
  - [ ] 创建 datasource/model.go
  - [ ] 实现 CreateDatasource handler
  - [ ] 实现 UpdateDatasource handler
  - [ ] 实现 DeleteDatasource handler
  - [ ] 实现 ListDatasources handler
  - [ ] 密码加密存储
  - 前置: T013

- [ ] **T024** 数据源连接测试 API
  - [ ] 实现 TestConnection handler
  - [ ] 支持 MySQL/PostgreSQL 连接测试
  - [ ] 返回连接状态和错误信息
  - 前置: T023

- [ ] **T025** 表结构浏览 API
  - [ ] 实现 GetTables handler (获取数据源下的表列表)
  - [ ] 实现 GetColumns handler (获取表的列信息)
  - [ ] 实现元数据缓存
  - 前置: T024

- [ ] **T026** 前端数据源管理页面
  - [ ] 创建 views/datasource/DatasourceManage.vue
  - [ ] 实现数据源列表展示
  - [ ] 实现添加/编辑数据源弹窗
  - [ ] 实现连接测试按钮
  - [ ] 实现删除确认
  - 前置: T023, T024

- [ ] **T027** 前端表结构浏览界面
  - [ ] 创建 views/datasource/TableBrowse.vue
  - [ ] 实现左侧数据源树形选择
  - [ ] 实现右侧表/列信息展示
  - [ ] 实现表结构详情弹窗
  - 前置: T025, T026

### 2.4 数据查询

- [ ] **T028** SQL 查询执行 API
  - [ ] 创建 query/service.go
  - [ ] 实现 ExecuteQuery handler
  - [ ] 实现查询参数校验
  - [ ] 实现结果集转换
  - [ ] 应用行列权限过滤
  - 前置: T013, T020

- [ ] **T029** 查询结果缓存
  - [ ] 实现缓存服务 cache/query_cache.go
  - [ ] 生成查询结果 Hash Key
  - [ ] 设置缓存过期时间
  - [ ] 实现缓存失效策略
  - 前置: T028, T008

- [ ] **T030** 查询历史记录 API
  - [ ] 创建 query_history/model.go
  - [ ] 实现保存查询历史
  - [ ] 实现查询历史列表 API
  - [ ] 实现查询历史详情 API
  - [ ] 实现清理历史记录 API
  - 前置: T028

- [ ] **T031** 前端 SQL 编辑器页面
  - [ ] 创建 views/query/SqlEditor.vue
  - [ ] 集成 CodeMirror 或 Monaco Editor
  - [ ] 实现 SQL 语法高亮
  - [ ] 实现格式化功能
  - [ ] 实现快捷键支持
  - 前置: T028

- [ ] **T032** 前端查询结果展示
  - [ ] 创建 views/query/QueryResult.vue
  - [ ] 实现表格展示查询结果
  - [ ] 实现分页展示
  - [ ] 实现列排序功能
  - [ ] 实现列宽调整
  - 前置: T028, T031

- [ ] **T033** 数据导出功能 (CSV)
  - [ ] 实现 ExportCSV handler
  - [ ] 实现大文件分片导出
  - [ ] 前端添加导出按钮
  - [ ] 实现进度显示
  - 前置: T032

---

## 阶段 3: 可视化 (第 6-8 周)

### 3.1 图表模块

- [ ] **T034** 图表数据模型设计
  - [ ] 创建 chart/model.go
  - [ ] 设计图表配置结构 (title, type, xAxis, yAxis, series)
  - [ ] 设计数据绑定配置
  - [ ] 支持图表类型枚举
  - 前置: T006

- [ ] **T035** 图表创建/编辑 API
  - [ ] 实现 CreateChart handler
  - [ ] 实现 UpdateChart handler
  - [ ] 实现图表配置校验
  - [ ] 实现图表预览数据生成
  - 前置: T034, T028

- [ ] **T036** 图表删除 API
  - [ ] 实现 DeleteChart handler
  - [ ] 实现批量删除
  - [ ] 删除时检查仪表盘引用
  - 前置: T035

- [ ] **T037** 前端图表组件 (ECharts 封装)
  - [ ] 创建 components/charts/BaseChart.vue
  - [ ] 实现折线图组件 LineChart
  - [ ] 实现柱状图组件 BarChart
  - [ ] 实现饼图组件 PieChart
  - [ ] 实现仪表盘组件 GaugeChart
  - [ ] 实现地图组件 MapChart
  - 前置: T005

- [ ] **T038** 前端图表创建向导
  - [ ] 创建 views/chart/ChartWizard.vue
  - [ ] 实现步骤条组件 (选择数据源 -> 选择表/查询 -> 配置图表 -> 预览)
  - [ ] 实现图表类型选择
  - [ ] 实现数据绑定配置
  - 前置: T035, T037

- [ ] **T039** 前端图表配置界面
  - [ ] 创建 views/chart/ChartConfig.vue
  - [ ] 实现标题配置
  - [ ] 实现坐标轴配置
  - [ ] 实现系列配置
  - [ ] 实现样式配置
  - 前置: T038

- [ ] **T040** 图表数据绑定功能
  - [ ] 实现数据源选择组件
  - [ ] 实现自定义 SQL 查询绑定
  - [ ] 实现字段映射配置
  - [ ] 实现数据预览
  - 前置: T039, T025

### 3.2 仪表盘模块

- [ ] **T041** 仪表盘数据模型设计
  - [ ] 创建 dashboard/model.go
  - [ ] 设计仪表盘布局结构 (grid, row, col, widget)
  - [ ] 设计发布配置
  - [ ] 设计分享配置
  - 前置: T034

- [ ] **T042** 仪表盘 CRUD API
  - [ ] 实现 CreateDashboard handler
  - [ ] 实现 UpdateDashboard handler
  - [ ] 实现 DeleteDashboard handler
  - [ ] 实现 ListDashboards handler
  - [ ] 实现 GetDashboard handler
  - 前置: T041

- [ ] **T043** 仪表盘发布/分享 API
  - [ ] 实现 PublishDashboard handler
  - [ ] 实现取消发布 handler
  - [ ] 实现生成分享链接
  - [ ] 实现分享权限控制
  - 前置: T042

- [ ] **T044** 前端仪表盘列表页面
  - [ ] 创建 views/dashboard/DashboardList.vue
  - [ ] 实现仪表盘卡片展示
  - [ ] 实现创建/编辑/删除操作
  - [ ] 实现发布状态标识
  - 前置: T042

- [ ] **T045** 前端仪表盘拖拽布局
  - [ ] 创建 views/dashboard/DashboardEditor.vue
  - [ ] 集成 vue-grid-layout
  - [ ] 实现图表拖拽添加
  - [ ] 实现布局调整保存
  - [ ] 实现网格吸附
  - 前置: T044

- [ ] **T046** 前端仪表盘预览功能
  - [ ] 实现仪表盘预览页面
  - [ ] 实现全屏预览
  - [ ] 实现刷新数据
  - [ ] 实现时间范围选择
  - 前置: T045

---

## 阶段 4: 完善与优化 (第 9-10 周)

### 4.1 权限细粒度控制

- [ ] **T047** 行级别数据过滤
  - [ ] 设计行级权限配置结构
  - [ ] 实现行级权限过滤中间件
  - [ ] 在查询时注入行级过滤条件
  - [ ] 支持动态行级条件
  - 前置: T020

- [ ] **T048** API 权限验证
  - [ ] 实现 API 权限码配置
  - [ ] 实现权限验证中间件
  - [ ] 实现权限不足统一响应
  - 前置: T019

### 4.2 性能优化

- [ ] **T049** 查询性能优化
  - [ ] 实现查询超时控制
  - [ ] 实现查询限流
  - [ ] 优化 SQL 执行计划
  - [ ] 添加查询索引建议
  - 前置: T028

- [ ] **T050** 前端加载优化
  - [ ] 实现路由懒加载
  - [ ] 实现组件懒加载
  - [ ] 实现 ECharts 按需加载
  - [ ] 添加骨架屏
  - 前置: T032

### 4.3 测试与修复

- [ ] **T051** 单元测试
  - [ ] 编写 Service 层单元测试
  - [ ] 编写 Handler 层单元测试
  - [ ] 使用 gomock 进行 mock
  - [ ] 统计测试覆盖率
  - 前置: T050

- [ ] **T052** 集成测试
  - [ ] 编写 API 集成测试
  - [ ] 编写数据库集成测试
  - [ ] 编写前端 E2E 测试
  - [ ] 使用 Testcontainers
  - 前置: T051

- [ ] **T053** Bug 修复
  - [ ] 修复测试发现的问题
  - [ ] 修复安全漏洞
  - [ ] 优化错误处理
  - [ ] 完善日志记录
  - 前置: T052

### 4.4 部署准备

- [ ] **T054** Docker 配置
  - [ ] 创建 backend/Dockerfile
  - [ ] 创建 frontend/Dockerfile
  - [ ] 创建 docker-compose.yml
  - [ ] 配置环境变量
  - [ ] 配置 Nginx 反向代理
  - 前置: T053

- [ ] **T055** 部署文档
  - [ ] 编写部署手册
  - [ ] 编写环境配置说明
  - [ ] 编写备份恢复方案
  - [ ] 编写监控告警配置
  - 前置: T054

---

## 开发顺序建议

### 第一周 (T001-T008)
```
T001 → T002 → T003 → T004 → T005 → T006 → T007 → T008
```

### 第二周 (T009-T014)
```
T009 → T010 → T011 → T012 → T013 → T014
```

### 第三周 (T015-T022)
```
T015 → T016 → T017
T018 → T019 → T020 → T021 → T022
```

### 第四周 (T023-T027)
```
T023 → T024 → T025 → T026 → T027
```

### 第五周 (T028-T033)
```
T028 → T029 → T030 → T031 → T032 → T033
```

### 第六周 (T034-T040)
```
T034 → T035 → T036
T037 → T038 → T039 → T040
```

### 第七周 (T041-T046)
```
T041 → T042 → T043
T044 → T045 → T046
```

### 第八周 (T047-T050)
```
T047 → T048 → T049 → T050
```

### 第九周 (T051-T055)
```
T051 → T052 → T053 → T054 → T055
```

---

## 进度统计

| 阶段 | 任务数 | 已完成 | 进度 |
|------|--------|--------|------|
| 阶段1: 基础架构 | 14 | 0 | 0% |
| 阶段2: 核心功能 | 19 | 0 | 0% |
| 阶段3: 可视化 | 13 | 0 | 0% |
| 阶段4: 完善与优化 | 9 | 0 | 0% |
| **总计** | **55** | **0** | **0%** |
