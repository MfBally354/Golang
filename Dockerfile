# Multi-stage build untuk ukuran image yang kecil

# Stage 1: Build
FROM golang:1.21-alpine AS builder

# Install git (diperlukan untuk go get)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build aplikasi (ganti API.go sesuai file yang ingin di-build)
# -ldflags="-w -s" untuk mengurangi ukuran binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main API.go

# Stage 2: Runtime
FROM alpine:latest

# Install ca-certificates untuk HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary dari builder stage
COPY --from=builder /app/main .

# Expose port (sesuaikan dengan port di aplikasi)
EXPOSE 8090

# Run aplikasi
CMD ["./main"]
