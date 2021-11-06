#!/bin/bash

go run ./src/services/storage/  \
  --bind-address localhost:7000 \
  --clickhouse-address 0.0.0.0 \
  --clickhouse-user streamer \
  --clickhouse-database stream \
  --clickhouse-password 'password' \
  --clickhouse-alt-hosts '0.0.0.0:9001,0.0.0.0:9002,0.0.0.0:9003' \
  --debug