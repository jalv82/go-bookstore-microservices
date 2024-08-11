all: databases-init run

#.PHONY: run
#run: build
#	@echo "🔥 Running executables..."
#	bin/bookstore-author-ms
#	bin/bookstore-book-ms
#	@echo "🌴️ All executables run!"

.PHONY: build
build: tests
	@echo "🔥 Building executables..."
	@mkdir bin
	@go build -o bin/bookstore-author-ms bookstore-author-ms/cmd/main.go
	@go build -o bin/bookstore-book-ms bookstore-book-ms/cmd/main.go
	@echo "🌴 Build done!"

.PHONY: tests
tests:
	@echo "🔥 Running tests..."
	@go test bookstore/bookstore-author-ms/internal/author/application/crud
	@go test bookstore/bookstore-author-ms/internal/author/domain/model
	@go test bookstore/bookstore-author-ms/internal/author/infrastructure/api_restful
	@go test bookstore/bookstore-author-ms/internal/author/infrastructure/sql_database
	@go test bookstore/bookstore-author-ms/acceptance_tests
	@go test bookstore/bookstore-book-ms/internal/book/domain
	@go test bookstore/bookstore-book-ms/internal/book/infrastructure/database
	@go test bookstore/bookstore-book-ms/acceptance_tests
	@echo "🌴 All tests passed!"

.PHONY: mocks-generate
mocks-generate:
	@mockgen -package=infrastructure -source=bookstore-book-ms/internal/book/domain/service.go -destination=bookstore-book-ms/internal/book/infrastructure/repository_mock.go

.PHONY: openapi-client
openapi-client:
	@oapi-codegen -package openapi bookstore-author-ms/internal/author/infrastructure/api_restful/openapi/api-author.yaml > bookstore-author-ms/internal/author/infrastructure/api_restful/openapi/client.gen.go

.PHONY: databases-init
databases-init: databases-up databases-populate

databases-up: CMD=up -d
databases-stop: CMD=stop
databases-down: CMD=down

databases-up databases-stop databases-down:
	@docker-compose $(CMD)

.PHONY: databases-populate
databases-populate:
	@echo "🗂️ Populating databases..."
	@migrate -database 'postgres://bookstore:_b00kSt0r3_@localhost:5431/bookstore-author?sslmode=disable' -path bookstore-author-ms/scripts/migrator up
	@migrate -database 'postgres://bookstore:_b00kSt0r3_@localhost:5432/bookstore-book?sslmode=disable' -path bookstore-book-ms/scripts/migrator up
	@echo "🌴 All databases populated!"

.PHONY: databases-clean
databases-clean:
	@echo "🧹 Cleaning databases..."
	@migrate -database 'postgres://bookstore:_b00kSt0r3_@localhost:5431/bookstore-author?sslmode=disable' -path bookstore-author-ms/scripts/migrator drop -f
	@migrate -database 'postgres://bookstore:_b00kSt0r3_@localhost:5432/bookstore-book?sslmode=disable' -path bookstore-book-ms/scripts/migrator drop -f
	@echo "🌴 All databases cleaned!"

.PHONY: executables-clean
executables-clean:
	@echo "🧹 Cleaning executables..."
	@rm -r bin
	@echo "🌴 All executables cleaned!"

.PHONY: clean
clean: databases-clean databases-down executables-clean
