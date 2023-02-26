all: databases-init run

.PHONY: run
run: build
	@echo "🔥 Running executables..."
	bin/bookstore-author
	bin/bookstore-book
	@echo "🌴️ All executables run!"

.PHONY: build
build: tests
	@echo "🔥 Building executables..."
	@mkdir bin
	@go build -o bin/bookstore-author cmd/author/main.go
	@go build -o bin/bookstore-book cmd/book/main.go
	@echo "🌴 Build done!"

.PHONY: tests
tests:
	@echo "🔥 Running tests..."
	@go test bookstore/internal/core/author/domain
	@go test bookstore/internal/core/author/infrastructure/database
	@go test bookstore/internal/core/book/domain
	@go test bookstore/internal/core/book/infrastructure/database
	@go test bookstore/acceptance-tests
	@echo "🌴 All tests passed!"

.PHONY: mocks-generate
mocks-generate:
	@mockgen -package=infrastructure -source=internal/core/author/domain/service.go -destination=internal/core/author/infrastructure/repository_mock.go
	@mockgen -package=infrastructure -source=internal/core/book/domain/service.go -destination=internal/core/book/infrastructure/repository_mock.go

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
	@migrate -database 'postgres://bookstore:_b00kSt0r3_@localhost:5431/bookstore-author?sslmode=disable' -path scripts/migrator/author up
	@migrate -database 'postgres://bookstore:_b00kSt0r3_@localhost:5432/bookstore-book?sslmode=disable' -path scripts/migrator/book up
	@echo "🌴 All databases populated!"

.PHONY: databases-clean
databases-clean:
	@echo "🧹 Cleaning databases..."
	@migrate -database 'postgres://bookstore:_b00kSt0r3_@localhost:5431/bookstore-author?sslmode=disable' -path scripts/migrator/author drop -f
	@migrate -database 'postgres://bookstore:_b00kSt0r3_@localhost:5432/bookstore-book?sslmode=disable' -path scripts/migrator/book drop -f
	@echo "🌴 All databases cleaned!"

.PHONY: executables-clean
executables-clean:
	@echo "🧹 Cleaning executables..."
	@rm -r bin
	@echo "🌴 All executables cleaned!"

.PHONY: clean
clean: databases-clean databases-down executables-clean





