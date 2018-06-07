package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic"
	"Sephiroth/config"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	rdsCli *redis.Client
	mgoCli *mgo.Session
	esCli  *elastic.Client
	start  int64
	end    int64
)

// 依赖注入
type Di struct {
	Config []map[string]map[string]string
}

func NewDi() *Di {
	di := new(Di)
	di.Config = di.GetConfig()
	return di
}

func (di *Di) GetConfig() []map[string]map[string]string {
	if di.Config != nil {
		return di.Config
	}
	return config.NewConfigParser().GetConfig()
}

func (di *Di) GetRedis() *redis.Client {
	if rdsCli != nil {
		return rdsCli
	}
	rdsConf := di.Config[1]["redis"]
	rdsConfig := &redis.Options{
		Addr:     rdsConf["host"] + ":" + rdsConf["port"],
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	rdsCli := redis.NewClient(rdsConfig)
	return rdsCli
}

func (di *Di) GetMongoDB() *mgo.Session {
	if mgoCli != nil {
		return mgoCli
	}
	mgoConf := di.Config[2]["mongodb"]
	url := mgoConf["host"] + ":" + mgoConf["port"]
	mgoCli, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(err)
	}
	return mgoCli
}

func (di *Di) GetElastic() *elastic.Client {
	if esCli != nil {
		return esCli
	}
	esCli, _ := elastic.NewClient()
	return esCli
}

// 运行消耗时间
func (di *Di) Cost(flag string) {
	if flag == "start" {
		start = time.Now().UnixNano()
		fmt.Println("start:", start)
	} else {
		end = time.Now().UnixNano()
		dura := (end - start) / 1000000
		fmt.Println("end:  ", end)
		fmt.Println("cost:", dura, " ms")
	}
}
