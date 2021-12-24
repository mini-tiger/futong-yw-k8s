# futong-yw-k8s

## 目录结构说明
- base: 公用基础内容
- api: 处理请求参数和返回结果
- service: 处理业务逻辑
- storage: 处理数据库操作
- model: 数据结构体
- router: 路由
- middle: 中间件
- deploy: 项目部署
- test: 测试

## 初始化用户
```shell script
# 内置后台管理用户
name = "builtin_root"
password = "builtin_0!@"
# 更改项目配置以下4个参数设置项目初始租户
INIT_TENANT_ID = "admin"
INIT_TENANT_NAME = "主账号"
INIT_PASSWORD = "admin_123"
INIT_EMAIL = ""
```

## 本地启动
```shell script
# 项目根路径执行
go run main.go
``` 

## 接口测试
```shell script
cd test/api
go test -v
``` 
