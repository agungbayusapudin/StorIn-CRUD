# Contributing Guide

Terima kasih telah tertarik untuk berkontribusi pada proyek Videocall! Panduan ini akan membantu Anda memulai.

## Code of Conduct

Kami berkomitmen untuk menyediakan lingkungan yang welcoming dan inclusive. Semua kontributor diharapkan untuk:
- Menghormati pendapat dan perspektif yang berbeda
- Memberikan feedback yang konstruktif
- Fokus pada apa yang terbaik untuk komunitas

## Getting Started

### 1. Fork Repository

```bash
# Fork repository di GitHub
# Clone fork Anda
git clone https://github.com/your-username/videocall.git
cd videocall

# Add upstream remote
git remote add upstream https://github.com/original-owner/videocall.git
```

### 2. Create Feature Branch

```bash
# Update main branch
git fetch upstream
git checkout main
git merge upstream/main

# Create feature branch
git checkout -b feature/your-feature-name
```

### 3. Make Changes

Ikuti guidelines berikut:

#### Code Style
- Follow Go conventions
- Run `go fmt` sebelum commit
- Run `go vet` untuk check errors
- Use meaningful variable names
- Add comments untuk complex logic

#### Commit Messages
```
feat: add user authentication
fix: resolve database connection issue
docs: update API documentation
test: add unit tests for user service
refactor: improve error handling
```

Format: `<type>: <description>`

Types:
- `feat` - New feature
- `fix` - Bug fix
- `docs` - Documentation
- `test` - Tests
- `refactor` - Code refactoring
- `perf` - Performance improvement
- `chore` - Build, dependencies, etc.

### 4. Test Your Changes

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestUserLogin ./internal/app/users/...
```

### 5. Push Changes

```bash
git push origin feature/your-feature-name
```

### 6. Create Pull Request

1. Go to GitHub repository
2. Click "New Pull Request"
3. Select your branch
4. Fill PR template:

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Related Issues
Closes #123

## Testing
- [ ] Unit tests added
- [ ] Integration tests added
- [ ] Manual testing done

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex logic
- [ ] Documentation updated
- [ ] No new warnings generated
```

## Development Workflow

### Project Structure

```
videocall/
├── cmd/              # Application entry points
├── internal/         # Private application code
├── pkg/              # Public library code
├── docs/             # Documentation
├── scripts/          # Build & deployment scripts
└── tests/            # Integration tests
```

### Adding New Feature

#### 1. Create Domain Structure

```bash
mkdir -p internal/app/newfeature/{controller,service,repository,schema}
```

#### 2. Define Schema

```go
// internal/app/newfeature/schema/newfeature_schema.go
package schema

type NewFeature struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
}
```

#### 3. Create Repository Interface

```go
// internal/app/newfeature/repository/interface.go
package repository

import "context"

type NewFeatureRepository interface {
    Create(ctx context.Context, feature *schema.NewFeature) error
    GetByID(ctx context.Context, id string) (*schema.NewFeature, error)
    Update(ctx context.Context, feature *schema.NewFeature) error
    Delete(ctx context.Context, id string) error
}
```

#### 4. Implement Repository

```go
// internal/app/newfeature/repository/newfeature_repository.go
package repository

import (
    "context"
    "database/sql"
)

type newFeatureRepository struct {
    db *sql.DB
}

func NewNewFeatureRepository(db *sql.DB) NewFeatureRepository {
    return &newFeatureRepository{db: db}
}

func (r *newFeatureRepository) Create(ctx context.Context, feature *schema.NewFeature) error {
    // Implementation
    return nil
}
```

#### 5. Create Service Interface

```go
// internal/app/newfeature/service/interface.go
package service

import "context"

type NewFeatureService interface {
    Create(ctx context.Context, feature *schema.NewFeature) error
    GetByID(ctx context.Context, id string) (*schema.NewFeature, error)
    Update(ctx context.Context, feature *schema.NewFeature) error
    Delete(ctx context.Context, id string) error
}
```

