# Transaction Routines

It is Transaction routine project with simple clean code design for 3 REST API

## Installation
#### Requirements

```
docker 25.0.3 or later
docker-compose 2.24.5 or later
```
Instalation:
```
brew install golang-migrate
bash run.sh
```
Tests:
```
make tests
```

Mocks:
```
Make mockgen
```

## ABOUT PROJECT 

### Libraries and Framework
* [https://github.com/gofiber/fiber](https://github.com/gofiber/fiber) Framework used for this project for routing, logging.
* [https://github.com/jackc/pgx](https://github.com/jackc/pgx) Postgres driver used to connect and interact with database
* [https://github.com/golang-migrate/migrate](https://github.com/jackc/pgx) Migrate is migration assistant for sql migration
* [https://github.com/testcontainers/testcontainers-go](https://github.com/testcontainers/testcontainers-go)  Test containers for go used for running database and integration test.
* [https://github.com/golang/mock](https://github.com/golang/mock) Mocking library to mock the services and repository for unit test.
* [https://github.com/swaggo](https://github.com/swaggo) To Generate the api documentation

### Documentation
Api documentation can be found on this After running the project locally
http://127.0.0.1:8080/v1/swagger/index.htm

## Database
* Postgres

### Technical Architecture
It is based on clean code architecture principles.It is DDD(domain driven design) as the project resembles banking and transactions so DDD suits the requirement.

The Main component of the Design are as below
1. Handlers -> They take the request from user and validate the request for required fields

2. Services-> Each Service acts as an “orchestrator” for a subset of functionality within the system, handling interactions between various components.

3. Repository-> It deals only with handling data in database for add, update or delete or specific.

4. Providers-> It will only help to Get the data from database without any manipulations

5. Domain -> All the business logic for system will go into this which help to modify domain easily

6. Mappers -> It will map domain to database object or vice versa

7. BaseDb ->It hold the implementation for insert, update.

The core principle for this is each component will work in its boundary.
This makes the code very modular and easily testable and extendable

### Source Tree
```
├── Dockerfile
├── Makefile
├── README.md
├── api
│   ├── handlers
│   └── router
├── cmd
│   ├── logger.go
│   └── server.go
├── commons
│   └── utils
├── core
│   ├── data
│   ├── domain
│   ├── persistence
│   └── service
├── customerror
│   ├── customerror.go
│   ├── error.go
│   ├── error_handler.go
│   └── error_handler_test.go
├── db
│   ├── base_db.go
│   ├── db.go
│   ├── db_config.go
│   └── migration
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── integration_test
│   ├── account_integration_test.go
│   ├── test_app.go
│   └── transaction_integration_test.go
├── main.go
└── run.sh
```





