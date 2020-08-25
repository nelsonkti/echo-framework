# 【echo-framework】

echo-framework是基于echo搭建用于快速开发的基础框架

## 目录

- [技术栈](#技术栈)

- [项目结构](#项目结构)

- [环境搭建](#环境搭建)
  
- [编码规范](#编码规范)
  - [分支管理 git](#分支管理)
  - [变量命名 驼峰式](#变量命名)
  - [路由风格 restful](#路由风格)
  - [生成protobuf](#生成protobuf)
  
## 技术栈 
[[返回目录]](#目录)
- Golang 尽量和最新版本保持一致
- MySQl 5.7
- gorm  数据库 orm
- echo  http路由框架 
- socketio  底层通信协议
- nsq 消息队列
- Redis 缓存设备在线状态

## 项目结构 
[[返回目录]](#目录)
- logic 包含了所有业务逻辑

- socket 通信相关的代码

- main 程序启动入口, 主要可以启动 http

## 环境搭建 
[[返回目录]](#目录)

- 项目的生产环境和测试环境都在远程服务器, 自己开发时需要部署本地环境

- 本地需要安装 nsq, redis, (建议使用 docker)

- 开发语言使用 golang, ide 使用 goland


## 编码规范 
[[返回目录]](#目录)

### 分支管理 
[[返回目录]](#目录)

每次开发时, 从最新的 uat 分支上, 创建自己的新功能分支 myname_feature, 
开发完成后提交 合并请求 到 uat分支,

测试时将自己的功能分支合并到 develop 进行测试,
不要将 develop 合并到任何其他分支.

### 变量命名 
[[返回目录]](#目录)

驼峰式, 小写开头, 例如 productTotal.

### 路由风格
[[返回目录]](#目录)

设计接口路由时遵循 restful 风格.
uri 表示资源定位, method 表示操作资源的方法;  
例如 put /chatroom/:id/member:mid 表示修改某群内某个成员

### 生成protobuf
[[返回目录]](#目录)

生成命令: protoc --go_out=. lib/protobuf/*.proto

也可以直接:make proto