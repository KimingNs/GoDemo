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
}

//编写一个ServerProcessLogin 函数
func (this *UserProcess) ServrProcessLogin(mes *message.Message) (err error) {

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
	var loginResMes message.LoginResMes

	//我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao 到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		loginResMes.Code = 101
		loginResMes.Error = err.Error()
	} else {
		loginResMes.Code = 100
		fmt.Println("user=", user)
	}

	//将loginResMes序列化
	data, err := json.Marshal(loginResMes)
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
