#! /bin/

# 此文件放在服务器上
# This document is prepared according to the actual situation 根据实际情况添加
cd /app/ # 代码目录
echo "run go mod tidy"
go mod tidy
echo "go build ./cmd/server/main.go -> test"
go build -o test ./cmd/server/main.go
echo "restart service -> starter"
sudo systemctl stop starter
sudo systemctl start starter

#sudo vi /etc/systemd/system/starter.service
#[Unit]
#Description=My starter service 描述
#
#[Service]
#Type=simple
#ExecStart=/xxx # 执行文件
#WorkingDirectory=/xxx # 工作目录
#Restart=on-failure # 重启策略
#RestartSec=30s # 启动等待的时间
#User=ubuntu # 启动用户
#
#[Install]
#WantedBy=multi-user.target

# sudo systemctl daemon-reload 让配置文件生效