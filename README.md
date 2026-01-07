# Videocall Application

Aplikasi videocall adalah platform komunikasi real-time yang memungkinkan pengguna untuk melakukan panggilan video, mengelola produk, dan mengelola billing. Dibangun dengan Go menggunakan arsitektur clean architecture dan microservice-ready.

## 📋 Daftar Isi

- [Fitur](#fitur)
- [Tech Stack](#tech-stack)
- [Prasyarat](#prasyarat)
- [Instalasi](#instalasi)
- [Konfigurasi](#konfigurasi)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [Struktur Proyek](#struktur-proyek)
- [API Documentation](#api-documentation)
- [Database](#database)
- [Testing](#testing)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

## ✨ Fitur

### User Management
- Registrasi dan login pengguna
- Manajemen profil pengguna
- Authentication & Authorization
- Role-based access control

### Product Management
- CRUD operasi untuk produk
- Kategori produk
- Inventory management
- Product listing

### Billing System
- Invoice generation
- Payment tracking
- Billing history
- Transaction management

### Video Call
- Real-time video communication
- Call history
- Call recording (optional)
- Call notifications

## 🛠 Tech Stack

| Komponen | Teknologi |
|----------|-----------|
| **Language** | Go 1.23.1 |
| **Framework** | Gin (HTTP) |
| **Database** | PostgreSQL / MySQL |
| **Authentication** | JWT |
| **Architecture** | Clean Architecture |
| **Containerization** | Docker |
| **Orchestration** | Docker Compose |

## 📦 Prasyarat

- Go 1.23.1 atau lebih tinggi
- Docker & Docker Compose
- PostgreSQL 12+ atau MySQL 8+
- Git

## 🚀 Instalasi

### 1. Clone Repository
```bash
git clone <repository-url>
cd videocall
```

### 2. Install Dependencies
```bash
go mod download
go mod tidy
```

### 3. Setup Environment
```bash
cp .env.example .env
# Edit .env dengan konfigurasi lokal Anda
```

### 4. Setup Database
```bash
# Menggunakan Docker Compose
docker-compose up -d db

# Atau manual migration
go run cmd/server/main.go migrate
```

## ⚙️ Konfigurasi

### Environment Variables

Buat file `.env` berdasarkan `.env.example`:

```env
# Server Configuration
SERVER_PORT=8080
SERVER_HOST=localhost
ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=videocall_db
DB_SSL_MODE=disable

# JWT Configuration
JWT_SECRET=your_secret_key_here
JWT_EXPIRATION=24h

# API Configuration
API_VERSION=v1
API_TIMEOUT=30s

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

### Konfigurasi Database

Edit file `configs/config.yaml`:

```yaml
database:
  driver: postgres
  host: localhost
  port: 5432
  user: postgres
  password: password
  name: videocall_db
  max_connections: 25
  idle_connections: 5

server:
  port: 8080
  timeout: 30s
  
jwt:
  secret: your_secret_key
  expiration: 24h
```

## 🏃 Menjalankan Aplikasi

### Development Mode

```bash
# Run dengan hot reload (gunakan air atau similar)
go run cmd/server/main.go

# Atau dengan make
make run
```

### Production Mode

```bash
# Build binary
go build -o bin/videocall cmd/server/main.go

# Run binary
./bin/videocall
```

### Docker

```bash
# Build image
docker build -t videocall:latest .

# Run container
docker run -p 8080:8080 --env-file .env videocall:latest

# Atau dengan docker-compose
docker-compose up -d
```

## 📁 Struktur Proyek

```
videocall/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point aplikasi
├── internal/
│   ├── app/
│   │   ├── users/                  # User service domain
│   │   │   ├── controller/         # HTTP handlers
│   │   │   ├── service/            # Business logic
│   │   │   ├── repository/         # Data access layer
│   │   │   └── schema/             # Data models
│   │   ├── product/                # Product service domain
│   │   │   ├── controller/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   ├── routers/
│   │   │   └── schema/
│   │   └── billing/                # Billing service domain
│   │       ├── controller/
│   │       ├── service/
│   │       ├── repository/
│   │       └── schema/
│   └── pkg/
│       └── handler.go              # Shared utilities
├── pkg/                            # Public library code
├── middleware/
│   └── auth_middelware.go          # Authentication middleware
├── db/
│   ├── db.go                       # Database connection
│   └── configurationDB.go          # Database configuration
├── errors/
│   └── valid_error.go              # Error handling
├── dto/
│   └── dto.go                      # Data transfer objects
├── configs/
│   ├── config.yaml                 # Configuration template
│   └── development.yaml            # Development config
├── scripts/
│   ├── build.sh                    # Build script
│   ├── deploy.sh                   # Deploy script
│   └── migrate.sh                  # Database migration
├── api/
│   └── openapi/
│       └── videocall.yaml          # API documentation
├── deployments/
│   ├── docker/
│   │   └── Dockerfile
│   └── docker-compose.yml
├── go.mod                          # Go module definition
├── go.sum                          # Dependency lock file
├── .env.example                    # Environment template
├── .gitignore                      # Git ignore rules
├── Dockerfile                      # Docker configuration
├── docker-compose.yml              # Docker compose config
└── README.md                       # This file
```

## 📚 API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication
Semua endpoint (kecuali login/register) memerlukan JWT token di header:
```
Authorization: Bearer <token>
```

### User Endpoints

#### Register User
```http
POST /users/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe"
}
```

#### Login
```http
POST /users/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}

Response:
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

#### Get User Profile
```http
GET /users/profile
Authorization: Bearer <token>
```

#### Update User Profile
```http
PUT /users/profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Jane Doe",
  "phone": "+62812345678"
}
```

### Product Endpoints

#### Get All Products
```http
GET /products?page=1&limit=10
Authorization: Bearer <token>
```

#### Get Product by ID
```http
GET /products/:id
Authorization: Bearer <token>
```

#### Create Product
```http
POST /products
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99,
  "category": "electronics"
}
```

#### Update Product
```http
PUT /products/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Updated Name",
  "price": 149.99
}
```

#### Delete Product
```http
DELETE /products/:id
Authorization: Bearer <token>
```

### Billing Endpoints

#### Get Invoices
```http
GET /billing/invoices
Authorization: Bearer <token>
```

#### Get Invoice by ID
```http
GET /billing/invoices/:id
Authorization: Bearer <token>
```

#### Create Invoice
```http
POST /billing/invoices
Authorization: Bearer <token>
Content-Type: application/json

{
  "user_id": "uuid",
  "items": [
    {
      "product_id": "uuid",
      "quantity": 2,
      "price": 99.99
    }
  ]
}
```

#### Pay Invoice
```http
POST /billing/invoices/:id/pay
Authorization: Bearer <token>
Content-Type: application/json

{
  "payment_method": "credit_card",
  "amount": 199.98
}
```

## 🗄️ Database

### Schema

#### Users Table
```sql
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  phone VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Products Table
```sql
CREATE TABLE products (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  price DECIMAL(10, 2) NOT NULL,
  category VARCHAR(100),
  stock INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Invoices Table
```sql
CREATE TABLE invoices (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id),
  total_amount DECIMAL(10, 2) NOT NULL,
  status VARCHAR(50) DEFAULT 'pending',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Migration

```bash
# Run migrations
go run cmd/server/main.go migrate

# Rollback migrations
go run cmd/server/main.go migrate:rollback
```

## 🧪 Testing

### Unit Tests
```bash
# Run all tests
go test ./...

# Run tests dengan coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Integration Tests
```bash
# Run integration tests
go test -tags=integration ./...
```

### Load Testing
```bash
# Menggunakan Apache Bench
ab -n 1000 -c 10 http://localhost:8080/api/v1/products

# Atau menggunakan wrk
wrk -t12 -c400 -d30s http://localhost:8080/api/v1/products
```

## 🐳 Deployment

### Docker Compose (Development)
```bash
docker-compose up -d
```

### Docker Compose (Production)
```bash
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

### Kubernetes (Optional)
```bash
# Build image
docker build -t videocall:latest .

# Push ke registry
docker tag videocall:latest your-registry/videocall:latest
docker push your-registry/videocall:latest

# Deploy ke Kubernetes
kubectl apply -f deployments/kubernetes/
```

### Environment-specific Deployment

**Development:**
```bash
make deploy-dev
```

**Staging:**
```bash
make deploy-staging
```

**Production:**
```bash
make deploy-prod
```

## 📊 Monitoring & Logging

### Logs
```bash
# View logs
docker-compose logs -f videocall

# View specific service logs
docker-compose logs -f db
```

### Health Check
```bash
curl http://localhost:8080/health
```

### Metrics
```bash
curl http://localhost:8080/metrics
```

## 🔒 Security

### Best Practices
- Selalu gunakan HTTPS di production
- Jangan commit `.env` file dengan secrets
- Rotate JWT secret secara berkala
- Implement rate limiting
- Validate semua input dari user
- Use parameterized queries untuk prevent SQL injection

### CORS Configuration
```go
// Configured in main.go
cors.DefaultConfig()
```

## 🤝 Contributing

1. Fork repository
2. Buat feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

### Code Style
- Follow Go conventions
- Run `go fmt` sebelum commit
- Run `go vet` untuk check errors
- Write tests untuk setiap feature

## 📝 License

Project ini dilisensikan di bawah MIT License - lihat file LICENSE untuk detail.

## 📞 Support

Untuk pertanyaan atau issues, silakan buat issue di repository atau hubungi tim development.

## 🗺️ Roadmap

- [ ] Real-time video call implementation
- [ ] Call recording feature
- [ ] Advanced billing reports
- [ ] Mobile app integration
- [ ] Payment gateway integration
- [ ] Analytics dashboard
- [ ] Multi-language support
- [ ] Two-factor authentication

---

**Last Updated:** January 7, 2026  
**Version:** 1.0.0  
**Maintainer:** Development Team
