
DB_URL=postgres://user:super-secret@0.0.0.0:5432/transaction-routines?sslmode=disable

migrate-up:
	migrate -database "${DB_URL}" -path db/migration -verbose up

migrate-down:
	migrate -database "${DB_URL}" -path db/migration -verbose down

mockgen:
	@printf "Removing existing mocks...\n"
	@rm -rf mocks
	@printf "Generating mocks...\n"
	@go generate ././core/...
	@go generate ./db/...
	@printf "Mocks Generated"