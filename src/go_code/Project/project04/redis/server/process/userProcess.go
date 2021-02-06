package process

import (
	"encoding/json"
	"fmt"
	"go_code/Project/project04/redis/client/model"
	"go_code/Project/project04/redis/common/message"
	"go_code/Project/project04/redis/common/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn `json:"conn"`
	//增加一个字段表示该conn 属于哪个用户
	UserId int `json:"user_id"`
}

//这里我们编写通知所有在线用户的方法
//userId 要通知其他
func (this *UserProcess) NotifyOthers() {
	//遍历UsersOnline，然后一个个的发送 NotifyUserStatusMes
	userMap := userManger.GetUserOnlineAll()
	for id, up := range userMap {
		if id == this.UserId {
			continue
		}
		up.Notify(up.UserId)
	}
}

func (this *UserProcess) Notify(userId int) {
	//组装我们的NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifuUserStatusMes message.NotifyUserStatusMes
	notifuUserStatusMes.UserId = userId
	notifuUserStatusMes.UserStatus = message.ZAIXIAN

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifuUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.data
	mes.Data = string(data)

	//对mes 再次进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送，创建我们Transfer实例，发送
	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyMe err=", err)
		return
	}
	return
}

//编写一个ServerProcessLogin 函数
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	//核心代码...
	//1.先从mes 中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResType

	//在声明一个LoginResMes
	var loginMesRes message.LoginMesRes

	//我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao 到redis去验证
	user, err := model.MyUserDao.Login(loginMes.User.UserId, loginMes.User.UserPwd)
	if err != nil {
		if err == message.ERROR_USER_NOT_EXISTS {
			loginMesRes.Res.Code = message.USER_NOT_EXISTS
		} else if err == message.ERROR_USER_PWD_FALSE {
			loginMesRes.Res.Code = message.USER_PWD_FALSE
		} else {
			loginMesRes.Res.Code = message.ERROR
		}
		loginMesRes.Res.Error = err.Error()
	} else {
		loginMesRes.Res.Code = message.SUCCESS
		//这里用户登录成功，我们就把该登陆成功的用户放入到userManger中
		//将登录成功的用户的userId 赋给this
		this.UserId = loginMes.User.UserId
		userManger.AddUserOnline(this)

		//通知其他用户该登录用户上线了
		this.NotifyOthers()

		//将当前在线用户的id放入到loginResMes的userId
		//遍历userManger.UserOnline
		for id, _ := range userManger.UsersOnline {
			loginMesRes.UserId = append(loginMesRes.UserId, id)
		}

		fmt.Println("user=", user)
	}

	//将loginResMes序列化
	data, err := json.Marshal(loginMesRes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMes.Data = string(data)

	//对resMes进行序列化,准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//发送data,我们将其封装到writePkg函数
	//因为使用分层模式，我们先创建一个Transfer实例，然后来读取
	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//核心代码...
	//1.先从mes 中取出mes.Data，并直接反序列化成LoginMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResType
	//2.在声明一个LoginResMes，并完成赋值

	//在声明一个LoginResMes
	var registerMesRes message.RegisterMesRes

	//我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao 到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == message.ERROR_USER_EXISTS {
			registerMesRes.Res.Code = message.USER_EXISTS
		} else {
			registerMesRes.Res.Code = message.ERROR
		}
		registerMesRes.Res.Error = err.Error()
	} else {
		registerMesRes.Res.Code = message.SUCCESS
		fmt.Println("user=", registerMes)
	}

	//将loginResMes序列化
	data, err := json.Marshal(registerMesRes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMes.Data = string(data)

	//对resMes进行序列化,准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//发送data,我们将其封装到writePkg函数
	//因为使用分层模式，我们先创建一个Transfer实例，然后来读取
	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	return
}
