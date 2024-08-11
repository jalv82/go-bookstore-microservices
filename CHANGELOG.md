### Unreleased
Next version will have:
- Upgrade acceptance-test of bookstore-author-ms
- Upgrade bookstore-book-ms with the same features that bookstore-author-ms
- [golangci-lint] library as linter
---

<a name="1.0.0"></a>
### [v1.0.0] - 2024-08-11 - Start milestone II
The focus has been put in bookstore-author-ms
- Added
  - [Postman collection] for play with microservice
  - Project scaffolding
    - Application layer
    - Domain layer
    - Infrastructure layer
      - API RESTful
        - OpenAPI specification
        - [oapi-codegen] to generate code from OpenAPI specification
      - HTTP server
        - [Echo] web framework
  - Testing 
    - More unit/integration tests 
    - [Testify] library to generate mocks for integration tests
    - [Validation] library for validations model
- Changed
  - Go v1.22
  - Upgrade database config with filed image
  - Updated libraries to last version
  - Updated makefile
  - Updated readme
- Removed
  - [Gomock] library to generate mocks for integration tests
---

<a name="v0.8.0"></a>
### [v0.8.0] - 2023-03-19 - Finish milestone I
- Added
  - Changelog file 
  - More acceptance tests
  - [Testconstainers] library for tests
- Changed
  - Every microservice has an independent structure
  - Updated libraries to last version
  - Updated makefile
  - Updated readme
- Removed
  - Redundant error logs
---

<a name="v0.7.0"></a>
### [v0.7.0] - 2023-02-26
- Added
  - Acceptance tests with [godog] library
- Changed
  - Updated makefile
  - Updated readme
---

<a name="v0.6.0"></a>
### [v0.6.0] - 2023-02-24
- Added
  - [Go-sqlmock] library to mock sql queries into the tests
- Changed
  - Renamed some files
  - Updated go to version v1.19
  - Updated libraries to last version
- Fixed
  - All tests run in parallel
---

<a name="v0.5.0"></a>
### [v0.5.0] - 2023-01-21
- Added
  - Project scaffolding
    - Domain driven design (since my point of view 😀)
    - Ports and adapter pattern (aka hexagonal architecture)
    - Two microservices: one for authors and the other one for books
  - Database persistence
    - PostgreSQL as database engine
    - [Migrate] library to manage sql scripts
    - [Gorm] library as ORM
  - Testing
    - [Gomock] library to generate mocks for integration tests
    - [Testify] library for tests asserts
  - Utils
    - [Zerolog] library to manage logs
    - Google [uuid] library for generate identification numbers
    - [Viper] library to read config files
  - Docker as container engine

[v1.0.0]: https://github.com/jalv82/go-bookstore-microservices/compare/0.8.0...1.0.0
[v0.8.0]: https://github.com/jalv82/go-bookstore-microservices/compare/0.7.0...0.8.0
[v0.7.0]: https://github.com/jalv82/go-bookstore-microservices/compare/0.6.0...0.7.0
[v0.6.0]: https://github.com/jalv82/go-bookstore-microservices/compare/0.5.0...0.6.0
[v0.5.0]: https://github.com/jalv82/go-bookstore-microservices/releases/tag/0.5.0

[Echo]: https://echo.labstack.com/
[godog]: https://github.com/cucumber/godog
[Gomock]: https://github.com/golang/mock
[golangci-lint]: https://golangci-lint.run/
[Gorm]: https://gorm.io/
[Go-sqlmock]: https://github.com/DATA-DOG/go-sqlmock
[Migrate]: https://github.com/golang-migrate/migrate
[Oapi-codegen]: https://github.com/deepmap/oapi-codegen
[Postman collection]: bookstore-author-ms/acceptance_tests/author-postman-collection.json
[Testconstainers]: https://github.com/testcontainers/testcontainers-go
[Testify]: https://github.com/stretchr/testify
[Validation]: https://github.com/invopop/validation
[Viper]: https://github.com/spf13/viper
[uuid]: https://github.com/google/uuid
[Zerolog]: https://github.com/rs/zerolog
