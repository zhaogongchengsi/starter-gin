# Using golang's gin-based backend template

[中文文档](./README-CN.md)
## Start

```sh
go run main.go
go build main.go
```

[API Documentation](https://www.apifox.cn/apidoc/project-2379970/api-65717385)

## Live reload for Go apps

Need to install [air](https://github.com/cosmtrek/air)
```shell
# install
go install github.com/cosmtrek/air@latest
# init air
air init
# start
air
```



## Command

```
  -c, --config:      Directory where configuration files are stored (default: configs)
  -t, --configType:  Type of configuration file (default: yaml)
  -s, --seed:        Initialize the database seed data parameter to database url
  -g, --gsc:         Generate ssl certificate
  -h, --help:        help
```

### command
```sh
-c="./" # Specify the configuration file directory
-t="yaml" # Specify configuration file type
-n="config" # Specify the configuration file name
```

### Init 
```shell
# Initializing app config.local will be ignored by git
go run main.go -c "./" -t yaml  -n "config" -i
```

### AutoMigrate
Automatic model migration
```sh
go run main.go --auto all
go run main.go -a all
go run main.go -a "user1,user2" # Can be specified, multiple are separated by commas
go run main.go -c "./configs" -t yaml -n config -a "user" # Specify the configuration file to specify the database
```

### Seed
Generate seed data
```sh
go run main.go --seed all
go run main.go -s all
go run main.go -s "user1,user2" # Can be specified, multiple are separated by commas
go run main.go -c "./configs" -t yaml -n config -s "user1" # Specify the configuration file to specify the database
```

### Gsc
Before creating a certificate, you need to create a certificate storage
```
go run main.go -g ssl
```
## Launch a local development environment

```sh
# Check if the file is normal
docker-compose -f docker-compose-dev.yaml config
# start
docker compose -f "docker-compose-dev.yaml" up -d --build
```

