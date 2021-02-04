package main

import (
	"fmt"
	"go_code/Project/project04/redis/common/message"
	"go_code/Project/project04/redis/common/utils"
	process2 "go_code/Project/project04/redis/server/process"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn `json:"conn"`
}

//编写一个ServerProcessMes 函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		up := &process2.UserProcess{Conn: this.Conn}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{Conn: this.Conn}
		err = up.ServerProcessRegister(mes)
		//err = up.
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	defer this.Conn.Close()
	return
}

func (this *Processor) process2() (err error) {
	//循环的客户端发送的信息
	for {
		tf := &utils.Transfer{Conn: this.Conn}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出")
				return err
			} else {
				fmt.Println("readPkg() err=", err)
				return err
			}
		}
		err = this.serverProcessMes(&msg)
		if err != nil {
			return err
		}
	}
}
