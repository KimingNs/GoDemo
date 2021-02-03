package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Send("auth", "callmens1005")
	if err != nil {
		fmt.Println(err)
	}
	RedisSet(c)
	RedisHSet(c)
	defer c.Close()
}

func RedisSet(c redis.Conn) {
	_, _ = redis.String(c.Do("set", "key5", "test5"))
	re, _ := redis.String(c.Do("get", "key2"))
	//re, err := redis.String(c.Do("get", "key4"))
	fmt.Println(re)
}

func RedisHSet(c redis.Conn) {
	_, _ = redis.String(c.Do("HSet", "user01", "name", "Tom", "age", "10", "address", "wuhan"))
	//re, _ := redis.String(c.Do("HGet", "user01", "address"))
	re, err := redis.Strings(c.Do("HGetAll", "user01"))
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println(re)
	for i, v := range re {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
