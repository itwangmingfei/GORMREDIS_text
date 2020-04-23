package tools

import (
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

type Goredis struct {
	client *redis.Client
}

var GOREDIS *Goredis

func InitRedisStore() *Goredis{
	config := Cfg.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:  config.Addr+":"+config.Port,
		Password: config.Password,
		DB: 0,
	})
	GOREDIS = &Goredis{client:client}
	return GOREDIS
}
//设置无过期时间
func(red *Goredis) Set(key,val string){
	err := red.client.Set(key,val,0).Err()
	if err!=nil{
		log.Println(err)
	}
}
//有过期时间
func (red *Goredis) SetT(key,val string,outtime time.Duration)  {
	err := red.client.Set(key,val,outtime).Err()
	if err!=nil{
		log.Println(err)
	}
}
//获取值
func(red *Goredis) Get(key string) string{
	val,err := red.client.Get(key).Result()
	if err!=nil{
		log.Println(err)
		return ""
	}
	return val
}
//获取值后删除
func(red *Goredis) GetC(key string,clear bool) string{
	val,err := red.client.Get(key).Result()
	if err!=nil{
		log.Println(err)
		return ""
	}
	if clear{
		err := red.client.Del(key).Err()
		if err!=nil{
			log.Println(err)
			return ""
		}
	}
	return val
}