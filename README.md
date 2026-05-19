an e-commerce admin and mall API. This repo provides REST APIs for shop management, RBAC, WeChat integration, orders, and related features.

> **Module path:** Imports use `yixiang.co/go-mall` (see `go.mod`). That is the Go module name, not a folder on disk—the code lives in this repository root.

## Tech stack

| Layer | Technology |
|-------|------------|
| HTTP | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io/) + MySQL 8 |
| Auth | JWT + [Casbin](https://casbin.org/) (RBAC) |
| Cache | Redis |
| Config | Viper (`config.yaml`) |
| Logging | Zap |
| Docs | Swagger (`/swagger/*any`) |
| CLI | Cobra (`serve` command) |
| Other | WeChat SDK, gopay, cron, code generator |

## Related repositories

| Component | Gitee | GitHub |
|-----------|-------|--------|
| Backend (this repo) | [yshop-gin](https://gitee.com/guchengwuyue/yshop-gin) | [yshop-gin](https://github.com/guchengwuyue/yshop-gin) |
| Admin UI (Vue) | [yshop-gin-vue](https://gitee.com/guchengwuyue/yshop-gin-vue) | [yshop-gin-vue](https://github.com/guchengwuyue/yshop-gin-vue) |

The PC storefront frontend is bundled separately (see upstream docs under `pc-vue/`).

## Features

**Admin / system**

- Users, roles, departments, jobs
- Menus with dynamic routes (Casbin URL + method checks)
- Data dictionaries, operation logs
- Material library, code generator, scheduled tasks

**Shop**

- Categories, SKU rules, products (single/multi spec)
- Orders, shipping, express (Kdniao)
- Canvas / homepage content

**WeChat**

- Official account menus, users, articles

**Mall API (storefront)**

- Described in upstream docs (login, cart, orders, pay, etc.). In this tree, `api.RegisterApiRouters` is commented out in `routers/router.go`; the active surface is mainly admin/management routes.

## Project layout

```
yshop-gin/
├── main.go              # Entry: init infra, Cobra root command
├── cmd/                 # CLI commands (default: serve)
├── config.yaml          # App, DB, Redis, WeChat, express settings
├── conf/                # Config struct definitions + rbac_model.conf
├── routers/             # HTTP route registration
│   └── admin/           # /admin, /shop, /weixin, /tools, /auth
├── middleware/          # CORS, JWT, Casbin, request logging
├── app/
│   ├── controllers/     # HTTP handlers
│   ├── service/         # Business logic (dto/, vo/ subpackages)
│   ├── models/          # GORM models and DB helpers
│   ├── params/          # Request binding structs
│   └── listen/          # Redis keyspace listener
├── pkg/                 # Shared packages (jwt, redis, upload, …)
├── docs/                # Swagger generated docs
├── sql/                 # Database schema dumps
├── template/            # Code generator templates
└── runtime/             # Runtime files (logs, uploads, etc.)
```

### Request flow

```
HTTP → routers → middleware (JWT, log, Casbin) → controller → service → models → MySQL
```

Global dependencies (`YSHOP_DB`, `YSHOP_CONFIG`, `YSHOP_LOG`, …) are initialized in `main.go` `init()` and stored in `pkg/global`.

## Prerequisites

- Go **1.15+** (see `go.mod`)
- MySQL **8**
- Redis

## Quick start

1. **Clone and install dependencies**

   ```bash
   git clone <your-remote-url> yshop-gin
   cd yshop-gin
   go mod tidy
   ```

2. **Database**

   Import the schema:

   ```bash
   mysql -u root -p yshop_go < sql/yshop_go.sql
   ```

   Create the database first if needed: `CREATE DATABASE yshop_go CHARACTER SET utf8mb4;`

3. **Configuration**

   Edit `config.yaml`—at minimum `database` and `redis`:

   ```yaml
   database:
     user: root
     password: your_password
     host: 127.0.0.1:3306
     name: yshop_go
   redis:
     host: 127.0.0.1:6379
   ```

4. **Run**

   ```bash
   go run main.go serve
   ```

   Or with [Air](https://github.com/air-verse/air) for live reload:

   ```bash
   air -c .air.conf
   ```

   Default HTTP port: **8000** (`config.yaml` → `server.http-port`).

5. **Verify**

   - API base: `http://127.0.0.1:8000`
   - Swagger UI: `http://127.0.0.1:8000/swagger/index.html`
   - Login: `POST /auth/login` (use admin UI or Swagger)

## API route groups

| Prefix | Purpose | Auth |
|--------|---------|------|
| `/auth/*` | Login, captcha | Public |
| `/admin/*` | System RBAC, users, menus, dicts, materials | JWT + Casbin |
| `/shop/*` | Products, categories, orders, express | JWT + Casbin |
| `/weixin/*` | WeChat admin | JWT + Casbin |
| `/tools/*` | Codegen, cron jobs | JWT + Casbin |
| `/upload/images` | Static uploads | — |

Protected routes expect: `Authorization: Bearer <token>`.

## Build & deploy

```bash
go build -o yshop-gin .
./yshop-gin serve
```

Run behind nginx or another reverse proxy in production. Set `server.run-mode` to `release` in `config.yaml`.

## Configuration reference

| File | Role |
|------|------|
| `config.yaml` | Main settings (Viper, hot-reload supported) |
| `conf/rbac_model.conf` | Casbin RBAC model |
| `-c path` | Alternate config file (see `pkg/config/viper.go`) |
| `--env` | Cobra flag for env-specific config (see `cmd/cmd.go`) |

## Development notes

- **Import path vs directory:** All packages import as `yixiang.co/go-mall/...` while the repo folder may be named `yshop-gin`. To rename the module, change the `module` line in `go.mod` and update imports project-wide.
- **Soft delete:** Models embedding `BaseModel` use GORM soft delete (`isDel`).
- **Permissions:** Non-`admin` roles are checked with Casbin against request path and HTTP method.