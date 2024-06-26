# Create a postgres docker container 
postgres:
	docker run --name postgres1 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5432:5432 -d postgres:alpine3.19

# Create or drop the job-scheduler db
createdb:
	docker exec -it postgres1 createdb --username=root --owner=root job-scheduler

dropdb:
	docker exec -it postgres1 dropdb job-scheduler

# run migrations
migrateup:
	migrate -path go-backend/db/migrations -database postgresql://root:secret@localhost:5432/job-scheduler?sslmode=disable -verbose up

migratedown:
	migrate -path go-backend/db/migrations -database postgresql://root:secret@localhost:5432/job-scheduler?sslmode=disable -verbose down

createMigrate:
	migrate create -ext sql -dir go-backend/db/migrations  -seq <migration_name>

# Generate sqlc
sqlc:
	sqlc generate

# Start server
server:
	cd go-backend && go run main.go

# Build go docker image
build:
	cd go-backend && docker build -t goimage:latest .

# Create a network and add the postgres into in
connect:
	docker network create job-scheduler-network && \
	docker network connect job-scheduler-network postgres1

# Run go image to create a container running in the same network as postgres
run:
	docker run --name go-cont --network job-scheduler-network -d -p 8080:8080 goimage:latest 

lint:
	cd go-backend && golangci-lint run --enable-all

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server lint build run