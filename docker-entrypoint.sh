#!/bin/sh
set -e

CONFIG_PATH="${CONFIG_PATH:-/app/config/config.yml}"
CONFIG_DIR=$(dirname "$CONFIG_PATH")

mkdir -p "$CONFIG_DIR"

cat > "$CONFIG_PATH" <<EOF
App:
  name: "${APP_NAME:-CurrencyExchangeApp}"
  port: "${APP_PORT:-:3000}"

Database:
  dsn: "${DB_DSN:-root:password@tcp(mysql:3306)/test?charset=utf8mb4&parseTime=True&loc=Local}"
  MaxIdleConns: ${DB_MAX_IDLE_CONNS:-10}
  MaxOpenCons: ${DB_MAX_OPEN_CONNS:-100}

Redis:
  addr: "${REDIS_ADDR:-redis:6379}"
  password: "${REDIS_PASSWORD:-}"
  db: ${REDIS_DB:-0}
EOF

exec /app/exchangeapp
