# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

BIGO server is a Go backend API system using Gin framework with JWT authentication and Casbin RBAC permission control. It provides user management, role/authority management, and menu-based access control.

## Commands

```bash
# Run the server
go run main.go

# Generate Swagger documentation (after modifying API comments)
go generate ./...

# Build
go build -o bigo-server .
```

## Architecture

### Layered Structure
```
api/          # API handlers (receive requests, call services)
service/      # Business logic layer
model/        # Data models and DTOs
router/       # Route definitions
middleware/   # JWT auth, Casbin RBAC
initialize/   # Startup initialization (DB, logging, routing)
core/         # Core utilities (viper config, zap logger, server)
config/       # Configuration structs
utils/        # JWT, Casbin, encryption helpers
global/       # Global singletons (DB, logger, config)
```

### Entry Point Flow
`main.go` → `core.Viper()` (load config) → `core.Zap()` (init logger) → `initialize.Gorm()` (DB) → `initialize.Routers()` → `core.RunServer()`

### Dependency Injection Pattern
Services are accessed via singleton pattern:
```go
ServiceGroupApp.SystemServiceGroup.UserService
```

### Key Technologies
- **Framework**: Gin v1.12
- **ORM**: GORM v1.31
- **Auth**: JWT (golang-jwt/jwt/v5)
- **RBAC**: Casbin v3 with gorm-adapter
- **Config**: Viper
- **Logging**: Zap
- **API Docs**: Swagger

### Authentication Flow
1. JWT middleware (`middleware/jwt.go`) extracts token from header
2. Parses and validates token, extracts claims
3. Auto-refreshes token if within buffer window
4. Sets claims in context via `global.ClaimsKey`

### RBAC/Permission Flow
1. `CasbinHandler()` middleware runs after JWT auth
2. Extracts `AuthorityId` from JWT claims
3. Enforces Casbin policy: `sub=role_id, obj=path, act=method`
4. Uses `keyMatch2` for path matching

### Database Models
- `SysUser` - Users with UUID, password (bcrypt), roles
- `SysAuthority` - Roles with hierarchy
- `SysBaseMenu` - Menu permissions (tree structure)
- `SysBaseMenuBtn` - Button-level permissions
- `SysLoginLog` - Login audit trail
- `CasbinRule` - Permission policies (managed by Casbin)

### Configuration (config.yaml)
- `system.env`: local/dev/pub - controls log output
- `jwt.signing-key`: JWT secret
- `jwt.expires-time`: Token TTL (seconds)
- `system.disable-auto-migrate`: Set true in production

### Adding New Modules
1. Create model in `model/system/`
2. Add service logic in `service/system/`
3. Add API handlers in `api/v1/system/`
4. Register routes in `router/system/`
5. Add to `service/enter.go` and `api/v1/system/enter.go`
