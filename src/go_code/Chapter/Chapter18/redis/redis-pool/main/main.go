package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
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

func main() {
	//先从pool取出一个链接池
	conn1 := pool.Get()
	defer conn1.Close()
	_, err := conn1.Do("Set", "name", "tom cat")
	if err != nil {
		fmt.Println("conn1.Do err=", err)
		return
	}
	//取出
	r1, err := redis.String(conn1.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn1.Do err=", err)
		return
	}
	fmt.Println("r1=", r1)

	//如果要从pool取出链接，一定保证链接池是没有关闭
	//pool.Close()

	conn2 := pool.Get()
	defer conn2.Close()
	_, err = conn2.Do("Set", "job", "GoCoder")
	if err != nil {
		fmt.Println("conn2.Do err=", err)
	}
	r2, err := redis.String(conn2.Do("Get", "job"))
	if err != nil {
		fmt.Println("conn2.Do err=", err)
		return
	}
	fmt.Println("r2=", r2)

}
