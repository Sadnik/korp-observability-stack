# Projeto Korp

Production-style observability platform built with Go, Nginx, Prometheus, Grafana, Docker Compose, and Ansible.

The project demonstrates containerized service deployment, infrastructure automation, metrics collection, dashboard provisioning, and performance validation under sustained load.

Load-tested at **101,219 requests/sec** with **0 failed requests** while maintaining **15.05ms P99 latency** at **400 concurrent connections**.

## Screenshots

### Grafana Dashboard
![Grafana Dashboard](https://github.com/user-attachments/assets/6e019ed7-1316-44bf-b1b9-352b20a492b9)

### Prometheus Targets
![Prometheus Targets](https://github.com/user-attachments/assets/990c2b4f-0101-459b-a047-a172a24e043d)

### Prometheus Request Rate
![Prometheus Request Rate](https://github.com/user-attachments/assets/6cd38172-94d9-40e7-b1d8-f95de38a9a54)


---

## Architecture

```text
                           +----------------+
                           |     Client     |
                           +--------+-------+
                                    |
                                    v
                           +----------------+
                           |     Nginx      |
                           |      :80       |
                           +--------+-------+
                                    |
                                    v
                           +----------------+
                           |   Go Service   |
                           |     :8080      |
                           +--------+-------+
                                    |
                    +---------------+---------------+
                    |                               |
                    v                               v
           +----------------+              +----------------+
           |  Prometheus    | -----------> |    Grafana     |
           |     :9090      |              |     :3000      |
           +----------------+              +----------------+
```

All services run on an isolated Docker bridge network.

Nginx serves as the public entry point and reverse proxy. The Go application exposes Prometheus metrics, which are collected by Prometheus and visualized through Grafana dashboards provisioned as code.

---

## Stack

| Component      | Purpose                                 |
| -------------- | --------------------------------------- |
| Go 1.21        | HTTP API and Prometheus instrumentation |
| Nginx          | Reverse proxy and request routing       |
| Prometheus     | Metrics collection and storage          |
| Grafana        | Metrics visualization and dashboards    |
| Docker Compose | Container orchestration                 |
| Ansible        | Automated deployment and validation     |

---

## Quick Start

### Start the stack

```bash
docker compose up -d --build
```

### Verify service availability

```bash
curl http://localhost/projeto-korp
```

### Access monitoring tools

| Service     | URL                           |
| ----------- | ----------------------------- |
| Application | http://localhost/projeto-korp |
| Prometheus  | http://localhost:9090         |
| Grafana     | http://localhost:3000         |

Default Grafana credentials:

```text
admin / admin
```

---

## API Endpoints

| Endpoint          | Description                                         |
| ----------------- | --------------------------------------------------- |
| GET /projeto-korp | Returns JSON response containing name and timestamp |
| GET /health       | Nginx health check endpoint                         |
| GET /metrics      | Internal Prometheus metrics endpoint                |
| :9090             | Prometheus UI                                       |
| :3000             | Grafana UI                                          |

---


## Ansible Deployment

The Ansible playbook automates:

* Docker installation
* Container runtime validation
* Project deployment
* Service startup
* Health verification

Designed to be idempotent and safely re-runnable.

```bash
ansible-playbook ansible/playbook.yml
```

This works from the repository root because `ansible.cfg` now contains the default inventory, local connection, and privilege escalation settings.

---

## Performance Benchmark

Benchmark executed using:

```bash
wrk -t16 -c400 -d30s --latency http://localhost/projeto-korp
```

### Results

| Metric                 | Value     |
| ---------------------- | --------- |
| Requests/sec           | 101,219   |
| Total Requests         | 3,046,709 |
| Failed Requests        | 0         |
| P50 Latency            | 3.62ms    |
| P75 Latency            | 5.81ms    |
| P90 Latency            | 8.54ms    |
| P99 Latency            | 15.05ms   |
| Max Latency            | 70.99ms   |
| Concurrent Connections | 400       |
| Test Duration          | 30s       |

---

## Metrics

The application exports custom and runtime metrics through Prometheus.

| Metric                        | Type      | Description                                   |
| ----------------------------- | --------- | --------------------------------------------- |
| http_requests_total           | Counter   | Request count by method, endpoint, and status |
| http_request_duration_seconds | Histogram | Request latency distribution                  |
| service_availability          | Gauge     | Service availability state                    |


Observed metrics are visualized through Grafana dashboards, including:

* P95 latency
* P99 latency
* Service availability
* Request Rate (5m)
* Total Requests

---

## Design Decisions

### Reverse Proxy Layer

Nginx is the only public-facing component. The Go service is not exposed externally and communicates exclusively through the internal Docker network.

### Metrics Isolation

Prometheus scrapes metrics directly from the Go service. External access to `/metrics` is blocked at the proxy layer to reduce unnecessary exposure.

### Health Check Separation

The `/health` endpoint is served directly by Nginx. This allows validation of the proxy layer independently from the application backend.

### Infrastructure as Code

Application deployment, monitoring configuration, and service provisioning are version-controlled and reproducible.

### Grafana Provisioning

Datasources and dashboards are automatically loaded during startup, eliminating manual configuration steps.

### Multi-Stage Container Builds

The Go application is compiled in a dedicated build stage and deployed into a minimal runtime image to reduce attack surface and image size.

### Explicit HTTP Timeouts

Read, write, and idle timeouts are configured to prevent stalled connections and improve resilience under sustained load.

### Observability-First Design

Performance metrics, latency distributions, availability indicators, and runtime statistics are available immediately after deployment without additional configuration.
