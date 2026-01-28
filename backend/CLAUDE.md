# backend

## Verify
`gofumpt -w . && golangci-lint run && go build .`

## Structure
- `cmd/main.go` - entry point
- `packages/api/` - routes, controllers, middleware
- `packages/db/` - queries, migrations, models
- `packages/config/` - env config

## Env
Requires `.env` with: POSTGRES_*, CLIENT_URL, SERVER_PORT, JWT_KEY
