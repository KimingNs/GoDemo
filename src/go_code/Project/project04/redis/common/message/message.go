package message

const (
	LoginMesType       = "LoginMes"
	LoginResType       = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的类型
}

//定义两个消息。。后面需要再增加
type LoginMes struct {
	User User `json:"user"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态码，100表示登录成功
	Error string `json:"error"` //返回错误信息
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
