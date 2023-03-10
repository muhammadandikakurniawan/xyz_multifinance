PACKAGE_NAME := xyz_multifinance_app

swagger:
	swagger generate spec --scan-models --output=./swagger.yaml

wire:
	go run github.com/google/wire/cmd/wire

run:
	cd cmd/app 
	go run main.go

install:
	go get golang.org/x/tools/cmd/cover
	go install github.com/google/wire/cmd/wire@latest
	go get github.com/google/wire/cmd/wire
	go mod download
	go mod tidy

test:
	@echo "=================================================================================="
	@echo "Coverage Test"
	@echo "=================================================================================="
	go fmt ./... && go test -coverprofile coverage.cov -cover ./... # use -v for verbose
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