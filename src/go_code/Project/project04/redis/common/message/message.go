package message

const (
	LoginMesType            = "LoginMes"
	LoginResType            = "LoginRes"
	RegisterMesType         = "RegisterMes"
	RegisterResType         = "RegisterRes"
	NotifyUserStatusMesType = "NotifyOthersMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的类型
}

type ResStruct struct {
	Code  int    `json:"code"`  //返回状态码，100表示登录成功
	Error string `json:"error"` //返回错误信息
}

//定义两个消息。。后面需要再增加
type LoginMes struct {
	User User `json:"user"`
}

type LoginMesRes struct {
	Res    ResStruct `json:"res"`
	UserId []int     `json:"user_online"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterMesRes struct {
	Res ResStruct `json:"res"`
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId     int `json:"user_id"`
	UserStatus int `json:"user_status"`
}
