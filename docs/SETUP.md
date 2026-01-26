# Setup & Development Guide

## Prerequisites

Sebelum memulai, pastikan Anda sudah install:

- **Go 1.23.1 atau lebih tinggi**
  ```bash
  go version
  ```

- **Docker & Docker Compose**
  ```bash
  docker --version
  docker-compose --version
  ```

- **Git**
  ```bash
  git --version
  ```

- **PostgreSQL Client (optional, untuk direct database access)**
  ```bash
  psql --version
  ```

## Initial Setup

### 1. Clone Repository

```bash
git clone <repository-url>
cd videocall
```

### 2. Install Go Dependencies

```bash
# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Verify dependencies
go mod verify
```

### 3. Setup Environment Variables

```bash
# Copy environment template
cp .env.example .env

# Edit .env dengan konfigurasi lokal
nano .env  # atau gunakan editor favorit Anda
```

**Contoh .env untuk development:**
```env
SERVER_PORT=8080
SERVER_HOST=localhost
ENV=development

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=videocall_db
DB_SSL_MODE=disable

JWT_SECRET=your_secret_key_development
JWT_EXPIRATION=24h

API_VERSION=v1
API_TIMEOUT=30s

LOG_LEVEL=debug
LOG_FORMAT=json
```

### 4. Setup Database

#### Option A: Menggunakan Docker Compose

```bash
# Start database container
docker-compose up -d db

# Verify database is running
docker-compose ps
```

#### Option B: Manual PostgreSQL Setup

```bash
# Create database
createdb videocall_db

# Create user (jika belum ada)
createuser -P postgres

# Grant privileges
psql -U postgres -d videocall_db -c "GRANT ALL PRIVILEGES ON DATABASE videocall_db TO postgres;"
```

### 5. Run Database Migrations

```bash
# Run migrations
go run cmd/server/main.go migrate

# Verify tables created
psql -U postgres -d videocall_db -c "\dt"
```

## Development Workflow

### Running the Application

#### Option 1: Direct Go Run

```bash
go run cmd/server/main.go
```

Server akan berjalan di `http://localhost:8080`

#### Option 2: Using Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f videocall

# Stop services
docker-compose down
```

#### Option 3: Using Make (jika Makefile tersedia)

```bash
make run
```

### Hot Reload Development

Untuk development dengan hot reload, gunakan `air`:

```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

Buat file `.air.toml` di root project:

```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ./cmd/server"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false
```

## Testing

### Run All Tests

```bash
go test ./...
```

### Run Tests with Coverage

```bash
go test -cover ./...
```

### Generate Coverage Report

```bash
# Generate coverage file
go test -coverprofile=coverage.out ./...

# View coverage in browser
go tool cover -html=coverage.out
```

### Run Specific Test

```bash
go test -run TestUserLogin ./internal/app/users/...
```

### Run Integration Tests

```bash
go test -tags=integration ./...
```

## Code Quality

### Format Code

```bash
# Format all Go files
go fmt ./...

# Or using gofmt
gofmt -s -w .
```

### Lint Code

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run
```

### Vet Code

```bash
go vet ./...
```

## Building

### Build Binary

```bash
# Build for current OS
go build -o bin/videocall cmd/server/main.go

# Run binary
./bin/videocall
```

### Build for Different OS

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o bin/videocall-linux cmd/server/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o bin/videocall-darwin cmd/server/main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o bin/videocall.exe cmd/server/main.go
```

## Docker Development

### Build Docker Image

```bash
docker build -t videocall:latest .
```

### Run Docker Container

```bash
docker run -p 8080:8080 \
  --env-file .env \
  --name videocall \
  videocall:latest
```

### Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Remove volumes
docker-compose down -v
```

## Database Management

### Connect to Database

```bash
# Using psql
psql -U postgres -d videocall_db -h localhost

# Using Docker
docker-compose exec db psql -U postgres -d videocall_db
```

### View Tables

```sql
\dt
```

### View Table Schema

```sql
\d users
\d products
\d invoices
```

### Run SQL Query

```bash
psql -U postgres -d videocall_db -c "SELECT * FROM users;"
```

### Backup Database

```bash
pg_dump -U postgres videocall_db > backup.sql
```

### Restore Database

```bash
psql -U postgres videocall_db < backup.sql
```

## Troubleshooting

### Port Already in Use

```bash
# Find process using port 8080
lsof -i :8080

# Kill process
kill -9 <PID>
```

### Database Connection Error

```bash
# Check if database is running
docker-compose ps

# Check database logs
docker-compose logs db

# Verify connection string in .env
```

### Go Module Issues

```bash
# Clear module cache
go clean -modcache

# Re-download dependencies
go mod download

# Tidy dependencies
go mod tidy
```

### Docker Issues

```bash
# Remove all containers
docker-compose down -v

# Rebuild images
docker-compose build --no-cache

# Start fresh
docker-compose up -d
```

## IDE Setup

### VS Code

Install extensions:
- Go (golang.go)
- REST Client (humao.rest-client)
- Docker (ms-azuretools.vscode-docker)

Create `.vscode/settings.json`:
```json
{
  "go.lintOnSave": "package",
  "go.lintTool": "golangci-lint",
  "go.lintArgs": ["--fast"],
  "editor.formatOnSave": true,
  "[go]": {
    "editor.defaultFormatter": "golang.go",
    "editor.formatOnSave": true
  }
}
```

### GoLand / IntelliJ IDEA

- Built-in Go support
- Enable gofmt on save
- Configure run configurations for `cmd/server/main.go`

## Useful Commands

```bash
# Check Go version
go version

# List all dependencies
go list -m all

# Update dependencies
go get -u ./...

# View module graph
go mod graph

# Download specific version
go get github.com/package@v1.2.3

# Remove unused dependencies
go mod tidy
```

## Next Steps

1. Read [ARCHITECTURE.md](./ARCHITECTURE.md) untuk memahami struktur aplikasi
2. Read [API.md](./API.md) untuk dokumentasi API
3. Mulai develop dengan membuat feature baru
4. Ikuti [CONTRIBUTING.md](../CONTRIBUTING.md) untuk contribution guidelines

---

**Last Updated:** January 7, 2026
