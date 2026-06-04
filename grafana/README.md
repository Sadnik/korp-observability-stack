# Grafana provisioning

This directory contains Grafana provisioning files for the Projeto Korp stack.

The provisioning setup ensures Grafana loads the Prometheus datasource and dashboards automatically.

## Structure

- `provisioning/datasources/prometheus.yml` - defines the Prometheus datasource
- `provisioning/dashboards/provider.yml` - configures static dashboard provisioning
- `provisioning/dashboards/dashboards/http-server-projeto-korp-dashboard.json` - JSON dashboard definition

## Usage

`docker-compose.yml` mounts this directory into the Grafana container at `/etc/grafana/provisioning`.
