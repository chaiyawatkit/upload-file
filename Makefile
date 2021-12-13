example-docs-gen:
	@echo "============= Docs -> https://github.com/swaggo/swag ============= "
	swag init --dir ./example/docs-generator --swagger ./example/docs-generator/docs/swagger/

docs-gen:
	@echo "============= Docs -> https://github.com/swaggo/swag ============= "
	swag init --dir ./app --output ./app/docs/swagger/
	npx redoc-cli bundle ./app/docs/swagger/swagger.json --output ./app/docs/swagger/index.html

docs-gen-no-html:
	@echo "============= Docs -> https://github.com/swaggo/swag ============= "
	swag init --dir ./app --output ./app/docs/swagger/

dev:
	@echo "=============starting locally============="
	docker-compose -f resources/docker/docker-compose.yaml up --build

dev-app:
	@echo "=============starting locally============="
	dep ensure
	docker-compose -f resources/docker/docker-compose.yaml up --build demo_app

db:
	docker-compose -f resources/docker/docker-compose.yaml up -d postgres_db pgadmin

logs:
	docker-compose -f resources/docker/docker-compose.yaml logs -f

down:
	docker-compose -f resources/docker/docker-compose.yaml down

test:
	export GIN_MODE=release && go test ./app/... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt

test-debug:
	go test ./app/... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt

ci_test:
	go test --tags="unit" ./app/... -v -coverprofile .coverage.txt

clean: down
	@echo "=============cleaning up============="
	docker system prune -f
	docker volume prune -f
	docker images prune -f

format:
	go fmt ./app/...

dep: ## Get the dependencies
	@go get -v -d ./...
	@go get -u github.com/golang/lint/golint

migrate-create:
	migrate create -ext sql -dir app/migrations $(name)

mockgen:
	mockgen -source=./app/layers/repositories/$(module_name)/init.go -destination=./app/mocks/$(module_name)/repo.go
	mockgen -source=./app/layers/usecases/$(module_name)/init.go -destination=./app/mocks/$(module_name)/usecase.go

mockgen-all:
	mockgen -source=./app/layers/repositories/customer/init.go -package mocks -destination=./app/mocks/customer/repo.go
	mockgen -source=./app/layers/usecases/customer/init.go -package mocks -destination=./app/mocks/customer/usecase.go

mockgen-customer:
	mockgen -source=./app/layers/repositories/customer/init.go -package mocks -destination=./app/mocks/customer/repo.go
	mockgen -source=./app/layers/usecases/customer/init.go -package mocks -destination=./app/mocks/customer/usecase.go
