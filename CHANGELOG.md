### Unreleased
- Next version each microservice will have its API RESTful to communicate to each other:
  - Based on OpenAPI specification
  - [Oapi-codegen] to generate code from OpenAPI specification
  - [Echo] web framework 
  - Wiremock library for test apis
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

[v0.8.0]: https://github.com/jalv82/go-bookstore-microservices/compare/0.7.0...0.8.0
[v0.7.0]: https://github.com/jalv82/go-bookstore-microservices/compare/0.6.0...0.7.0
[v0.6.0]: https://github.com/jalv82/go-bookstore-microservices/compare/0.5.0...0.6.0
[v0.5.0]: https://github.com/jalv82/go-bookstore-microservices/releases/tag/0.5.0

[Echo]: https://echo.labstack.com/
[godog]: https://github.com/cucumber/godog
[Gomock]: https://github.com/golang/mock
[Gorm]: https://gorm.io/
[Go-sqlmock]: https://github.com/DATA-DOG/go-sqlmock
[Migrate]: https://github.com/golang-migrate/migrate
[Oapi-codegen]: https://github.com/deepmap/oapi-codegen
[Testconstainers]: https://github.com/testcontainers/testcontainers-go
[Testify]: https://github.com/stretchr/testify
[Viper]: https://github.com/spf13/viper
[uuid]: https://github.com/google/uuid
[Zerolog]: https://github.com/rs/zerolog
