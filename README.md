# Scaleway Homework - Servers API

## Run migrations
Don't forget to fill your .env file with your database credentials (see .env.example)
```bash
$ go run migrate/migrate.go
```

## Start the Go API
### Development environment
```bash
$ go run main.go
```

### With watcher
```bash
$ CompileDaemon -build="go build -o main" -command="./main"
```

## You can check health of the API
```bash
$ curl http://localhost:8080/health
```

## Documentation
You can find the documentation of the API here: [http://localhost:8080/openapi/index.html](http://localhost:8080/openapi/index.html)

## Run tests
```bash
$ go test ./...
```

### Run tests with coverage
```bash
$ go test ./... -coverpkg=./...
```
