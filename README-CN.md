# 一个 gin 模板项目

[英文文档](./README.md)

[API 接口文档](https://www.apifox.cn/apidoc/project-2379970/api-65717385)

[starter-vue](https://github.com/zhaogongchengsi/starter-vue) 对应的前端项目

## 命令

### Server
*cmd/server/main.go* 服务的入口文件
```sh
# 启动
go run /cmd/server/main.go
# 打包
go build /cmd/server/main.go
```

### tools
*cmd/tools/main.go* An 一个工具类的可执行文件
```sh
go build -o ./bin/tool /cmd/tools/main.go

./bin/tool -h
```

## 热更新

需要先下载[air](https://github.com/cosmtrek/air)
```shell
# install
go install github.com/cosmtrek/air@latest
# init air
air init
# start
air
```

## 快速启动开发环境

```sh
# Check if the file is normal
docker-compose -f ./deploy/docker-compose-dev.yaml config
# start
docker compose -f "./deploy/docker-compose-dev.yaml" up -d --build
```
