
export CWX_WEBSERVER_ENABLE=True
export CWX_WEBSERVER_PORT=19002

export CWX_DB_TYPE=mysql
export CWX_DB_ADDR="10.2.1.5"
export CWX_DB_PORT="3306"
export CWX_DB_USER="wait"
export CWX_DB_PASSWORD="passw0rd"
export CWX_DB_CHARSET="utf8mb4"


# export CWX_LOGTASK_ENABLE=True

# export CWX_GENERATEDATA_MYSQL_ENABLE=True
export CWX_GENERATEDATA_DB_NAME="mytest"

export CWX_DB_SLEEP_MS="100"

# export CWX_GENERATEDATA_REDIS_ENABLE=True
export CWX_REDIS_ADDRESS="10.2.1.18"
export CWX_REDIS_PORT="6379"
export CWX_REDIS_SLEEP_MS="200"

# export CWX_METRICS_ENABLE=True
export CWX_METRICS_PORT=19003

# export CWX_GINWEB_ENABLE=True
export CWX_GINWEB_PORT=19004
# export GIN_MODE=release


# blog api
export CWX_BLOG_API_ENABLE=True
export CWX_BLOG_API_PORT=19005
export CWX_BLOG_API_DB_NAME="doc"

go run src/main.go
