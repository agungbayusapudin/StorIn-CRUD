# Architecture Documentation

## Overview

Aplikasi Videocall menggunakan **Clean Architecture** dengan prinsip separation of concerns. Arsitektur ini memungkinkan aplikasi untuk scalable, testable, dan maintainable.

## Architecture Layers

### 1. Controller Layer (Presentation)
- Menangani HTTP requests dan responses
- Validasi input dari user
- Mengkonversi request ke format yang dimengerti service layer
- Lokasi: `internal/app/{domain}/controller/`

```
Request → Controller → Service
```

### 2. Service Layer (Business Logic)
- Mengimplementasikan business logic
- Orchestration antar repository
- Validasi business rules
- Lokasi: `internal/app/{domain}/service/`

```
Controller → Service → Repository
```

### 3. Repository Layer (Data Access)
- Abstraksi database operations
- Query building dan execution
- Data transformation
- Lokasi: `internal/app/{domain}/repository/`

```
Service → Repository → Database
```

### 4. Schema Layer (Data Models)
- Define struktur data
- Database models
- Request/Response DTOs
- Lokasi: `internal/app/{domain}/schema/`

## Domain Structure

Setiap domain (users, product, billing) memiliki struktur yang sama:

```
internal/app/{domain}/
├── controller/
│   └── {domain}_controller.go      # HTTP handlers
├── service/
│   ├── interface.go                # Service interface
│   └── {domain}_service.go         # Business logic
├── repository/
│   ├── interface.go                # Repository interface
│   └── {domain}_repository.go      # Data access
└── schema/
    └── {domain}_schema.go          # Data models
```

## Data Flow

### Create Product Flow

```
1. HTTP Request
   POST /api/v1/products
   {
     "name": "Product Name",
     "price": 99.99
   }

2. Controller Layer
   - Validate request
   - Parse JSON
   - Call service.CreateProduct()

3. Service Layer
   - Validate business rules
   - Check inventory
   - Call repository.Create()

4. Repository Layer
   - Build SQL query
   - Execute query
   - Return created product

5. Controller Layer
   - Format response
   - Return HTTP 201 Created
```

## Dependency Injection

Aplikasi menggunakan dependency injection untuk loose coupling:

```go
// Service depends on Repository interface
type ProductService struct {
    repo repository.ProductRepository
}

// Repository is injected
func NewProductService(repo repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}
```

## Error Handling

Error handling mengikuti pattern:

```
Repository Error → Service Error → Controller Error Response
```

Lokasi error handling: `errors/valid_error.go`

## Authentication & Authorization

### JWT Flow

```
1. User Login
   POST /api/v1/users/login
   
2. Generate JWT Token
   - User credentials validated
   - JWT token generated
   - Token returned to client

3. Subsequent Requests
   - Client sends token in Authorization header
   - Middleware validates token
   - Request processed if valid
```

### Middleware

Lokasi: `middleware/auth_middelware.go`

- Validate JWT token
- Extract user information
- Attach user to request context

## Database Design

### Entity Relationship

```
Users (1) ──── (N) Invoices
  │
  └──── (N) Products (through Invoices)

Products (1) ──── (N) Invoices
```

### Database Connection

Lokasi: `db/db.go` dan `db/configurationDB.go`

- Connection pooling
- Configuration management
- Migration support

## Scalability Considerations

### Current (Monolith)
- Single binary
- Shared database
- All services in one process

### Future (Microservices)
- Separate services per domain
- Service-specific databases
- API Gateway for routing
- Message queue for async communication

## Configuration Management

Konfigurasi dikelola melalui:

1. **Environment Variables** (`.env`)
   - Development configuration
   - Secrets management

2. **Config Files** (`configs/`)
   - YAML-based configuration
   - Environment-specific configs

## Testing Strategy

### Unit Tests
- Test individual functions
- Mock dependencies
- Located alongside source code (`*_test.go`)

### Integration Tests
- Test multiple layers
- Use test database
- Located in `test/` directory

### Test Coverage
```bash
go test -cover ./...
```

## Deployment Architecture

### Development
```
Docker Compose
├── videocall-app (Go application)
├── PostgreSQL (Database)
└── Redis (Cache - optional)
```

### Production
```
Kubernetes / Docker Swarm
├── videocall-app (multiple replicas)
├── PostgreSQL (managed service)
├── Redis (managed service)
├── Load Balancer
└── Monitoring (Prometheus, Grafana)
```

## Performance Optimization

### Database
- Connection pooling
- Query optimization
- Indexing strategy

### Caching
- Redis for session management
- Cache frequently accessed data

### API
- Pagination for list endpoints
- Rate limiting
- Request timeout

## Security Architecture

### Authentication
- JWT-based authentication
- Token expiration
- Refresh token mechanism

### Authorization
- Role-based access control (RBAC)
- Resource-level permissions

### Data Protection
- Password hashing (bcrypt)
- SQL injection prevention (parameterized queries)
- CORS configuration

## Monitoring & Logging

### Logging
- Structured logging (JSON format)
- Log levels (debug, info, warn, error)
- Centralized log aggregation

### Metrics
- Application metrics (requests, latency)
- Database metrics (connections, queries)
- System metrics (CPU, memory)

### Health Checks
- `/health` endpoint
- Database connectivity check
- Service availability check

## Future Improvements

1. **Event-Driven Architecture**
   - Implement event sourcing
   - Message queue integration

2. **Microservices**
   - Split into independent services
   - Service discovery
   - API Gateway

3. **Advanced Features**
   - Real-time notifications (WebSocket)
   - Video streaming optimization
   - Advanced analytics

4. **Infrastructure**
   - Kubernetes deployment
   - Service mesh (Istio)
   - Advanced monitoring

---

**Last Updated:** January 7, 2026
