export POSTGRESQL_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_USER}?sslmode=disable

migrate -database ${POSTGRESQL_URL} -path db/migrations up 1