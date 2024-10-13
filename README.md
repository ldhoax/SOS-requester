
# SOS requester API - Golang Gin Gorm


## Prerequisites

1. Install go. You can download the Golang in this [page](https://go.dev/doc/install). You should install version 1.20
2. Install Postgres database. You can download the Postgres in this [page](https://www.postgresql.org/download/). You should install version 14.1
3. Make an `.env` file from `.env.example`
4. Go to `pgadmin` create a database. Note the name of it and add to `.env`
5. Install **Air - live reload fo Go apps**. You can visit this [page](https://github.com/cosmtrek/air).
6. Install migrate to migration table to database. You can visti this [page](https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md).

## 💿 Installation

#### Via `go`

1. You run this command to install packages
   ```sh
   go mod download && go mod tidy
   ```
2. Create `.env` file from `.env.example` file.
3. run this command to start (hot reload):
   ```sh
   make watch
   ```
   run without hot reload
   ```sh
   make run
   ```
4. Visit: http://localhost:8088/api/v1/health

5. Migration database:
    ```
    migrate -database ${POSTGRESQL_URL} -path database/migrations up
    ```

#### Via `docker`
Run by docker
```sh
docker-compose up
```

## Folder structure of project

**/cmd**
The folder will contain the main applications for this project.
There might be more than 1 application within cmd folder, and each application should have its own folder so the path would be cmd -> applicationNameFolder.

**/api**
This folder can have OpenAPI/Swagger specs, JSON schema files, and protocol definition files.

**/configs**
It can have configuration file templates or default configs.

**/internal**
Private application and library code. This is the code you don't want others importing into their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. Note that you are not limited to the top-level internal directory. You can have more than one internal directory at any level of your project tree.

**/pkg**
Library codes which can be used by external applications (e.g. /pkg/mypubliclib). Other projects will import these libraries expecting them to work, so double-check before you put something here.

**/web**
Web application-specific components: static web assets, server-side templates and SPAs.

**/http**
You might want to expose your application through several means of communication such as rest-api and grpc.
This way you can have a separate separation between each type of communication layer by creating a separate directory such as /http/rest or /http/grpc

**/tests**
The folder will contain unit test, testcase of the project.

**/utils**
The directory contains supporting tools for this project.


## 🧪 Testing
```sh
make test
# or
go test ./tests -cover
# with cover
go test ./tests -cover
# with verbose
go test -v ./test -cover
# specific folder
go test -v ./utils -cover
# specific test file
go test ./utils/array_test.go ./utils/array.go
# one unit test
# - utils is a package name
# - TestChunkSlice is a testcase
go test utils -run TestChunkSlice
```

### 🧪 Improve code with lint checks
```sh
make lint
```

## Demo
### Health

```sh
curl -L 'localhost:8088/api/v1/health'
```

response:
```json
{
  "status": "success",
  "message": "live",
  "data": {
    "welcome": "Requester API alive"
  }
}
```

## Authentication

Users no need to login to use the API.

## Authorization

For the first time creating a `SOS request`, we will create a `requester` and return a token.
After that, the user can access to 'my requests' section by adding the token to `Authorization` header with `Bearer {token}` format.
Users also update their information by using the token.