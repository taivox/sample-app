# Sample App

A sample full-stack authentication app built for learning and practicing.

## Tech Stack

| Layer    | Technology                   |
| -------- | ---------------------------- |
| Frontend | React 19, Vite, React Router |
| Backend  | Go 1.25, Fiber               |
| Database | PostgreSQL                   |

## Quick Start

### Run with Docker

```bash
docker compose build
docker compose up -d
```

Access the app at **http://localhost:3000**

### Clean Up

```bash
docker compose down -v
```

## API Endpoints

| Endpoint        | Method | Body                        | Description         |
| --------------- | ------ | --------------------------- | ------------------- |
| `/api/ping`     | GET    | -                           | Health check        |
| `/api/register` | POST   | `{ name, email, password }` | Register user       |
| `/api/login`    | POST   | `{ email, password }`       | Login user          |
| `/api/session`  | GET    | -                           | Get current session |
| `/api/logout`   | GET    | -                           | Logout user         |
