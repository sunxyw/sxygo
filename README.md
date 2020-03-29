# LwGO

LwGO: An simple framework based on [Singo](https://github.com/Gourouting/singo)
LwGO: 基于 [Singo](https://github.com/Gourouting/singo) 的简单框架

## 特色

本项目已经整合了许多日常开发所需要的组件：

1. [Gin](https://github.com/gin-gonic/gin)：轻量级 Web 框架，自称路由速度是 golang 最快的
2. [GORM](http://gorm.io/docs/index.html)：ORM 工具。本项目需要配合 Mysql 使用
3. [Gin-Session](https://github.com/gin-contrib/sessions)：Gin 框架提供的 Session 操作工具
4. [Go-Redis](https://github.com/go-redis/redis)：Golang Redis 客户端
5. [godotenv](https://github.com/joho/godotenv)：开发环境下的环境变量工具，方便使用环境变量
6. [Gin-Cors](https://github.com/gin-contrib/cors)：Gin 框架提供的跨域中间件
7. 自行实现了国际化 i18n 的一些基本功能
8. 本项目是使用基于 cookie 实现的 session 来保存登录状态的，如果需要可以自行修改为 token 验证

本项目已经预先实现了一些常用的代码方便参考和复用：

1. 创建了用户模型
2. 实现了 `register` 用户注册接口
3. 实现了 `login` 用户登录接口
4. 实现了 `user` 用户资料接口（需要登录后获取 session）
5. 实现了 `logout` 用户登出接口（需要登录后获取 session）

本项目已经预先创建了一系列文件夹划分出下列模块：

1. controllers 文件夹负责控制层代码，根据输入的数据进行相应操作并输出结果
2. models 文件夹负责存储数据库模型和数据库操作相关的代码
3. services 负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. transformers 储存通用的 json 模型，把 models 得到的数据库模型转换成 controllers 需要的 json 对象
5. cache 负责 redis 缓存相关的代码
6. auth 权限控制文件夹
7. utils 一些通用的小工具
8. config 放一些静态存放的配置文件，其中 locales 内放置翻译相关的配置文件
9. middleware 文件夹存放中间件

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env 文件设置环境变量便于使用

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug" # Gin 运行模式
LOG_LEVEL="debug" # 日志记录等级
```

## Go Mod

本项目使用 [Go Mod](https://github.com/golang/go/wiki/Modules) 管理依赖。

```shell
go mod init lwgo
go run main.go // 自动安装
```

## 运行

```shell
go run main.go
```

项目运行后启动在 3000 端口（可以修改，参考 gin 文档)
