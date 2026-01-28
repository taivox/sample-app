# sample-app

React + Go Fiber + PostgreSQL auth app.

## Structure
- `backend/` - Go Fiber API (port 8080)
- `frontend/` - React/Vite app (port 3000)

## Run locally
```
docker-compose up -d db
cd backend && go run cmd/main.go
cd frontend && npm run dev
```

## Token Usage
For verbose commands (`npm install`, `yarn upgrade`), instruct user to run manually.

## Dependencies
Read changelog before updating. Warn about breaking changes.
