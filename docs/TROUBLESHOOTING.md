# Troubleshooting Guide

Panduan ini membantu mengatasi masalah umum yang mungkin Anda hadapi.

## Installation Issues

### Go Module Issues

**Problem:** `go mod download` gagal

**Solution:**
```bash
# Clear module cache
go clean -modcache

# Re-download dependencies
go mod download

# Verify dependencies
go mod verify

# Tidy dependencies
go mod tidy
```

### Dependency Conflicts

**Problem:** Version conflict antar dependencies

**Solution:**
```bash
# Update all dependencies
go get -u ./...

# Or update specific dependency
go get -u github.com/package@latest

# Tidy to remove unused
go mod tidy
```

## Database Issues

### Connection Failed

**Problem:** `connection refused` atau `dial tcp: lookup db: no such host`

**Solution:**

1. **Check if database is running:**
```bash
docker-compose ps db
```

2. **If not running, start it:**
```bash
docker-compose up -d db
```

3. **Verify connection string in .env:**
```env
DB_HOST=localhost  # or db if using docker-compose
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=videocall_db
```

4. **Test connection manually:**
```bash
psql -U postgres -h localhost -d videocall_db
```

### Migration Failed

**Problem:** Migration tidak berjalan atau error

**Solution:**

1. **Check migration status:**
```bash
psql -U postgres -d videocall_db -c "SELECT * FROM schema_migrations;"
```

2. **Reset database (development only):**
```bash
docker-compose down -v
docker-compose up -d db
go run cmd/server/main.go migrate
```

3. **Manual migration:**
```bash
# Connect to database
psql -U postgres -d videocall_db

# Run SQL manually
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Database Locked

**Problem:** `database is locked` error

**Solution:**

1. **Check active connections:**
```bash
psql -U postgres -d videocall_db -c "SELECT * FROM pg_stat_activity;"
```

2. **Kill idle connections:**
```bash
psql -U postgres -d videocall_db -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'videocall_db' AND pid <> pg_backend_pid();"
```

3. **Restart database:**
```bash
docker-compose restart db
```

## Application Issues

### Port Already in Use

**Problem:** `listen tcp :8080: bind: address already in use`

**Solution:**

1. **Find process using port:**
```bash
# macOS/Linux
lsof -i :8080

# Windows
netstat -ano | findstr :8080
```

2. **Kill process:**
```bash
# macOS/Linux
kill -9 <PID>

# Windows
taskkill /PID <PID> /F
```

3. **Or use different port:**
```bash
SERVER_PORT=8081 go run cmd/server/main.go
```

### Application Won't Start

**Problem:** Application crashes immediately

**Solution:**

1. **Check logs:**
```bash
go run cmd/server/main.go 2>&1 | head -50
```

2. **Enable debug logging:**
```bash
LOG_LEVEL=debug go run cmd/server/main.go
```

3. **Check configuration:**
```bash
# Verify .env file exists
ls -la .env

# Check environment variables
echo $DB_HOST
echo $SERVER_PORT
```

4. **Verify database connection:**
```bash
psql -U postgres -h localhost -d videocall_db -c "SELECT 1;"
```

### High Memory Usage

**Problem:** Application menggunakan memory terlalu banyak

**Solution:**

1. **Check memory usage:**
```bash
# macOS/Linux
ps aux | grep videocall

# Docker
docker stats videocall
```

2. **Profile memory:**
```bash
# Add to main.go
import _ "net/http/pprof"

// In init or main
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()

# Access profile
curl http://localhost:6060/debug/pprof/heap > heap.prof
go tool pprof heap.prof
```

3. **Optimize code:**
- Check for memory leaks
- Use connection pooling
- Implement caching

### Slow Performance

**Problem:** API response lambat

**Solution:**

1. **Check database queries:**
```bash
# Enable query logging
LOG_LEVEL=debug go run cmd/server/main.go

# Check slow queries
psql -U postgres -d videocall_db -c "SELECT * FROM pg_stat_statements ORDER BY mean_time DESC LIMIT 10;"
```

2. **Add database indexes:**
```sql
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_products_category ON products(category);
CREATE INDEX idx_invoices_user_id ON invoices(user_id);
```

3. **Optimize queries:**
- Use pagination
- Add WHERE clauses
- Use EXPLAIN ANALYZE

4. **Check network latency:**
```bash
ping localhost
```

## Docker Issues

### Container Won't Start

**Problem:** `docker-compose up` gagal

**Solution:**

1. **Check logs:**
```bash
docker-compose logs videocall
```

2. **Rebuild image:**
```bash
docker-compose build --no-cache
```

3. **Remove old containers:**
```bash
docker-compose down -v
docker-compose up -d
```

### Image Build Failed

**Problem:** `docker build` error

**Solution:**

1. **Check Dockerfile:**
```bash
# Verify Dockerfile exists
ls -la Dockerfile

