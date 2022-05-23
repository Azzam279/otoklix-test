# Otoklix Test

Otoklix Technical Test

## Precondition
```bash
go get github.com/mattn/go-sqlite3
go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get github.com/stretchr/testify
```

## Run App

Run the following command to run the app

```bash
go run .
```

## Build App
Run the following command to build the app

```bash
go build
```

## Run Tests
Run the following command to run unit tests

```bash
cd tests
go test -v
```
Or
```bash
go test -v ./...
```
