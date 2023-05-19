# Using golang's gin-based backend template

[中文文档](./README-CN.md)

[API Document](https://www.apifox.cn/apidoc/project-2379970/api-65717385)

[starter-vue](https://github.com/zhaogongchengsi/starter-vue) corresponding front-end project


## Build
```shell
# Build all
make
# Build separately
make build-windows
make build-linux
make build-macos
```

## Command

### Server
*cmd/server/main.go* Entrance to the project
```sh
# start
go run /cmd/server/main.go
# pack
go build /cmd/server/main.go
```

### tools
*cmd/tools/main.go* An executable of a utility class
```sh
go build -o bin/tools.exe ./cmd/tools/main.go

./bin/tools.exe -h
```

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

## Launch a local development environment

```sh
# Check if the file is normal
docker-compose -f docker-compose-dev.yaml config
# start
docker compose -f "docker-compose-dev.yaml" up -d --build
```

