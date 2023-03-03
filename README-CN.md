# 使用 golang 基于 gin 框架创建的一个后端模板

## 启动

```sh
go run main.go
go build main.go
```

[在线 Api 文档](https://www.apifox.cn/apidoc/project-2379970/api-65717385)

## 命令

```
  -c, --config:      设置配置文件的路径 (默认:./)
  -t, --configType:  设置配置文件的类型 (默认: yaml)
  -s, --seed:        初始化一些数据库的种子数据，需要指定数据库的url
  -g, --gsc:         生成一个本地的ssl证书
  -h, --help:        帮助
```

### Seed

```sh
go run main.go --seed "root@tcp(localhost:3306)/starter_gin?charset=utf8mb4&parseTime=True&loc=Local"
go run main.go -s "root@tcp(localhost:3306)/starter_gin?charset=utf8mb4&parseTime=True&loc=Local"
```

### Gsc

创建证书之前需要确认存放证书的文件夹是否存在 例如 ssl文件夹必须存在

```
go run main.go -g ssl
```
