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

migrateUp1:
	migrate   -path db/migration -database "postgresql://root:nde2024@localhost:5433/simple_bank?sslmode=disable" -verbose up 1

migrateDown:
	 migrate -path db/migration -database "postgresql://root:nde2024@localhost:5433/simple_bank?sslmode=disable" -verbose down

migrateDown1:
	migrate  -path db/migration -database "postgresql://root:nde2024@localhost:5433/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go  github.com/simpleBank/db/sqlc Store

.PHONY: createDb, dropDb, setPostgres, stopPostgres, restartPostgres, migrateUp, migrateDown, sqlc, test, server, mock, migrateUp1, migrateDown1