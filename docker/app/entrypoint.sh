#!/bin/sh
export POSTGRESQL_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_USER}?sslmode=disable

# migration
migrate -database ${POSTGRESQL_URL} -path docker/db/migrations up

# コンテナを起動し続けるためにshを実行
/bin/sh