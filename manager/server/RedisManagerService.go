//author: richard
package server

import (
	"fmt"
	lm "github.com/advancevillage/letaotao/manager"
	"github.com/go-redis/redis"
	"sync"
)

var cacheConn 		*redis.Client
var cacheConfig 	*lm.CacheConfig
var cacheErr	 	error
var cacheOnce 		sync.Once

func InitCache(c *lm.CacheConfig) error {
	cacheOnce.Do(func () {
		cacheConfig = c
		cacheConn = redis.NewClient(&redis.Options{Addr:fmt.Sprintf("%s:%s",cacheConfig.Host,cacheConfig.Port),Password:cacheConfig.Token,DB: cacheConfig.Schema})
		_, cacheErr = cacheConn.Ping().Result()
	})
	return cacheErr
}

func ConnCache() *redis.Client {
	cacheOnce.Do(func (){
		cacheConn = redis.NewClient(&redis.Options{Addr:fmt.Sprintf("%s:%s",cacheConfig.Host,cacheConfig.Port),Password:cacheConfig.Token,DB: cacheConfig.Schema})
		_, cacheErr = cacheConn.Ping().Result()
	})
	return cacheConn
}
