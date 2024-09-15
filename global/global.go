package global

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"umex.com/pkg/logger"
	"umex.com/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *sql.DB
	Rdb    *redis.Client
)
