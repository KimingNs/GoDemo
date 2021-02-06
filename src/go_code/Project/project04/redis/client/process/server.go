package process

import (
	"encoding/json"
	"fmt"
	"go_code/Project/project04/redis/common/message"
	"go_code/Project/project04/redis/common/utils"
	"net"
	"os"
)

//显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("--------恭喜xxx登陆成功")
	fmt.Println("--------1.显示在线用户列表")
	fmt.Println("--------2.发送信息")
	fmt.Println("--------3.信息列表")
	fmt.Println("--------4.退出系统")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送信息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确")
	}
}

//和服务器保持通讯
func ServerProcessMes(conn net.Conn) {
	//创建一个transfer实例
	//1.
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err=", err)
			return
		}
		//如果读取到消息，又是下一步处理逻辑
		switch mes.Type {
		//有人上线了
		case message.NotifyUserStatusMesType:
			//1.取出 NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("59json.Unmarshal() err=", err)
			}
			//2.把这个用户的信息，状态保存到客户map[int]User中
			updateUserStatus(&notifyUserStatusMes)
		default:
			fmt.Println("mes=", mes)
		}
	}
}
