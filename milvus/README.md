# Milvus + Prometheus + Grafana Monitoring Stack

This docker-compose setup brings up Milvus (v2.6.6) with its dependencies (etcd, minio) along with Prometheus and Grafana to monitor Milvus metrics.

## Services

- **Milvus** (v2.6.6): Vector database server with metrics endpoint on port 9091
- **etcd**: Configuration and coordination backend for Milvus
- **minio**: Object storage backend for Milvus
- **Prometheus**: Scrapes Milvus metrics at `http://milvus:9091/metrics` every 15 seconds
- **Grafana**: Visualization layer with Prometheus as the default datasource (admin/admin)

## Quick Start

1. From this directory (`milvus/`), start the stack:

```bash
docker compose up -d
```

2. Wait for services to be healthy (~30 seconds):

```bash
docker compose ps
```

3. Access the services:

- **Milvus gRPC**: localhost:19530
- **Milvus Metrics**: http://localhost:9091/metrics
- **Prometheus UI**: http://localhost:9090
- **Grafana UI**: http://localhost:3000 (admin / admin)

## Verify Prometheus is scraping Milvus

Open http://localhost:9090 → Status → Targets and confirm the `milvus` target is UP.

Or query via API:

```bash
curl -s 'http://localhost:9090/api/v1/targets' | jq '.data.activeTargets[] | {job: .labels.job, health}'
```

## Explore Milvus Metrics

1. Open Prometheus UI (http://localhost:9090) and search for metrics starting with `milvus_`:
   - `milvus_system_startup_millis_total`
   - `milvus_grpc_server_request_count`
   - `milvus_proxy_node_execute_latency_bucket`
   - etc.

2. Or in Grafana (http://localhost:3000):
   - Click Explore → select Prometheus datasource
   - Type `milvus_` to see available metrics
   - Build custom panels to visualize metrics

## Files

- `docker-compose.yml` — Milvus stack with Prometheus and Grafana
- `prometheus/prometheus.yml` — Prometheus scrape config (targets Milvus)
- `grafana/provisioning/datasources/prometheus.yml` — Grafana provisioning (Prometheus datasource)

## Troubleshooting

### Prometheus target DOWN

Check Milvus logs and healthcheck:

```bash
docker compose logs milvus
docker compose ps milvus
```

Verify Milvus metrics endpoint is reachable:

```bash
curl -s http://localhost:9091/metrics | head -20
```

### Grafana doesn't show datasource

Check Grafana logs:

```bash
docker compose logs grafana
```

Ensure `grafana/provisioning/datasources/prometheus.yml` exists and is readable.

## Stop the stack

```bash
docker compose down
```

To remove volumes (including Milvus data):

```bash
docker compose down -v
```
