package main

import (
	"fmt"
	"go_code/Project/project04/redis/client/process"
)

/**
问题未解决：客户端完成登录操作后无法同时开启端口对服务器端传输数据导致端口关闭
*/

//定义两个变量，一个表示用户id，一个表示用户的密码
var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	var loop = true
	for {
		fmt.Println("--------欢迎登陆多人聊天系统--------")
		fmt.Println("\t\t\t\t1.登录聊天系统")
		fmt.Println("\t\t\t\t2.注册用户")
		fmt.Println("\t\t\t\t3.退出系统")
		fmt.Println("请选择（1-3）：")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)

			//说明用户要登录
			//先把登录函数写到另外一个文件，比如login.go
			//1.创建一个UserProcess实例
			up := &process.UserProcess{}
			_, err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("up.Login() err=", err)
			}
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名：")
			fmt.Scanf("%s\n", &userName)
			//2.调用UserProcess实例，完成注册的请求
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
		if loop == false {
			break
		}
	}
}
