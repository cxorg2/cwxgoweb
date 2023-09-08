
[![Build Status](https://drone.services.wait/api/badges/chenwx/cwxgoweb/status.svg)](https://drone.services.wait/chenwx/cwxgoweb)


# 描述
一个 golang 的测试和练手项目


# 核心功能

1. 功能开关: 通过 env 决定启动时是否开启某些功能
2. 数据生成: 往 redis, mysql, 内生成随机数据
3. 环境信息: 通过 api 获取环境信息


# 发布
1. 支持本地二进制启动
2. 支持 docker 本地打包
3. 支持 docker-compose 启动
4. 支持 k8s 环境

# 运行
CWX_WEBSERVER_ON=True go run src/main.go

# TODO
1. [ ] 增加一些杂七杂八的功能
2. [ ] 增加 github actions 的 workflows 编译功能
3. [ ] 增加 CPU 资源消耗波动和内存波动的功能
