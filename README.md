# go-sql-crud docker-composed app

<!-- TOC -->
* [go-sql-crud docker-composed app](#go-sql-crud-docker-composed-app)
    * [About](#about)
    * [Technical key points](#technical-key-points)
      * [Ports](#ports)
    * [How to run app](#how-to-run-app)
    * [About Tests](#about-tests)
      * [How to run tests](#how-to-run-tests)
    * [Directory structure](#directory-structure)
    * [About APIs](#about-apis)
      * [APIs contract](#apis-contract)
      * [Sample request/response](#sample-requestresponse)
<!-- TOC -->

### About
- a `docker-composed golang sql` app that performs CRUD operations.
- areas focused are `maintainability`, `clean code`, `simplicity`, `scalability`, `testability`.
- demonstrates system design principles, code organization, OOP principles, and best practices in golang.

### Technical key points
#### Ports
- `8080` - web port
- `3306` - mysql port

### How to run app

https://github.com/harssRajput/go_crud_sql/assets/82873133/65c3fb63-e068-4ff9-a583-582952007c82

https://github.com/harssRajput/go_crud_sql/assets/82873133/4dbd6cbd-da2e-4b25-87c9-48ff7f369cd5


1. clone the repo
```shell
git clone git@github.com:harssRajput/go_crud_sql.git
```
2. install [docker][docker] and docker-compose.
3. open terminal and head over to the project directory
3. Now, Two ways to run the app:
   1. Using docker-compose (recommended one)
```shell
docker-compose up
```
   ii. manual deployment of mysql and golang app
```shell
docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -v ./scripts/init.sql:/docker-entrypoint-initdb.d/1.sql -p 3306:3306 mysql
go run main.go
``` 

### About Tests
Two type of test cases are written which are covering all the scenarios.
1. Unit test cases - it covers all the scenarios of the functions.
2. Integration test cases - it covers end to end functionality of the APIs.

#### How to run tests

1. install docker and docker-compose. 
2. take a clone and head over to the project directory as mentioned above in the [`How to run app`](#how-to-run-app) section.
3. run the below command. it executes **all testcases** in the **docker container**.
```shell
go test ./... -v
```

### Directory structure
```shell
.
├── Dockerfile
├── README.md
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── scripts/
├── tests/
└── internal
    ├── entity/
    ├── handler/
    ├── server/
    ├── service/
    ├── utilityStore/
```
`main.go` - entry point of the application.

`scripts/` - contains the sql script to create the database and tables schema.

`tests/` - contains the integration test cases. (unit tests present alongside the functions)

`internal/` - contains the main codebase. all business/domain logic resides here.

`entity/` - contains the entity objects.

`handler/` - contains the request and response objects.

`server/` - contains the server configuration.

`service/` - contains the business logic of the application.

`utilityStore/` - provides everything that server needs during initialization. like db connection, logger, env vars, router, error handling, etc.


### About APIs
#### APIs contract
| API Name            | Signature                           | method |
|---------------------|-------------------------------------|--------|
| Create Account      | http://localhost:8080/accounts/     | POST   |
| Get Account         | http://localhost:8080/accounts/2    | GET    |
| Create Transactions | http://localhost:8080/transactions/ | POST   |
(NOTE: catch-all fallback route also added) 

#### Sample request/response
- Create Account

json request
```json
{
  "document_number": "12345678909"
}
```
json response
```json
{
  "account_id": 7,
  "document_number": "12345678909"
}
```

- Get Account

json response
```json
{
  "account_id": 7,
  "document_number": "12345678909"
}
```

- Create Transactions

json request
```json
{
  "account_id": 7,
  "operation_type_id": 1,
  "amount": -10
}
```

json response
```json
{
  "transaction_id": 7,
  "account_id": 7,
  "operation_type_id": 1,
  "amount": -10,
  "event_date": "2024-05-28T12:54:05.851Z"
}
```


[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. no need to format nicely because it shouldn't be seen.)
[grpcui]: <https://github.com/fullstorydev/grpcui>
[docker]: <https://docs.docker.com/desktop/install/mac-install/>
