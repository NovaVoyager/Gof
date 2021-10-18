package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/miaogu-go/Gof/global"
	"go.uber.org/zap"
)

//LoadRedis 加载redis
func LoadRedis() {
	rdsConf := global.GofConfig.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rdsConf.Host, rdsConf.Port),
		Password: rdsConf.Password,
		DB:       rdsConf.Db,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		global.GofLog.Error("redis connect ping failed,err:", zap.Error(err))
		return
	}
	global.GofRedis = rdb
}
