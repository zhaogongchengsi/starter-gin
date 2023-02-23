# Golang 版本的后端模板

## Start
```sh
go run main.go
go build main.go
```

# 配置文件

## 配置文件路径

配置文件路径为 `configs/*.yaml`, 配置文件格式为 `yaml`。

- database.yaml 数据库配置文件
- server.yaml 服务配置文件

### Specify profile directory
```sh
go run main.go -config=configs -configType=yaml
```