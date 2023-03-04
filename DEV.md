## 配置开发运行环境

### 1. 安装 docker

[点击查看安装方式](https://docs.docker.com/get-docker/)

### 2. 启动 mysql

```sh
docker pull mysql
# F:\mysql8\data 容器卷挂载 数据持久化
# MYSQL_ROOT_PASSWORD=123456 用户 root 密码 123456
# MYSQL_DATABASE=starter 启动容器后创建的数据库
docker run --name starter-mysql-8089 -v F:\mysql8\data:/var/lib/mysql -p 8089:3306 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=starter -d mysql:latest
docker exec -it starter-mysql-8089 bash
```

### 3. 启动 redis

```sh
docker pull redis
docker run --name starter-gin-redis-8088 -p 8088:3306 -d mysql:latest
docker exec -it starter-gin-redis-8088 bash
```
