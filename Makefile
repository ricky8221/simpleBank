createDb:
	docker exec -it postgresSimpleBank createdb --username=root --owner=root simple_bank

dropDb:
	docker exec -it postgresSimpleBank dropdb simple_bank

gotoDb:
	 docker exec -it postgresSimpleBank psql -U root simple_bank

 setPostgres:
	docker run --name postgresSimpleBank -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=nde2024 -d postgres

stopPostgres:
	docker stop postgresSimpleBank;

restartPostgres:
	docker restart postgresSimpleBank;

migrateUp:
	 migrate -path db/migration -database "postgresql://root:nde2024@localhost:5433/simple_bank?sslmode=disable" -verbose up

migrateDown:
	 migrate -path db/migration -database "postgresql://root:nde2024@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createDb, dropDb, setPostgres, stopPostgres, restartPostgres, migrateUp, migrateDown, sqlc, test, server