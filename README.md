# full-moon-backend

```
docker-compose build

docker-compose up -d

docker exec -it full-moon-backend /bin/sh

docker-compose exec bash
```

```
go run ./server.go
```

```
docker-compose exec db /bin/bash

psql -U postgres


```

```
## full-moon-backend

migrate create -ext sql -dir db/migrations -seq create_bookmarks_table

export POSTGRESQL_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_USER}?sslmode=disable

migrate -database ${POSTGRESQL_URL} -path db/migrations up 1
```
```
## db

export POSTGRESQL_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_USER}?sslmode=disable

psql ${POSTGRESQL_URL}
```