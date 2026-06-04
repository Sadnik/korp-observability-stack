# Go Application - Projeto Korp

This directory contains the Go HTTP server used by Projeto Korp.

## Files

- `main.go` - Go application source
- `Dockerfile` - multi-stage Docker build for the Go service
- `go.mod` - Go module definition

## Endpoints

- `GET /projeto-korp` - returns JSON with `nome` and `horario`
- `GET /health` - returns health status
- `GET /metrics` - exposes Prometheus metrics

## Local development

Run locally:

```bash
go run main.go
```

Build the Docker image:

```bash
docker build -t http-server-projeto-korp .
```
