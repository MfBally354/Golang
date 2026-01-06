# 🐳 Docker Guide untuk Golang Project

Panduan lengkap menggunakan Docker untuk project Golang ini.

---

## 📁 File-file Docker

```
golang-tutorial/
├── Dockerfile              # Build image untuk API.go (Gin)
├── Dockerfile.backend      # Build image untuk backend.go
├── docker-compose.yml      # Orchestrasi multiple containers
├── .dockerignore          # File yang diabaikan saat build
├── Makefile               # Shortcut commands
└── DOCKER_GUIDE.md        # Panduan ini
```

---

## 🚀 Quick Start

### 1. Install Docker

**Windows & Mac:**
- Download [Docker Desktop](https://www.docker.com/products/docker-desktop)

**Linux:**
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
```

**Verifikasi instalasi:**
```bash
docker --version
docker-compose --version
```

### 2. Build & Run dengan Docker Compose

```bash
# Build images
docker-compose build

# Run containers
docker-compose up -d

# Cek status
docker-compose ps
```

### 3. Test API

```bash
# Test API (Gin)
curl http://localhost:8090/users

# Test Backend (net/http)
curl http://localhost:8091/
```

### 4. Stop Containers

```bash
docker-compose down
```

---

## 🛠️ Menggunakan Makefile

Makefile menyediakan shortcut untuk Docker commands.

```bash
# Lihat semua commands
make help

# Build images
make build

# Run containers
make run

# Stop containers
make stop

# Restart
make restart

# Lihat logs
make logs

# Lihat logs API saja
make logs-api

# Clean up semua
make clean

# Development mode (foreground)
make dev
```

---

## 📖 Penjelasan File Docker

### 1. `Dockerfile` (Multi-stage Build)

```dockerfile
# Stage 1: Build
FROM golang:1.21-alpine AS builder
# ... compile aplikasi

# Stage 2: Runtime
FROM alpine:latest
# ... copy binary saja
```

**Keuntungan Multi-stage:**
- ✅ Image size kecil (dari ~800MB jadi ~15MB)
- ✅ Hanya berisi binary, tidak ada Go compiler
- ✅ Lebih aman (less attack surface)
- ✅ Faster deployment

### 2. `docker-compose.yml`

Menjalankan multiple services sekaligus:
- **api**: Service Gin (port 8090)
- **backend**: Service net/http (port 8091)
- **postgres**: Database (optional)
- **redis**: Cache (optional)

### 3. `.dockerignore`

Sama seperti `.gitignore`, tapi untuk Docker build context.

File yang diabaikan:
- `.git/`
- `README.md`
- `*_test.go`
- `.env`
- dll

Ini membuat build process lebih cepat.

---

## 🎯 Docker Commands Penting

### Build & Run

```bash
# Build single image
docker build -t golang-api .

# Build dengan nama file spesifik
docker build -f Dockerfile.backend -t golang-backend .

# Run container
docker run -p 8090:8090 golang-api

# Run di background
docker run -d -p 8090:8090 --name my-api golang-api

# Run dengan environment variables
docker run -e GO_ENV=production -p 8090:8090 golang-api
```

### Management

```bash
# Lihat running containers
docker ps

# Lihat semua containers (including stopped)
docker ps -a

# Stop container
docker stop golang-api

# Remove container
docker rm golang-api

# Remove image
docker rmi golang-api
```

### Logs & Debug

```bash
# Lihat logs
docker logs golang-api

# Follow logs (real-time)
docker logs -f golang-api

# Lihat 100 baris terakhir
docker logs --tail 100 golang-api

# Masuk ke container
docker exec -it golang-api sh

# Run command di container
docker exec golang-api ps aux
```

### Cleanup

```bash
# Remove stopped containers
docker container prune

# Remove unused images
docker image prune

# Remove everything unused
docker system prune -a

# Remove dengan volumes
docker system prune -a --volumes
```

---

## 🔧 Docker Compose Commands

```bash
# Build services
docker-compose build

# Build specific service
docker-compose build api

# Run all services
docker-compose up

# Run di background
docker-compose up -d

# Run specific service
docker-compose up api

# Stop services
docker-compose down

# Stop dan remove volumes
docker-compose down -v

# Restart service
docker-compose restart api

# Lihat logs
docker-compose logs

# Follow logs
docker-compose logs -f

# Logs specific service
docker-compose logs -f api

# Scale service
docker-compose up -d --scale api=3

# Execute command in service
docker-compose exec api sh
```

---

## 🎨 Customization

### Mengubah Port

**Edit `docker-compose.yml`:**
```yaml
ports:
  - "9000:8090"  # localhost:9000 -> container:8090
```

### Menambah Environment Variables

**Buat file `.env`:**
```env
GO_ENV=production
DATABASE_URL=postgres://user:pass@postgres:5432/db
REDIS_URL=redis://redis:6379
API_KEY=your-secret-key
```

**Update `docker-compose.yml`:**
```yaml
services:
  api:
    env_file:
      - .env
```

### Menambah Volume (Hot Reload)

Untuk development dengan hot reload:

```yaml
services:
  api:
    volumes:
      - .:/app  # Mount source code
    command: go run API.go
```

---

## 🏗️ Best Practices

### 1. Gunakan Multi-stage Build
```dockerfile
FROM golang:1.21-alpine AS builder
# ... build
FROM alpine:latest
# ... runtime
```

### 2. Gunakan .dockerignore
```
.git
*.md
*_test.go
```

### 3. Set User (Security)
```dockerfile
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser
USER appuser
```

### 4. Health Check
```dockerfile
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget --quiet --tries=1 --spider http://localhost:8090/health || exit 1
```

### 5. Optimize Layer Caching
```dockerfile
# Copy go.mod dulu (jarang berubah)
COPY go.mod go.sum ./
RUN go mod download

# Baru copy source code
COPY . .
RUN go build
```

---

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Cek port yang dipakai
lsof -i :8090

# Kill process
kill -9 <PID>

# Atau ubah port di docker-compose.yml
```

### Cannot Connect to Database
```bash
# Cek container running
docker-compose ps

# Cek logs database
docker-compose logs postgres

# Cek network
docker network ls
docker network inspect golang-tutorial_golang-network
```

### Image Size Too Large
- ✅ Gunakan alpine base image
- ✅ Gunakan multi-stage build
- ✅ Hapus cache: `RUN go clean -cache`
- ✅ Build dengan `-ldflags="-w -s"`

### Build Fails
```bash
# Clean rebuild
docker-compose build --no-cache

# Cek .dockerignore
cat .dockerignore

# Build dengan verbose
docker-compose build --progress=plain
```

---

## 🚢 Production Deployment

### 1. Build untuk Production

```bash
# Build dengan tag version
docker build -t myregistry/golang-api:v1.0.0 .

# Push ke registry
docker push myregistry/golang-api:v1.0.0
```

### 2. Deploy ke Server

```bash
# Di server
docker pull myregistry/golang-api:v1.0.0
docker run -d -p 80:8090 myregistry/golang-api:v1.0.0
```

### 3. Dengan Docker Swarm

```bash
# Init swarm
docker swarm init

# Deploy stack
docker stack deploy -c docker-compose.yml golang-app
```

### 4. Dengan Kubernetes

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: golang-api
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: api
        image: myregistry/golang-api:v1.0.0
        ports:
        - containerPort: 8090
```

---

## 📊 Monitoring

### Docker Stats
```bash
# Real-time stats
docker stats

# Specific container
docker stats golang-api
```

### Resource Limits
```yaml
# docker-compose.yml
services:
  api:
    deploy:
      resources:
        limits:
          cpus: '0.5'
          memory: 512M
        reservations:
          cpus: '0.25'
          memory: 256M
```

---

## 🧪 Testing dengan Docker

```bash
# Run tests
docker-compose exec api go test ./...

# Run tests dengan coverage
docker-compose exec api go test -cover ./...

# Run benchmark
docker-compose exec api go test -bench=. ./...
```

---

## 📚 Resources

- [Docker Docs](https://docs.docker.com/)
- [Docker Compose Docs](https://docs.docker.com/compose/)
- [Go Docker Best Practices](https://docs.docker.com/language/golang/)
- [Multi-stage Builds](https://docs.docker.com/build/building/multi-stage/)

---

## 💡 Tips

1. **Gunakan Docker Compose** untuk development
2. **Gunakan Makefile** untuk shortcut commands
3. **Multi-stage build** untuk production
4. **Health checks** untuk monitoring
5. **Resource limits** untuk stability
6. **Volumes** untuk persistent data
7. **Networks** untuk isolasi
8. **Secrets** untuk sensitive data

---

## ✅ Checklist

```
□ Install Docker & Docker Compose
□ Buat Dockerfile
□ Buat docker-compose.yml
□ Buat .dockerignore
□ Test build locally
□ Test run locally
□ Setup Makefile
□ Add health checks
□ Setup monitoring
□ Push to registry
□ Deploy to production
```

---

**Happy Dockerizing! 🐳**
