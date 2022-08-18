build:
	docker-compose build

run:
	docker-compose up 

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:testtask@0.0.0.0:5432/postgres?sslmode=disable' up

install-swag:
	go get -u github.com/swaggo/swag/cmd/swag
	
swag:
	swag init -g cmd/main.go