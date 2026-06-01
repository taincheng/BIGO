# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

BIGO is a monorepo containing a Go backend API system (`server/`) and a Vue 3 frontend (`web/`). The system provides user authentication, role-based access control (RBAC), and menu-based permission management.

## Directory Structure

```
server/          # Go backend (Gin + GORM)
  api/v1/        # HTTP handlers
  service/       # Business logic
  model/         # Data models and DTOs
  router/       # Route definitions
  middleware/    # JWT auth, Casbin RBAC
  initialize/   # Startup initialization
  core/          # Config loader, logger, server runner
  config/        # Configuration structs
  utils/         # JWT, Casbin, encryption, upload helpers
  global/        # Global singletons (DB, logger, config)
  main.go        # Entry point

web/             # Vue 3 frontend (pure-admin-thin i18n version)
  src/api/       # Axios API clients
  src/views/     # Page components
  src/router/    # Vue Router config
  src/store/      # Pinia stores
  src/utils/      # Utility functions
  locales/       # i18n translation files
```

## Server Commands

```bash
cd server

# Run the server
go run main.go

# Generate Swagger docs (after modifying API comments)
go generate ./...

# Build binary
go build -o bigo-server .
```

## Web Commands

```bash
cd web

# Install dependencies (requires pnpm)
pnpm install

# Development server
pnpm dev

# Production build
pnpm build

# Type check
pnpm typecheck

# Lint
pnpm lint:eslint
pnpm lint:prettier
pnpm lint:stylelint
```

## Architecture

### Server Layered Flow
`main.go` → `core.Viper()` (load config) → `core.Zap()` (init logger) → `initialize.Gorm()` (DB) → `initialize.Routers()` → `core.RunServer()`

### Server Dependency Injection
Services accessed via singleton pattern:
```go
ServiceGroupApp.SystemServiceGroup.UserService
```

### Authentication Flow
1. JWT middleware extracts token from `x-token` header
2. Parses/validates token, auto-refreshes if within buffer window
3. Sets claims in context via `global.ClaimsKey`

### RBAC/Permission Flow
1. `CasbinHandler()` middleware runs after JWT auth
2. Enforces policy: `sub=role_id, obj=path, act=method`
3. Uses `keyMatch2` for path matching

### Key Technologies (Server)
- **Framework**: Gin v1.12
- **ORM**: GORM v1.31
- **Auth**: JWT (golang-jwt/jwt/v5)
- **RBAC**: Casbin v3 with gorm-adapter
- **Config**: Viper
- **Logging**: Zap
- **API Docs**: Swagger

### Key Technologies (Web)
- **Framework**: Vue 3 + TypeScript
- **Build**: Vite 7
- **State**: Pinia
- **UI**: Element Plus
- **i18n**: vue-i18n
- **Styling**: Tailwind CSS + SCSS

## Configuration

Server config is in `server/config.yaml`:
- `system.env`: local/dev/pub - controls log output level
- `jwt.signing-key`: JWT secret
- `jwt.expires-time`: Token TTL in seconds
- `system.disable-auto-migrate`: Set true in production

## Adding New Server Modules

1. Create model in `model/system/`
2. Add service logic in `service/system/`
3. Add API handlers in `api/v1/system/`
4. Register routes in `router/system/`
5. Add to `service/enter.go` and `api/v1/system/enter.go`