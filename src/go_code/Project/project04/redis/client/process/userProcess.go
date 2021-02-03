package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/Project/project04/redis/common/message"
	"go_code/Project/project04/redis/common/utils"
	"net"
)

type UserProcess struct {
	//暂时不需要字段
}

//写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (loginResMes message.LoginResMes, err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return loginResMes, err
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return loginResMes, err
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return loginResMes, err
	}
	//fmt.Println("data=", string(data))

	//7.到这个时候，data就是我们要发送的消息
	//7.1 先把data的长度发送给服务器
	// 先获取到 data的长度-》转成一个表示长度的byte切片
	//conn.Write(len(data))
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes)")
		return loginResMes, err
	}
	//fmt.Printf("客户端，发送消息的长度=%d,内容=%s\n", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return loginResMes, err
	}
	tf := &utils.Transfer{Conn: conn}
	mes, err = tf.ReadPkg()
	//将mes的data部分反序列化
	if err != nil {
		fmt.Println("readPkg() err=", err)
		return loginResMes, err
	}

	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if err != nil {
		fmt.Println("json.Unmarshal() err=", err)
		return loginResMes, err
	}
	//这里我们还需要在客户端启动一个协程
	//该协程保持在和服务器端的通讯，如果服务器有数据推送给客户端，则接收并显示在客户端的终端
	if loginResMes.Code == 100 {
		//登录成功
		go ServerProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println("loginResMes.err=", loginResMes.Error)
	}
	return
}
