#!/bin/bash

go run ./src/services/api/  \
  --bind-address localhost:3001 \
  --storage-address 0.0.0.0:7000 \
  --repository-address 0.0.0.0:7001 \
  --monitoring-address 0.0.0.0:7002 \
  --debug