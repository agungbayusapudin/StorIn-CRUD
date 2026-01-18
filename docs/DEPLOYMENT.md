# Deployment Guide

## Deployment Strategies

### 1. Docker Compose (Development & Small Scale)

#### Prerequisites
- Docker & Docker Compose installed
- `.env` file configured

#### Steps

1. **Build Image**
```bash
docker-compose build
```

2. **Start Services**
```bash
docker-compose up -d
```

3. **Verify Services**
```bash
docker-compose ps
```

4. **View Logs**
```bash
docker-compose logs -f videocall
```

5. **Stop Services**
```bash
docker-compose down
```

#### Docker Compose File Structure

```yaml
version: '3.8'

services:
  videocall:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=db
      - DB_PORT=5432
    depends_on:
      - db
    networks:
      - videocall-network

  db:
    image: postgres:15
    environment:
      - POSTGRES_DB=videocall_db
      - POSTGRES_PASSWORD=postgres
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - videocall-network

volumes:
  postgres_data:

networks:
  videocall-network:
    driver: bridge
```

### 2. Kubernetes Deployment (Production Scale)

#### Prerequisites
- Kubernetes cluster (EKS, GKE, AKS, or local)
- kubectl configured
- Docker image pushed to registry

#### Deployment Manifest

**deployment.yaml:**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: videocall-app
  labels:
    app: videocall
spec:
  replicas: 3
  selector:
    matchLabels:
      app: videocall
  template:
    metadata:
      labels:
        app: videocall
    spec:
      containers:
      - name: videocall
        image: your-registry/videocall:latest
        ports:
        - containerPort: 8080
        env:
        - name: DB_HOST
          valueFrom:
            configMapKeyRef:
              name: videocall-config
              key: db_host
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: videocall-secrets
              key: db_password
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
```

**service.yaml:**
```yaml
apiVersion: v1
kind: Service
metadata:
  name: videocall-service
spec:
  selector:
    app: videocall
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: LoadBalancer
```

#### Deploy to Kubernetes

```bash
# Create namespace
kubectl create namespace videocall

# Create ConfigMap
kubectl create configmap videocall-config \
  --from-literal=db_host=postgres.videocall.svc.cluster.local \
  -n videocall

# Create Secret
kubectl create secret generic videocall-secrets \
  --from-literal=db_password=your_password \
  -n videocall

# Apply manifests
kubectl apply -f deployment.yaml -n videocall
kubectl apply -f service.yaml -n videocall

# Verify deployment
kubectl get pods -n videocall
kubectl get svc -n videocall
```

### 3. AWS Deployment

#### Using ECS (Elastic Container Service)

1. **Create ECR Repository**
```bash
aws ecr create-repository --repository-name videocall
```

2. **Push Image**
```bash
# Get login token
aws ecr get-login-password --region us-east-1 | \
  docker login --username AWS --password-stdin <account-id>.dkr.ecr.us-east-1.amazonaws.com

# Tag image
docker tag videocall:latest <account-id>.dkr.ecr.us-east-1.amazonaws.com/videocall:latest

# Push image
docker push <account-id>.dkr.ecr.us-east-1.amazonaws.com/videocall:latest
```

3. **Create ECS Task Definition**
```json
{
  "family": "videocall",
  "networkMode": "awsvpc",
  "requiresCompatibilities": ["FARGATE"],
  "cpu": "256",
  "memory": "512",
  "containerDefinitions": [
    {
      "name": "videocall",
      "image": "<account-id>.dkr.ecr.us-east-1.amazonaws.com/videocall:latest",
      "portMappings": [
        {
          "containerPort": 8080,
          "hostPort": 8080,
          "protocol": "tcp"
        }
      ],
      "environment": [
        {
          "name": "DB_HOST",
          "value": "your-rds-endpoint"
        }
      ],
      "secrets": [
        {
          "name": "DB_PASSWORD",
          "valueFrom": "arn:aws:secretsmanager:us-east-1:<account-id>:secret:videocall/db-password"
        }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/videocall",
          "awslogs-region": "us-east-1",
          "awslogs-stream-prefix": "ecs"
        }
      }
    }
  ]
}
```

4. **Create ECS Service**
```bash
aws ecs create-service \
  --cluster videocall-cluster \
  --service-name videocall-service \
  --task-definition videocall:1 \
  --desired-count 3 \
  --launch-type FARGATE \
  --network-configuration "awsvpcConfiguration={subnets=[subnet-xxx],securityGroups=[sg-xxx],assignPublicIp=ENABLED}"
