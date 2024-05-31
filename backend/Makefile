.PHONY: clean build run run.build test.unit.brand test.unit.product docker.build docker.run docker.run.db docker.run.test docker.stop docker.stop.test docker.clean

### VARIABLES ###
APP_NAME = sacachispa
BUILD_DIR = ./build
DATABASE_URL = postgres://${DATABASE_USER}:${DATABASE_PASS}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable

### COMMANDS ###
clean:
	rm -rf $(BUILD_DIR)

build: clean
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api

run: docker.run.db
	@echo "Running local server..."
	DATABASE_HOST=localhost go run ./cmd/api/main.go

run.build:
	@echo "Running build app..."
	$(BUILD_DIR)/$(APP_NAME)

test.unit.brand:
	@echo "Running unit test for brand..."
	go test -v ./internal/entity/brand_test.go

test.unit.product:
	@echo "Running unit test for product..."
	go test -v ./internal/entity/product_test.go

docker.build:
	docker-compose build

docker.run: docker.clean docker.build
	@echo "Runnung server in docker container..."
	 docker-compose up -d database apiserver

docker.run.test:
	@echo "Running database for testing..."
	docker-compose up -d database-test

docker.run.db:
	docker-compose up -d database

docker.stop:
	@echo "Stop docker container..."
	docker-compose stop database apiserver database-test

docker.stop.test:
	@echo "Stop docker container for testing..."
	docker-compose stop database-test
	docker-compose rm -f database-test

docker.clean:
	docker-compose down --volumes
