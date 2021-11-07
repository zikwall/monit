#!/bin/bash

go run ./src/services/storage/  \
  --bind-address localhost:7000 \
  --clickhouse-address 194.35.48.20 \
  --clickhouse-user default \
  --clickhouse-database default \
  --debug