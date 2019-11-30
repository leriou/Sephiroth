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
	once.Do(func() {
		di.Config = config.NewConfigParser().GetConfig()
	})
	return di.Config
}

func (di *Di) GetRedis() *redis.Client {
	once.Do(func() {
		rdsConf := di.Config[1]["redis"]
		rdsConfig := &redis.Options{
			Addr:     rdsConf["host"] + ":" + rdsConf["port"],
			Password: "", // no password set
			DB:       0,  // use default DB
		}
		rdsCli = redis.NewClient(rdsConfig)
	})
	return rdsCli
}

func (di *Di) GetMongoDB() *mgo.Session {
	once.Do(func() {
		mgoConf := di.Config[2]["mongodb"]
		url := mgoConf["host"] + ":" + mgoConf["port"]
		mgoCli, _ = mgo.Dial(url)
	})
	return mgoCli
}

func (di *Di) GetElastic() *elastic.Client {
	once.Do(func() {
		esCli, _ = elastic.NewClient()
	})
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
