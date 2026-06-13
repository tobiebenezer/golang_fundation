using go-migrate to migrate the database

migrate -path=./cmd/migration -database="postgres://postgres:Awodumila@localhost:5432/gosocial?sslmode=disable" up
migrate -path=./cmd/migration -database="postgres://postgres:Awodumila@localhost:5432/gosocial?sslmode=disable" down