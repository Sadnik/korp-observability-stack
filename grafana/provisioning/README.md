# Grafana provisioning

This directory contains the provisioning configuration used by Grafana.

## Contents

- `datasources/` - datasource definitions
- `dashboards/` - dashboard provisioning and dashboard JSON files

## Usage

The Grafana container loads these files automatically at startup when mounted under `/etc/grafana/provisioning`.
