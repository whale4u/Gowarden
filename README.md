# 简介

基于gin+casbin+jwt实现REST接口。

目录结构参考：https://github.com/eddycjy/go-gin-example

## 实现功能

* [ ]  全局错误统一处理
* [ ]  casbin
* [X]  gorm
* [X]  jwt生成token

# 接口

## 登录接口

```http request
curl --location --request POST 'http://127.0.0.1:8090/login' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"admin", "password":"pass"}'
```

## 注册接口

`curl --location --request POST 'http://127.0.0.1:8090/register'  --header 'Content-Type: application/json'  --data-raw '{"username":"admin", "password":"pass", "phone":"181232725678"}'`

## 主页接口

```http request
curl --location --request POST 'http://127.0.0.1:8090/v1/home' \
--header 'token: xxx'
```
