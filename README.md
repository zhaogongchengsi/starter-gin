# Using golang's gin-based backend template

## Start

```sh
go run main.go
go build main.go
```

## Command

```
  -c, -config:      Directory where configuration files are stored (default: configs)
  -t, -configType:  Type of configuration file (default: yaml)
  -i, -init:        Initialize the database seed data parameter to database url
```

### Seed

```sh
go run main.go --init "root@tcp(localhost:3306)/starter_gin?charset=utf8mb4&parseTime=True&loc=Local"
go run main.go -i "root@tcp(localhost:3306)/starter_gin?charset=utf8mb4&parseTime=True&loc=Local"
```
