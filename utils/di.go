package utils

import (
	"fmt"
	"sephiroth/config"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
)

var (
	di     *Di
	rdsCli *redis.Client
	mgoCli *mgo.Session
	esCli  *elastic.Client
	start  int64
	end    int64
	once   sync.Once
)

// 依赖注入
type Di struct {
	Config []map[string]map[string]string
}

func NewDi() *Di {
	once.Do(func() {
		di := new(Di)
		di.Config = di.GetConfig()
	})
	return di
}

func (di *Di) GetConfig() []map[string]map[string]string {
	return config.NewConfigParser().GetConfig()
}

func (di *Di) GetRedis() *redis.Client {
	if rdsCli == nil {
		rdsConf := di.Config[1]["redis"]
		rdsConfig := &redis.Options{
			Addr:     rdsConf["host"] + ":" + rdsConf["port"],
			Password: "",
			DB:       0,
		}
		rdsCli = redis.NewClient(rdsConfig)
	}
	return rdsCli
}

func (di *Di) GetMongoDB() *mgo.Session {
	if mgoCli == nil {
		mgoCli, _ = mgo.Dial("localhost:27017")
	}
	return mgoCli
}

func (di *Di) GetElastic() *elastic.Client {
	if esCli == nil {
		esCli, _ = elastic.NewClient()
	}
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
