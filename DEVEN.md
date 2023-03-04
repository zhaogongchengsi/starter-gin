## Configure the development runtime environment

### 1. Install docker

[Click to view the installation method](https://docs.docker.com/get-docker/)

### 2. Start mysql

```sh
docker pull mysql
# F:\mysql8\data Container volume mount, data persistence
# MYSQL_ROOT_PASSWORD=123456 user:root password:123456
# MYSQL_DATABASE=starter Database created after starting the container
docker run --name starter-mysql-8089 -v F:\mysql8\data:/var/lib/mysql -p 8089:3306 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=starter -d mysql:latest
docker exec -it starter-mysql-8089 bash
```

### 3. Start redis

```sh
docker pull redis
docker run --name starter-gin-redis-8088 -p 8088:3306 -d mysql:latest
docker exec -it starter-gin-redis-8088 bash
```
