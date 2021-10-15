package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/miaogu-go/Gof/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GofConfig config.Server
	GofDB     *gorm.DB
	GofRedis  *redis.Client
	GofLog    *zap.Logger
)
