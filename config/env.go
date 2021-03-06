package config

import "github.com/go-sql-driver/mysql"

// 环境配置文件
// 可配置多个环境配置，进行切换

type Env struct {
	DEBUG            bool
	DATABASE         mysql.Config
	MAXIDLECONNS     int
	MAXOPENCONNS     int
	SERVER_PORT      string
	REDIS_IP         string
	REDIS_PORT       string
	REDIS_PASSWORD   string
	REDIS_DB         int
	REDIS_SESSION_DB int
	REDIS_CACHE_DB   int
	APP_SECRET       string

	ACCESS_LOG      bool
	ACCESS_LOG_PATH string
	ERROR_LOG       bool
	ERROR_LOG_PATH  string
	INFO_LOG        bool
	INFO_LOG_PATH   string

	SQL_LOG bool

	TEMPLATE_PATH string // 静态文件相对路径
}

var env = Env{
	DEBUG: true,

	SERVER_PORT: "4000",

	DATABASE: mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3306",
		DBName:               "gin-template",
		Collation:            "utf8mb4",
		Net:                  "tcp",
		AllowNativePasswords: true,
	},
	MAXIDLECONNS: 50,
	MAXOPENCONNS: 100,

	REDIS_IP:       "127.0.0.1",
	REDIS_PORT:     "6379",
	REDIS_PASSWORD: "",
	REDIS_DB:       0,

	REDIS_SESSION_DB: 1,
	REDIS_CACHE_DB:   2,

	ACCESS_LOG:      true,
	ACCESS_LOG_PATH: "storage/logs/access.log",

	ERROR_LOG:      true,
	ERROR_LOG_PATH: "storage/logs/error.log",

	INFO_LOG:      true,
	INFO_LOG_PATH: "storage/logs/info.log",

	TEMPLATE_PATH: "frontend/templates",

	//APP_SECRET: "YbskZqLNT6TEVLUA9HWdnHmZErypNJpL",
	APP_SECRET: "something-very-secret",
}

func GetEnv() *Env {
	return &env
}
