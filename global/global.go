package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/miaogu-go/Gof/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GofConfig config.Server
	GofDB     *gorm.DB
	GofRedis  *redis.Client
	GofLog    *zap.Logger
	GofMongo  *mongo.Client
)
