fmt:
	go fmt ./...

run:
	go run cmd/main/main.go

migrate.up:
	go run cmd/migration/main.go

# test:
# 	go test -coverprofile cover.out ./src/...
# 	go tool cover -html=cover.out

up:
	docker compose up -d

stop:
	docker compose stop

down:
	docker compose down

restart:
	docker compose restart

logs:
	docker compose logs -n 30 -f

docs-update:
	rm -rf swagger/v1
	swag fmt
	swag init -d ./cmd/main/,./  -o swagger
