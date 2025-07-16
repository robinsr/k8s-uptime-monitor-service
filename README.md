# Uptime Monitor

A lightweight uptime monitoring service written in Go, designed for deployment on Kubernetes.

## Features

- Periodically checks the availability of configured HTTP endpoints.
- Records uptime status, response durations, and failure counts.
- Exposes Prometheus metrics at `/metrics` for observability.
- Simple health check endpoint at `/healthz`.
- Configurable via YAML file specifying targets, intervals, and timeouts.

## Usage

1. Define monitoring targets in `config/service-targets.yaml`.
2. Run the service locally:
   ```
   go run ./cmd/main.go
   ```
3. Access metrics:
   - [http://localhost:8080/metrics](http://localhost:8080/metrics)
   - [http://localhost:8080/healthz](http://localhost:8080/healthz)

## Project Status

This project is an in-progress personal portfolio project intended to demonstrate proficiency with:

- Go programming
- Prometheus metrics
- Kubernetes-ready application design

Further enhancements are planned.

---

Developed by Ryan Robinson