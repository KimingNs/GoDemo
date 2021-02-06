package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/Project/project04/redis/common/message"
	"go_code/Project/project04/redis/common/utils"
	"net"
	"os"
)

type UserProcess struct {
	//暂时不需要字段
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (registerMesRes message.RegisterMesRes, err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Register.net.Dial() err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务
	var mes message.Message
	var registerMes message.RegisterMes
	mes.Type = message.RegisterMesType
	//3.创建一个User 结构体
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将loginMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("38Register.json.Marshal() err=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("47Register.json.Marshal() err=", err)
		return
	}

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("58Register.tf.WritePkg() err=", err)
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("63Register.tf.ReadPkg() err=", err)
		return
	}

	//将mes的Data部分反序列化成RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerMesRes)
	if registerMesRes.Res.Code == message.SUCCESS {
		fmt.Println("注册成功")
	} else {
		fmt.Println("registerResMes.error=", registerMesRes.Res.Error)
	}
	os.Exit(0)
	return
}

//写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (loginMesRes message.LoginMesRes, err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("82Login.net.Dial() err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.User.UserId = userId
	loginMes.User.UserPwd = userPwd

	//4.将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("100Login.json.Marshal() err=", err)
		return
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("110Login.json.Marshal() err=", err)
		return
	}
	//fmt.Println("data=", string(data))

	//7.到这个时候，data就是我们要发送的消息
	//7.1 先把data的长度发送给服务器
	// 先获取到 data的长度-》转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("125Login.conn.Write() err=", err)
		return
	}
	//fmt.Printf("客户端，发送消息的长度=%d,内容=%s\n", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("132Login.conn.Write() err=", err)
		return loginMesRes, err
	}
	tf := &utils.Transfer{Conn: conn}
	mes, err = tf.ReadPkg()
	//将mes的data部分反序列化
	if err != nil {
		fmt.Println("139Login.readPkg() err=", err)
		return
	}

	err = json.Unmarshal([]byte(mes.Data), &loginMesRes)
	if err != nil {
		fmt.Println("146Login.json.Unmarshal() err=", err)
		return
	}
	//这里我们还需要在客户端启动一个协程
	//该协程保持在和服务器端的通讯，如果服务器有数据推送给客户端，则接收并显示在客户端的终端
	if loginMesRes.Res.Code == message.SUCCESS {
		//登录成功
		//可以显示当前在线用户列表，遍历loginResMes.UserId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginMesRes.UserId {
			//如果我们要求不显示自己在线，
			if v == userId {
				continue
			}
			fmt.Printf("用户id：%d\n", v)
			//完成对 客户端的onlineUsers 的初始化
			user := &message.User{UserId: v, UserStatus: message.ZAIXIAN}
			onlineUsers[v] = user

		}
		fmt.Println()
		fmt.Println()

		go ServerProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println("loginResMes.err=", loginMesRes.Res.Error)
	}
	return
}