#### 6. Implement Service

```go
// internal/app/newfeature/service/newfeature_service.go
package service

import (
    "context"
    "errors"
)

type newFeatureService struct {
    repo repository.NewFeatureRepository
}

func NewNewFeatureService(repo repository.NewFeatureRepository) NewFeatureService {
    return &newFeatureService{repo: repo}
}

func (s *newFeatureService) Create(ctx context.Context, feature *schema.NewFeature) error {
    if feature.Name == "" {
        return errors.New("name is required")
    }
    return s.repo.Create(ctx, feature)
}
```

#### 7. Create Controller

```go
// internal/app/newfeature/controller/newfeature_controller.go
package controller

import (
    "github.com/gin-gonic/gin"
)

type NewFeatureController struct {
    service service.NewFeatureService
}

func NewNewFeatureController(service service.NewFeatureService) *NewFeatureController {
    return &NewFeatureController{service: service}
}

func (c *NewFeatureController) Create(ctx *gin.Context) {
    // Implementation
}
```

#### 8. Add Routes

```go
// internal/app/newfeature/routers/newfeature_router.go
package routers

import (
    "github.com/gin-gonic/gin"
)

func SetupNewFeatureRoutes(router *gin.Engine, controller *controller.NewFeatureController) {
    group := router.Group("/api/v1/newfeature")
    {
        group.POST("", controller.Create)
        group.GET("/:id", controller.GetByID)
        group.PUT("/:id", controller.Update)
        group.DELETE("/:id", controller.Delete)
    }
}
```

#### 9. Write Tests

```go
// internal/app/newfeature/service/newfeature_service_test.go
package service

import (
    "context"
    "testing"
)

func TestCreate(t *testing.T) {
    // Test implementation
}
```

## Testing Guidelines

### Unit Tests

```bash
go test ./internal/app/users/service -v
```

### Integration Tests

```bash
go test -tags=integration ./...
```

### Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Minimum Coverage

- Aim for 80% code coverage
- Critical paths should have 100% coverage

## Documentation

### Code Comments

```go
// Create creates a new user in the database
func (s *userService) Create(ctx context.Context, user *schema.User) error {
    // Validate user data
    if err := validateUser(user); err != nil {
        return err
    }
    
    // Save to database
    return s.repo.Create(ctx, user)
}
```

### API Documentation

Update `docs/API.md` untuk setiap endpoint baru.

### README Updates

Update `README.md` jika ada perubahan signifikan.

## Performance Considerations

- Use database indexes untuk frequently queried fields
- Implement pagination untuk large datasets
- Cache frequently accessed data
- Use connection pooling
- Optimize queries

## Security Considerations

- Validate semua user input
- Use parameterized queries
- Hash passwords dengan bcrypt
- Implement rate limiting
- Use HTTPS di production
- Rotate secrets regularly

## Debugging Tips

### Enable Debug Logging

```bash
LOG_LEVEL=debug go run cmd/server/main.go
```

### Use Debugger

```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Run with debugger
dlv debug cmd/server/main.go
```

### Database Debugging

```bash
# Connect to database
psql -U postgres -d videocall_db

# View tables
\dt

# View table schema
\d users
```

## Common Issues

### Import Cycle

Jika terjadi import cycle, refactor code untuk menghindari circular dependencies.

### Test Failures

```bash
# Run specific test with verbose output
go test -v -run TestName ./...

# Run with race detector
go test -race ./...
```

### Database Issues

```bash
# Reset database
docker-compose down -v
docker-compose up -d db
go run cmd/server/main.go migrate
```

## Review Process

1. Maintainer akan review PR Anda
2. Mungkin ada feedback atau request untuk changes
3. Update PR berdasarkan feedback
4. Setelah approval, PR akan di-merge

## Questions?

- Buka issue untuk questions
- Diskusi di GitHub discussions
- Hubungi maintainers

---

**Last Updated:** January 7, 2026
