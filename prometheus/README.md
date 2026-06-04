# Prometheus configuration

This directory contains the Prometheus scrape configuration for the Go service.

## Files

- `prometheus.yml` - scrape config targeting `http-server-projeto-korp:8080`

## Behavior

Prometheus scrapes the Go application metrics endpoint every 5 seconds.

## Usage

The `docker-compose.yml` file mounts this directory into the Prometheus container at `/etc/prometheus`.
