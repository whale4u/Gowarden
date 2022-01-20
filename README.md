# 简介
基于gin+casbin+jwt实现REST接口。

# 接口
## 登录接口
```http request
curl --location --request POST 'http://127.0.0.1:8090/login' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"admin", "password":"pass"}'
```
## 主页接口
```http request
curl --location --request POST 'http://127.0.0.1:8090/v1/home' \
--header 'token: xxx'
```