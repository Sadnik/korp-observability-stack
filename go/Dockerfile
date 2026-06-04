# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go files
COPY go/main.go .
COPY go/go.mod .

# Download dependencies without checksum verification
RUN go mod download -x || go get github.com/prometheus/client_golang@v1.17.0
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http-server-projeto-korp .

# Runtime stage
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/http-server-projeto-korp .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./http-server-projeto-korp"]
