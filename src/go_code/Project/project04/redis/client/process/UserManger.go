package process

import (
	"fmt"
	"go_code/Project/project04/redis/common/message"
)

//客户要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

//显示当前在线用户
func outputOnlineUser() {
	//遍历一把onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, user := range onlineUsers {
		fmt.Printf("用户id：%d,用户状态:%d\n", id, user.UserStatus)
	}
}

//编写一个方法，返回处理信息
func updateUserStatus(mes *message.NotifyUserStatusMes) {
	//适当优化
	user, ok := onlineUsers[mes.UserId]
	if !ok { //原来没有
		user = &message.User{UserId: mes.UserId}
	}
	user.UserStatus = mes.UserStatus
	onlineUsers[mes.UserId] = user

	outputOnlineUser()
}
