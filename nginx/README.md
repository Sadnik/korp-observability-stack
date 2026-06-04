# NGINX configuration

This directory contains the NGINX proxy configuration used by the stack.

## Files

- `http-server-projeto-korp.conf` - NGINX server configuration

## Behavior

- Routes `/projeto-korp` to the Go service at `http://http-server-projeto-korp:8080`
- Exposes a simple root response for `/`
- Returns `ok` on `/health`
- Blocks `/metrics` on the proxy side

## Usage

The `docker-compose.yml` file mounts this directory into the NGINX container at `/etc/nginx/conf.d`.
