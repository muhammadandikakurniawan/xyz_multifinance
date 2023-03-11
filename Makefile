PACKAGE_NAME := github.com/muhammadandikakurniawan/xyz_multifinance

deploy:
	docker-compose up --build -d

swagger:
	cd cmd/app
	swagger generate spec --scan-models --output=./docs/swagger.json

swaggo:
	cd cmd/app
	swag init

wire:
	go run github.com/google/wire/cmd/wire

run:
	cd cmd/app 
	go run main.go

install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get golang.org/x/tools/cmd/cover
	go install github.com/google/wire/cmd/wire@latest
	go get github.com/google/wire/cmd/wire
	go mod download
	go mod tidy

test:
	@echo "=================================================================================="
	@echo "Coverage Test"
	@echo "=================================================================================="
	go fmt ./... 
	go test -coverprofile coverage.cov -cover ./...
	@echo "\n"
	@echo "=================================================================================="
	@echo "All Package Coverage"
	@echo "=================================================================================="
	go tool cover -func coverage.cov

mock:
	@echo "=================================================================================="
	@echo "Generating Mock"
	@echo "=================================================================================="
	@echo "Make sure install mockery https://github.com/vektra/mockery#installation"
	@echo "\n"
	mockery --all --output src/shared/mocks --case underscore