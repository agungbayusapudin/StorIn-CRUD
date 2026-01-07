# Documentation Index

Dokumentasi lengkap untuk aplikasi Videocall. Pilih topik yang ingin Anda pelajari:

## 📚 Getting Started

- **[README.md](../README.md)** - Overview aplikasi, fitur, dan quick start
- **[SETUP.md](./SETUP.md)** - Panduan instalasi dan setup development environment

## 🏗️ Architecture & Design

- **[ARCHITECTURE.md](./ARCHITECTURE.md)** - Penjelasan arsitektur clean architecture
  - Layer structure
  - Data flow
  - Dependency injection
  - Database design
  - Scalability considerations

## 📖 API Documentation

- **[API.md](./API.md)** - Dokumentasi lengkap semua endpoint
  - User endpoints
  - Product endpoints
  - Billing endpoints
  - Authentication
  - Error codes
  - Rate limiting

## 🚀 Deployment

- **[DEPLOYMENT.md](./DEPLOYMENT.md)** - Panduan deployment ke berbagai platform
  - Docker Compose
  - Kubernetes
  - AWS (ECS, Elastic Beanstalk)
  - Environment configuration
  - Database migration
  - Monitoring & logging
  - Scaling strategies

## 🤝 Contributing

- **[CONTRIBUTING.md](../CONTRIBUTING.md)** - Panduan untuk berkontribusi
  - Code style
  - Commit messages
  - Testing guidelines
  - Pull request process
  - Adding new features

## 🔧 Troubleshooting

- **[TROUBLESHOOTING.md](./TROUBLESHOOTING.md)** - Solusi untuk masalah umum
  - Installation issues
  - Database issues
  - Application issues
  - Docker issues
  - Testing issues
  - API issues
  - Performance profiling

## 📋 Project Files

- **[README.md](../README.md)** - Main project documentation
- **[CHANGELOG.md](../CHANGELOG.md)** - Version history dan changes
- **[LICENSE](../LICENSE)** - MIT License
- **[.env.example](../.env.example)** - Environment variables template
- **[.gitignore](../.gitignore)** - Git ignore rules

## 🗂️ Project Structure

```
videocall/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point
├── internal/
│   ├── app/
│   │   ├── users/                  # User domain
│   │   ├── product/                # Product domain
│   │   └── billing/                # Billing domain
│   └── pkg/
│       └── handler.go              # Shared utilities
├── pkg/                            # Public library code
├── middleware/
│   └── auth_middelware.go          # Authentication
├── db/
│   ├── db.go                       # Database connection
│   └── configurationDB.go          # Configuration
├── errors/
│   └── valid_error.go              # Error handling
├── dto/
│   └── dto.go                      # Data transfer objects
├── configs/
│   └── config.yaml                 # Configuration
├── scripts/
│   ├── build.sh
│   ├── deploy.sh
│   └── migrate.sh
├── api/
│   └── openapi/
│       └── videocall.yaml          # API spec
├── deployments/
│   ├── docker/
│   │   └── Dockerfile
│   └── docker-compose.yml
├── docs/
│   ├── ARCHITECTURE.md
│   ├── SETUP.md
│   ├── API.md
│   ├── DEPLOYMENT.md
│   ├── TROUBLESHOOTING.md
│   └── INDEX.md                    # This file
├── go.mod
├── go.sum
├── .env.example
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── README.md
├── CONTRIBUTING.md
├── CHANGELOG.md
└── LICENSE
```

## 🎯 Quick Links

### For Developers
1. Start with [SETUP.md](./SETUP.md) untuk setup environment
2. Read [ARCHITECTURE.md](./ARCHITECTURE.md) untuk memahami struktur
3. Check [API.md](./API.md) untuk dokumentasi API
4. Follow [CONTRIBUTING.md](../CONTRIBUTING.md) untuk kontribusi

### For DevOps/Deployment
1. Read [DEPLOYMENT.md](./DEPLOYMENT.md) untuk deployment options
2. Check [TROUBLESHOOTING.md](./TROUBLESHOOTING.md) untuk common issues
3. Review [ARCHITECTURE.md](./ARCHITECTURE.md) untuk scalability

### For API Users
1. Check [API.md](./API.md) untuk endpoint documentation
2. Review authentication section untuk JWT setup
3. Check error codes untuk error handling

### For Contributors
1. Read [CONTRIBUTING.md](../CONTRIBUTING.md)
2. Follow code style guidelines
3. Write tests untuk setiap feature
4. Update documentation

## 📞 Support

- **Issues:** Buat issue di GitHub
- **Discussions:** Gunakan GitHub discussions
- **Email:** Hubungi maintainers

## 🔄 Documentation Updates

Dokumentasi di-update secara berkala. Untuk versi terbaru, selalu check repository.

**Last Updated:** January 7, 2026

---

## Navigation

- [← Back to README](../README.md)
- [Architecture →](./ARCHITECTURE.md)
- [Setup Guide →](./SETUP.md)
- [API Documentation →](./API.md)
- [Deployment →](./DEPLOYMENT.md)
- [Troubleshooting →](./TROUBLESHOOTING.md)
