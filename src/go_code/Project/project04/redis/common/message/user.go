package message

//定义一个用户的结构体
type User struct {
	//确定字段信息
	//为了序列化和反序列化成功，我们必须保证用户信息的json字符串的key 和结构体的字段对应的tag名字一致
	UserId     int    `json:"user_id"`
	UserPwd    string `json:"user_pwd"`
	UserName   string `json:"user_name"`
	UserStatus int    `json:"status"`
}

//用户当前状态code对应值
const (
	LIXIAN  = 0
	ZAIXIAN = 1
	MANGLU  = 2
	YINSHEN = 3
)
