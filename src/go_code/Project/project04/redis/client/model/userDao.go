package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go_code/Project/project04/redis/common/message"
)

//定义一个UserDao结构体
//完成对User 结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

//我们把服务器启动后，初始化一个UserDao实例，把他做成全局变量，在需要和redis操作时就直接使用即可
var (
	MyUserDao *UserDao
)

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{pool: pool}
	return
}

//思考一下UserDao应该提供哪些方法给我们
//1.根据用户id 返回一个User实例 + err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user message.User, err error) {
	//通过给定id 去 redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = message.ERROR_USER_NOT_EXISTS
		}
		return
	}

	//需要把res 反序列化成User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//完成对登录的校验
//1.Login 完成对用户的验证
//2.如果用户的id和pwd都正确，则返回一个user实例
//3.如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user message.User, err error) {
	//先从UserDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//这时证明这个用户时获取到
	if user.UserPwd != userPwd {
		err = message.ERROR_USER_PWD_FALSE
		return
	}
	return
}

//完成对注册的校验
func (this *UserDao) Register(user *message.User) (err error) {
	//先从UserDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = message.ERROR_USER_EXISTS
		return
	}
	//这时，说明id在redis还没有，则可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}
