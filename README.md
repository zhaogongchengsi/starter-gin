# Using golang's gin-based backend template

## Start

```sh
go run main.go
go build main.go
```

## Command

```
  -c, --config:      Directory where configuration files are stored (default: configs)
  -t, --configType:  Type of configuration file (default: yaml)
  -s, --seed:        Initialize the database seed data parameter to database url
  -h, --help:        help
```

### Seed

```sh
go run main.go --seed "root@tcp(localhost:3306)/starter_gin?charset=utf8mb4&parseTime=True&loc=Local"
go run main.go -s "root@tcp(localhost:3306)/starter_gin?charset=utf8mb4&parseTime=True&loc=Local"
```
