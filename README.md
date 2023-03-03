# Using golang's gin-based backend template


[中文文档](./README-CN.md)
## Start

```sh
go run main.go
go build main.go
```

[Online API Documentation](https://www.apifox.cn/apidoc/project-2379970/api-65717385)


## Command

```
  -c, --config:      Directory where configuration files are stored (default: configs)
  -t, --configType:  Type of configuration file (default: yaml)
  -s, --seed:        Initialize the database seed data parameter to database url
  -g, --gsc:         Generate ssl certificate
  -h, --help:        help
```

### Seed

```sh
go run main.go --seed "root@tcp(localhost:3306)/starter_gin?charset=utf8mb4&parseTime=True&loc=Local"
go run main.go -s "root@tcp(localhost:3306)/starter_gin?charset=utf8mb4&parseTime=True&loc=Local"
```

### Gsc
Before creating a certificate, you need to create a certificate storage
```
go run main.go -g ssl
```
