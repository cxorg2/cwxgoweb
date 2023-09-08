
export CWX_WEBSERVER_ON=True
export CWX_WEBSERVER_PORT=19002

# export CWX_LOGTASK_ON=True

# export CWX_GENERATEDATA_MYSQL_ON=True
export CWX_DB_ADDR="10.2.1.5"
export CWX_DB_PORT="3306"
export CWX_DB_USER="wait"
export CWX_DB_PASSWORD="passw0rd"
export CWX_DB_NAME="mytest"
export CWX_DB_CHARSET="utf8mb4"
export CWX_DB_SLEEP_MS="100"

# export CWX_GENERATEDATA_REDIS_ON=True
export CWX_REDIS_ADDRESS="10.2.1.18"
export CWX_REDIS_PORT="6379"
export CWX_REDIS_SLEEP_MS="200"

# export CWX_METRICS_ON=True
export CWX_METRICS_PORT=19003

# export CWX_GINWEB_ON=True
export CWX_GINWEB_PORT=19004
# export GIN_MODE=release

go run src/main.go
