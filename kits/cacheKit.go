package kits

import (
	"sync"
)

type cacheKit struct {
	Config map[string]interface{}
}

var once sync.Once
var instance *cacheKit
var config_filename = "/conf.ini"

func GetInstance() *cacheKit{
	once.Do(func(){
		conf,_ := LoadFile(config_filename)
		instance = &cacheKit{
			Config : conf,
		}
	})
	return instance
}
