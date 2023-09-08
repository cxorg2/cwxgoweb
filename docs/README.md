
# generate data


# 创建过程
go mod init git.services.wait/chenwx/cwxgoweb
go mod tidy


go get -u github.com/go-redis/redis/v9



# 测试

curl http://127.0.0.1:19002/api/env -H "X-Forwarded-For: 1.1.1.1, 2.2.2.2"
