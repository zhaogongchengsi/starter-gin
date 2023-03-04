## 配置开发运行环境

### 1. 安装 docker

```sh

```

### 2. 启动 mysql

```sh
docker pull mysql
docker run --name starter-gin-mysql-8089 -p 8089:8080 -e MYSQL_ROOT_PASSWORD=12345 -d mysql:latest
```

### 3. 启动 redis

```sh
docker pull redis
docker run --name starter-gin-redis-8088 -p 8088:3306 -d mysql:latest
```
