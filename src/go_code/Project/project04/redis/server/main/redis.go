package main

import "github.com/garyburd/redigo/redis"

//定义一个全局的pool
var redisPool *redis.Pool

func initPool() {
	redisPool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			_ = c.Send("auth", "callmens1005")
			return c, err
		},
	}
}
