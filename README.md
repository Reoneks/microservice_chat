# api-gateway

- Prometheus and Grafana usage:
  1. Launch run-local.sh in deploy/local directory
  2. Open [localhost:3000](localhost:3000)
  3. Go to Configuration -> Data Sources -> Add Data Source
  4. Choose `Prometheus`
  5. Type <http://localhost:9090> in `URL` field
  6. Push `Save & test` button
  7. Go to Dashboards -> Manage -> Import and type 3362 in `Import via grafana.com` field
  8. Push `Load` button
  9. Go to Dashboards -> Manage -> Import and type 1860 in `Import via grafana.com` field
  10. Push `Load` button
  11. Go to Dashboards -> Manage and choose required dashboard
