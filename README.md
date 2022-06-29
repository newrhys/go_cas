## 说明

根据gin框架开发的MVC分层demo

### 基础功能

1. 登录/注册
2. jwt鉴权
3. 图片上传
4. redis缓存
5. gorm
6. zap日志库
7. swagger文档

### go mod tidy

    goland 对 go.mod 文件右击，然后点击 Go Mod Tidy

### swagger 文档

    swag init -g main.go

### 配置文件

    config/application.yml

### 公共Redis key

    cache/redis_key/redis_key.yml
    cache/redis_key/config.go