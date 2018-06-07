package config

import (
	"fmt"
	"github.com/widuu/goini"
	// "runtime"
)

const (
	config_file = "/Users/lixiumeng/go/src/Sephiroth/config/config.ini"
)

// ini配置解析器
type ConfigParser struct {
}

func NewConfigParser() *ConfigParser {
	cf := new(ConfigParser)
	return cf
}

func (cf *ConfigParser) GetConfig() []map[string]map[string]string {
	return goini.SetConfig(config_file).ReadList()
}

func (cf *ConfigParser) Test() {
	fmt.Println(cf.GetConfig())
}
