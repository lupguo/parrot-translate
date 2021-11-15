# parrot-translate

鹦鹉翻译，支持对接有道、Google多个云API的HTTP服务

## 安装
```shell
go get -u github.com/lupguo/parrot-translate
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

## 使用
```shell
./parrot-translate -c app.yaml 
```