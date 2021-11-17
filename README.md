# parrot-translate

鹦鹉翻译，支持对接有道、Google多个云API的HTTP服务

## 安装
```shell
go get -u -v github.com/lupguo/parrot-translate
```

## 应用配置文件
```shell
server:
  listen: :5567 # httpd服务监听端口
api:
  google:
    proj_id: sage-ace-xxx
    auth_file: ./sage-ace-331915-d22cf04c186b.json
```

## 单节点systemd服务配置


```shell
# 新增www用户
useradd -r www
sudo mkdir -p /var/log/parrot-translate
sudo chown -R www:www /var/log/parrot-translate

# 配置启动服务相关， 将以下内容加入到`/usr/lib/systemd/system/parrot-translate.service`
[Unit]
Description=parrot-translate httpd server

[Service]
Type=forking
User=www
Group=www
ExecStart=/usr/local/sbin/parrot-translate
Restart=always
KillMode=process

[Install]
WantedBy=multi-user.target
```

## 使用
```shell
./parrot-translate -c app.yaml 
```