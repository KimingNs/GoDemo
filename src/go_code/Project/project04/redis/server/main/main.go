package main

import (
	"fmt"
	"go_code/Project/project04/redis/client/model"
	"net"
)

func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//读客户端发送的信息
	//for {
	//这里我们将读取整个数据包，直接封装成一个函数readPkg()，返回Message,Err
	//这里调用总控，创建一个
	processor := &Processor{Conn: conn}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误 , err=", err)
		return
	}
	//}
}

//这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	//这里需要注意一个初始化顺序的问题
	model.MyUserDao = model.NewUserDao(redisPool)
}

func init() {
	//当服务器启动时，我们就去初始化我们的redis连接池
	initPool()
	initUserDao()
}

func main() {

	//提示信息
	fmt.Println("服务器在8889监听。。。。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
