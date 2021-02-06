package message

import "errors"

//根据业务逻辑的需要，自定义一些错误
var (
	ERROR_USER_NOT_EXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS     = errors.New("用户已存在")
	ERROR_USER_PWD_FALSE  = errors.New("密码错误")
)

//code码
var (
	SUCCESS         = 100
	ERROR           = 900
	USER_NOT_EXISTS = 101
	USER_EXISTS     = 102
	USER_PWD_FALSE  = 103
)
