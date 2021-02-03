package main

import (
	"fmt"
	"go_code/Project/project04/redis/client/process"
)

//定义两个变量，一个表示用户id，一个表示用户的密码
var userId int
var userPwd string

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
			up := &process.UserProcess{}
			_, err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("up.Login() err=", err)
			}
		case 2:
			fmt.Println("注册用户")
			loop = false
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