# Check syntax
docker build --dry-run .
```

2. **Build with verbose output:**
```bash
docker build -t videocall:latest --progress=plain .
```

3. **Check base image:**
```bash
# Verify base image is available
docker pull golang:1.23.1
```

### Volume Issues

**Problem:** Data tidak persist atau permission denied

**Solution:**

1. **Check volumes:**
```bash
docker volume ls
docker volume inspect videocall_postgres_data
```

2. **Fix permissions:**
```bash
docker-compose down -v
docker-compose up -d
```

3. **Manual volume cleanup:**
```bash
docker volume rm videocall_postgres_data
```

## Testing Issues

### Tests Fail

**Problem:** Unit tests gagal

**Solution:**

1. **Run tests with verbose output:**
```bash
go test -v ./...
```

2. **Run specific test:**
```bash
go test -run TestName -v ./internal/app/users/...
```

3. **Check test database:**
```bash
# Ensure test database exists
createdb videocall_test_db
```

4. **Run with race detector:**
```bash
go test -race ./...
```

### Coverage Report Issues

**Problem:** Coverage report tidak generate

**Solution:**

```bash
# Generate coverage
go test -coverprofile=coverage.out ./...

# View coverage
go tool cover -html=coverage.out

# Check coverage percentage
go tool cover -func=coverage.out | tail -1
```

## API Issues

### 401 Unauthorized

**Problem:** Request dengan token gagal

**Solution:**

1. **Verify token:**
```bash
# Get token
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password"}'

# Use token
curl -H "Authorization: Bearer <token>" \
  http://localhost:8080/api/v1/users/profile
```

2. **Check JWT secret:**
```bash
# Verify JWT_SECRET in .env
echo $JWT_SECRET
```

3. **Check token expiration:**
```bash
# Decode token (online tool)
# https://jwt.io/

# Or use jq
echo $TOKEN | cut -d'.' -f2 | base64 -d | jq .
```

### 404 Not Found

**Problem:** Endpoint tidak ditemukan

**Solution:**

1. **Verify endpoint exists:**
```bash
curl -v http://localhost:8080/api/v1/products
```

2. **Check routing:**
```bash
# Verify routes registered in main.go
grep -r "router.GET\|router.POST" cmd/
```

3. **Check API version:**
```bash
# Verify API version in .env
echo $API_VERSION
```

### 500 Internal Server Error

**Problem:** Server error

**Solution:**

1. **Check server logs:**
```bash
docker-compose logs -f videocall
```

2. **Enable debug logging:**
```bash
LOG_LEVEL=debug docker-compose up
```

3. **Check database:**
```bash
docker-compose exec db psql -U postgres -d videocall_db -c "SELECT 1;"
```

## Git Issues

### Merge Conflicts

**Problem:** Git merge conflict

**Solution:**

```bash
# View conflicts
git status

# Edit conflicted files
nano <file>

# Mark as resolved
git add <file>

# Complete merge
git commit -m "Resolve merge conflicts"
```

### Accidentally Committed Secrets

**Problem:** Committed .env dengan secrets

**Solution:**

```bash
# Remove from git history
git rm --cached .env
git commit --amend -m "Remove .env"

# Add to .gitignore
echo ".env" >> .gitignore
git add .gitignore
git commit -m "Add .env to gitignore"

# Force push (jika belum di-push)
git push --force-with-lease
```

## Performance Profiling

### CPU Profiling

```bash
# Add to main.go
import _ "net/http/pprof"

# Profile CPU
curl http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.prof
go tool pprof cpu.prof
```

### Memory Profiling

```bash
curl http://localhost:6060/debug/pprof/heap > heap.prof
go tool pprof heap.prof
```

### Goroutine Profiling

```bash
curl http://localhost:6060/debug/pprof/goroutine > goroutine.prof
go tool pprof goroutine.prof
```

## Getting Help

1. **Check logs first**
2. **Search existing issues**
3. **Create detailed issue report:**
   - Error message
   - Steps to reproduce
   - Environment info
   - Logs/screenshots

4. **Contact maintainers**

---

**Last Updated:** January 7, 2026
