# GoFrame Project

https://goframe.org

```
启动
go run main.go
http://localhost:8199
或者
go build main.go
./main.exe


```

# 1 初始化项目
```
下载客户端工具
https://github.com/gogf/gf-cli/blob/master/README_ZH.MD
https://goframe.org/cli/windows_amd64/gf.exe
将gf.exe放至%gopath%\bin目录下
gf init gf-demo
目录结构
/
├── app                 业务逻辑层        所有的业务逻辑存放目录。
│   ├── api             业务接口         接收/解析用户输入参数的入口/接口层。
│   ├── dao             数据访问         数据库的访问操作，仅包含最基础的数据库CURD方法
│   ├── model           结构模型         数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义
│   └── service         逻辑封装         业务逻辑封装层，实现特定的业务需求，可供不同的包调用。
├── boot                初始化包         用于项目初始化参数设置，往往作为main.go中第一个被import的包。
├── config              配置管理         所有的配置文件存放目录。
├── docker              镜像文件         Docker镜像相关依赖文件，脚本文件等等。
├── document            项目文档         Documentation项目文档，如: 设计文档、帮助文档等等。
├── i18n                I18N国际化       I18N国际化配置文件目录。
├── library             公共库包         公共的功能封装包，往往不包含业务需求实现。
├── packed              打包目录         将资源文件打包的Go文件存放在这里，boot包初始化时会自动调用。
├── public              静态目录         仅有该目录下的文件才能对外提供静态服务访问。
├── router              路由注册         用于路由统一的注册管理。
├── template            模板文件         MVC模板文件存放的目录。
├── Dockerfile          镜像描述         云原生时代用于编译生成Docker镜像的描述文件。
├── go.mod              依赖管理         使用Go Module包管理的依赖描述文件。
└── main.go             入口文件         程序入口文件。
go mod tidy

```
# 2 建数据库 和 表
```
数据库名：gf-demo
字符集：utf8mb4
排序：utf8mb4_bin
用户表：document/sql/create.sql
```
# 3 命名设计
```
包名
1.简单
2.推荐通过不同的import路径来区分相同包名的包引入
3.推荐使用面向对象的封装方式，项目中app目录下往往只存在api, dao, model, service 四个包名，
  每个包内部通过对象的形式来封装具体的数据模型或者业务逻辑。

```
# 4 功能设计
```
1.用户注册
2.用户登录
3.判断用户是否已经登录
4.用户注销/退出
5.检测用户账号是否存在 注册时检测
6.检测用户昵称是否存在 注册时检测
7.获取用户详情信息
```
# 5 代码生成
```
默认支持的数据库类型为：MySQL/MariaDB、PostgreSQL、SQLServer。
gf gen -h

生成dao
gf gen dao -h
gf gen dao -l "mysql:root:123@tcp(192.168.230.129:3306)/gf-demo"  -path ./app  -t user

生成的目录
|-- dao
|   |-- internal
|   |   `-- user.go
|   `-- user.go
|-- model
|   |-- internal
|   |   `-- user.go
|   `-- user.go


```
# 6 功能开发
```
完成注册功能
1.model新建结构体 UserServiceSignUpReq 和 UserApiSignUpReq (含校验)
2.service中创建一个userService 空结构体
3.service给userService创建一个方法 SignUp，参数为*model.UserServiceSignUpReq，返回值为 error
4.library/response封装返回信息
5.api中创建一个userApi 空结构体
6.api中完成注册功能并编写swagger，生成命令：gf swagger --pack
  boot中加入swagger
  查看地址：http://localhost:8199/swagger
7.router中编写路由
8.config中编写配置文件

```

# 7 测试
```
注册
curl -d 'nickname=john&passport=test001&password=123456&password2=123456' http://localhost:8199/user/sign-up

```
# 8 部署
```
独立部署
1 nohup ./gf-app &
2 tmux
3 supervisor
4 systemctl
5 screen

代理部署
见document/nginx


容器部署
docker build -t main .
docker run main
```
