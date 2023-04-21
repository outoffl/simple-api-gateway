package cfg

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

var cfg *Cfg

type Cfg struct {
	Proxy map[string][]string `yaml:"proxy"`
}

const proxyCfgFileName = "proxy"

func init() {
	cfgBytes, err := readCfgBytes(proxyCfgFileName)
	if err != nil {
		panic(err.Error())
	}
	if err := yaml.Unmarshal(cfgBytes, &cfg); err != nil {
		panic(err.Error())
	}
	fmt.Printf("init cfg: %v \n", cfg)
}

func LoadCfg() Cfg {
	return *cfg
}
