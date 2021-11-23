package cache

import (
	"github.com/ipeace/go-web-server/cache/logic"
)

type Config struct {
	Option logic.ICache
}

var client logic.ICache

func NewCache(config *Config) {
	client = config.Option

	err := client.Run()

	if err != nil {
		panic("Cache New Error：" + err.Error())
	}
}
