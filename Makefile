postgres:
	 docker run --name mydb -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_DB=my_db -v pgdata:/var/lib/postgresql/data -p 5432:5432 postgres:16
start:
	docker start mydb
stop:
	docker stop mydb
createdb:
	docker exec -it mydb createdb --username=root --owner=root simple_db
dropdb:
	docker exec -it mydb dropdb --username=root --owner=root simple_db

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/simple_db?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/simple_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: postgres start stop createdb dropdb migrateup migratedown sqlc test server