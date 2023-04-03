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

### 前置指令
```sh
-c="./" # 指定配置文件目录
-t="yaml" # 指定配置文件类型
-n="config" # 指定配置文件名字
```

### Init
```shell
# 初始化app config.local 将会被git 忽略
go run main.go -c "./" -t yaml  -n "config.local" -i
```

### AutoMigrate
Automatic model migration
```sh
go run main.go --auto all
go run main.go -a all
go run main.go -a "user1,user2" # 指定模型名字 多个使用逗号分隔
go run main.go -c "./configs" -t yaml -n config -a "user" # 指定配置文件 以指定数据库
```

### Seed
Generate seed data
```sh
go run main.go --seed all
go run main.go -s all
go run main.go -s "user1,user2" # 指定模型名字 多个使用逗号分隔
go run main.go -c "./configs" -t yaml -n config -s "user1" # 指定配置文件 以指定数据库
```

### Gsc

创建证书之前需要确认存放证书的文件夹是否存在 例如 ssl文件夹必须存在

```
go run main.go -g ssl
```

## 启动一个开发环境

```sh
docker compose -f "docker-compose-dev.yaml" up -d --build
```
