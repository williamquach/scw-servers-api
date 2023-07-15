# Scaleway Homework - Servers API

## Fill environment variables
### Copy the .env.example file
```bash
cp .env.example .env
```
### Fill the .env file
Fill the .env file with needed environment variables

## Launch the database with Docker
```bash
docker-compose -f docker-compose.postgres.yml up -d
```

## Run migrations to create the database
Don't forget to fill your .env file with your database credentials (see .env.example)
```bash
go run migrate/migrate.go
```

## Start the Go API
### Install dependencies
```bash
go mod download
```

### Development environment
```bash
go run main.go
```

### With watcher
```bash
CompileDaemon -build="go build -o main" -command="./main"
```

## You can check health of the API
```bash
curl http://localhost:8080/health
```

## Documentation
You can find the documentation of the API here: [http://localhost:8080/openapi/index.html](http://localhost:8080/openapi/index.html)

## Run tests
```bash
go test ./...
```

### Run tests with coverage
```bash
go test ./... -coverpkg=./...
```
