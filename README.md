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
filePath-fileType-fileName
```sh
go run main.go --seed ./-yaml
go run main.go -s ./-yaml
```

### Gsc
Before creating a certificate, you need to create a certificate storage
```
go run main.go -g ssl
```
## Launch a local development environment

```sh
docker compose -f "docker-compose-dev.yaml" up -d --build
```