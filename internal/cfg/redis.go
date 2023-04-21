package cfg

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type RedisCfg struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var redisConfig RedisCfg

const redisCfgFileName = "redis"

func init() {
	cfgBytes, err := readCfgBytes(redisCfgFileName)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(cfgBytes, &redisConfig); err != nil {
		panic(err.Error())
	}
	fmt.Printf("init redisConfig: %v \n", redisConfig)
}
func LoadRedisCfg() *RedisCfg {
	return &redisConfig
}
