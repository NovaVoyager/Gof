package initialize

import (
	"context"
	"fmt"
	"github.com/miaogu-go/Gof/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

//LoadMongodb 加载mongo
func LoadMongodb() {
	mongoConf := global.GofConfig.Mongodb
	//mongodb://username:password@host:%d
	connStr := ""
	if mongoConf.Username != "" && mongoConf.Password != "" {
		connStr = fmt.Sprintf("mongodb://%s:%s@%s:%d", mongoConf.Username, mongoConf.Password, mongoConf.Host,
			mongoConf.Port)
	} else {
		connStr = fmt.Sprintf("mongodb://%s:%d", mongoConf.Host, mongoConf.Port)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connStr))
	if err != nil {
		global.GofLog.Error("mongo connect failed,err:", zap.Error(err))
		return
	}

	//check connect
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		global.GofLog.Error("mongo ping failed,err:", zap.Error(err))
		return
	}

	global.GofMongo = client
}
