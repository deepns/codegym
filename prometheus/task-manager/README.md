# Local monitoring stack (task-manager + Prometheus + Grafana)

This folder contains a docker-compose setup to build and run the `task-manager` app together with Prometheus and Grafana. Prometheus is configured to scrape the app at `/metrics` on port `8080`.

Files added:
- `docker-compose.yml` — builds the app and runs Prometheus and Grafana
- `prometheus/prometheus.yml` — Prometheus scrape config (scrapes `task-manager:8080`)
- `grafana/provisioning/datasources/prometheus.yml` — Grafana provisioning to auto-add Prometheus as a datasource

Run the stack

1. From this directory run:

```bash
docker compose up --build
```

2. After containers are up:

- The app will be accessible on http://localhost:8080
- Prometheus UI: http://localhost:9090
- Grafana UI: http://localhost:3000 (default admin/admin)

Prometheus will scrape `http://task-manager:8080/metrics` (container DNS) and Grafana is provisioned with that Prometheus datasource.

Notes
- If you want to customize Grafana dashboards, add provisioning `dashboards` files under `grafana/provisioning` or add dashboards manually in the UI.
- If you already have an image registry and prefer not to build locally, change `task-manager` service `image`/`build` section in `docker-compose.yml`.
