package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//这里我们循环的接受客户端发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write【发送】,那么协程阻塞在这里
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的Read err=", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听......")
	//1.tcp表示使用网络协议是tcp
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()
	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
			return
		}
		//这里准备一个协程，为客户端服务
		go process(conn)
	}
}
