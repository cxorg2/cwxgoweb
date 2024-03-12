package config

// 本地配置
type ConfigLocal struct {
	localAddr   string
	ConfSourcce string // 配置来源, env or cli
}

// webServer 模块
type ConfWebServer struct {
	On   bool
	Port int
}

// GinWeb 模块
type ConfGinWeb struct {
	On   bool
	Port int
}

// LogTask 模块
type ConfLogTask struct {
	On bool
}

// 压力模块配置
type ConfStress struct {
	On           bool
	StressSize   int    // 预设压力大小; 1: low, 2: medium, 3: high
	ConfSourcce  string // 配置来源, env or cli
	CacheSleepMs int
	MysqlSleepMs int
}

type ConfMetrics struct {
	On          bool
	MetricsPort int
	ConfSourcce string // 配置来源, env or cli
}

type ConfRedis struct {
	On                bool
	RedisAddress      string `yaml:"redisAddress"`
	RedisPort         string
	RedisAuthPassword string
	ConfSourcce       string // 配置来源, env or cli
	SleepMs           int
}

type ConfMysql struct {
	On          bool
	Dsn         string
	ConfSourcce string // 配置来源, env or cli
	SleepMs     int
}

type ConfBlogApi struct {
	On          bool
	Port        string
	ConfSourcce string
}

type ConfBlogMysql struct {
	On          bool
	Dsn         string
	ConfSourcce string // 配置来源, env or cli
}
