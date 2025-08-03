## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Run Migration

```shell
migrate -database "mysql://root:@tcp(localhost:3306)/golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```

## Run Application

### Run unit test

```bash
go test -v ./test/
```

### Run web server

```bash
go run cmd/web/main.go
```

### Run worker

```bash
go run cmd/worker/main.go
```