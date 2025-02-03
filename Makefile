migrate_up:
	migrate -path db/migrations -database "postgresql://root:postgres@localhost:5432/parking_lot?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://root:postgres@localhost:5432/parking_lot?sslmode=disable" -verbose down

sqlc:
	sqlc generate