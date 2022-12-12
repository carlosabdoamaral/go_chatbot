postgres:
    docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

createdb:
    docker exec -it postgres14 createdb --username=postgres --owner=postgres go_chatbot_db

dropdb:
    docker exec -it postgres14 dropdb go_chatbot_db

migrateup:
    migrate -path internal/db/migration -database "postgresql://postgres:postgres@localhost:5432/go_chatbot_db?sslmode=disable" -verbose up

migratedown:
    migrate -path internal/db/migration -database "postgresql://postgres:postgres@localhost:5432/go_chatbot_db?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb