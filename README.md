
[![Build Status](https://drone.services.wait/api/badges/chenwx/cwxgoweb/status.svg)](https://drone.services.wait/chenwx/cwxgoweb)

GITHUB 自动构建状态
![example workflow](https://github.com/github/docs/actions/workflows/work.yml/badge.svg)

# 描述

一个 golang 的练手综合项目

# 核心功能

api 原生 net/http 功能使用
base 程序自身配置
blog 博客后端
config 启动配置设置
generatedata 生成测试数据
ginweb 框架测试
logtask 生成随机日志
metrics 指标
stress 压力测试, 执行大随机字符串的sha512
unit 工具类

1. 功能开关: 通过 env 决定启动时是否开启某些功能
2. 数据生成: 往 redis, mysql, 内生成随机数据
3. 环境信息: 通过 api 获取环境信息

# 发布

1. 支持本地二进制启动
2. 支持 docker 本地打包
3. 支持 docker-compose 启动
4. 支持 k8s 环境

# 运行

go run src/main.go

# TODO

1. [ ] 增加一些杂七杂八的功能
2. [ ] 增加 github actions 的 workflows 编译功能
3. [ ] 增加 CPU 资源消耗波动和内存波动的功能

# run

## build

./script/build.sh

## help

```sh
wait@wait-fedora cwxgoweb]$ ./bin/cwxgoweb --help
Usage of ./bin/cwxgoweb:
  -n string
     local ipaddress, default: 10.x.x.x
  -p int
     pressure 1: low, 2: medium, 3: high; default 0 used env
  -w int
     stress an int work Goroutine number (default 30)

```
