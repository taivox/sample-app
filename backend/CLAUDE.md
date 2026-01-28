# backend

## Verify
`gofumpt -w . && golangci-lint run && go build ./cmd/main.go && go test ./...`

## Test
```
# Unit tests only
go test ./packages/...

# Integration tests (requires running database)
docker compose up -d db
go test ./tests/... -v
```

## Structure
- `cmd/main.go` - entry point
- `packages/api/` - routes, controllers, middleware
- `packages/db/` - queries, migrations, models
- `packages/config/` - env config
- `packages/utils/` - validation helpers
- `tests/` - integration tests

## Env
Requires `.env` with: POSTGRES_*, CLIENT_URL, SERVER_PORT, JWT_KEY
Use `.env.test` for integration tests (port 8081)