```

#### Using Elastic Beanstalk

1. **Create Dockerrun.aws.json**
```json
{
  "AWSEBDockerrunVersion": 2,
  "containerDefinitions": [
    {
      "name": "videocall",
      "image": "<account-id>.dkr.ecr.us-east-1.amazonaws.com/videocall:latest",
      "essential": true,
      "memory": 512,
      "portMappings": [
        {
          "hostPort": 80,
          "containerPort": 8080
        }
      ],
      "environment": [
        {
          "name": "DB_HOST",
          "value": "your-rds-endpoint"
        }
      ]
    }
  ]
}
```

2. **Deploy**
```bash
eb create videocall-env
eb deploy
```

## Environment Configuration

### Development Environment

```env
ENV=development
SERVER_PORT=8080
LOG_LEVEL=debug
DB_SSL_MODE=disable
```

### Staging Environment

```env
ENV=staging
SERVER_PORT=8080
LOG_LEVEL=info
DB_SSL_MODE=require
```

### Production Environment

```env
ENV=production
SERVER_PORT=8080
LOG_LEVEL=warn
DB_SSL_MODE=require
```

## Database Migration

### Before Deployment

```bash
# Run migrations
go run cmd/server/main.go migrate

# Verify migrations
psql -U postgres -d videocall_db -c "\dt"
```

### Automated Migration

Add to Dockerfile:
```dockerfile
RUN go run cmd/server/main.go migrate
```

Or use init container in Kubernetes:
```yaml
initContainers:
- name: db-migrate
  image: your-registry/videocall:latest
  command: ["go", "run", "cmd/server/main.go", "migrate"]
```

## Monitoring & Logging

### CloudWatch (AWS)

```bash
# Create log group
aws logs create-log-group --log-group-name /videocall/app

# View logs
aws logs tail /videocall/app --follow
```

### Prometheus & Grafana

1. **Add Prometheus metrics endpoint**
```go
import "github.com/prometheus/client_golang/prometheus/promhttp"

router.GET("/metrics", gin.WrapH(promhttp.Handler()))
```

2. **Deploy Prometheus**
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: prometheus-config
data:
  prometheus.yml: |
    global:
      scrape_interval: 15s
    scrape_configs:
    - job_name: 'videocall'
      static_configs:
      - targets: ['localhost:8080']
```

## Scaling

### Horizontal Scaling

#### Docker Compose
```bash
docker-compose up -d --scale videocall=3
```

#### Kubernetes
```bash
kubectl scale deployment videocall-app --replicas=5 -n videocall
```

#### AWS ECS
```bash
aws ecs update-service \
  --cluster videocall-cluster \
  --service videocall-service \
  --desired-count 5
```

### Vertical Scaling

Increase resource limits in deployment configuration.

## Health Checks

### Liveness Probe
```bash
curl http://localhost:8080/health
```

### Readiness Probe
```bash
curl http://localhost:8080/ready
```

## Backup & Recovery

### Database Backup

```bash
# Backup
pg_dump -U postgres videocall_db > backup.sql

# Restore
psql -U postgres videocall_db < backup.sql
```

### Automated Backup (AWS RDS)

```bash
aws rds create-db-snapshot \
  --db-instance-identifier videocall-db \
  --db-snapshot-identifier videocall-backup-$(date +%Y%m%d)
```

## Rollback Strategy

### Docker Compose
```bash
# Rollback to previous image
docker-compose down
git checkout previous-version
docker-compose up -d
```

### Kubernetes
```bash
# View rollout history
kubectl rollout history deployment/videocall-app -n videocall

# Rollback to previous version
kubectl rollout undo deployment/videocall-app -n videocall
```

## Security Checklist

- [ ] Use HTTPS/TLS in production
- [ ] Rotate JWT secrets regularly
- [ ] Use environment variables for secrets
- [ ] Enable database encryption
- [ ] Configure firewall rules
- [ ] Enable logging and monitoring
- [ ] Regular security updates
- [ ] Use private container registry
- [ ] Implement rate limiting
- [ ] Enable CORS properly

## Troubleshooting

### Application won't start
```bash
# Check logs
docker-compose logs videocall

# Check database connection
docker-compose exec videocall curl http://db:5432
```

### High memory usage
```bash
# Check memory limits
docker stats

# Increase memory in docker-compose.yml
```

### Database connection timeout
```bash
# Verify database is running
docker-compose ps db

# Check connection string
echo $DB_HOST
```

---

**Last Updated:** January 7, 2026
