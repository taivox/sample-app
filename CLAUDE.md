# sample-app

React + Go Fiber + PostgreSQL auth app for learning and practicing.

## Structure
- `backend/` - Go Fiber API (port 8080)
- `frontend/` - React/Vite app (port 3000)

## Run locally
```
docker compose up -d db
cd backend && go run cmd/main.go
cd frontend && npm run dev
```

## Test
```
cd backend && go test ./...
cd frontend && npm run test:run
```

## Token Usage
For verbose commands (`npm install`, `yarn upgrade`), instruct user to run manually.

## Dependencies
Read changelog before updating. Warn about breaking changes.
