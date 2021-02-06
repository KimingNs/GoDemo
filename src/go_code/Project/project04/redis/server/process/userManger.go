package process

import "fmt"

//因为UserManger 实例在服务器端有且只有一个
//因为在很多的地方都会使用到，因此我们将其定位全局变量

var (
	userManger *UserManger
)

type UserManger struct {
	UsersOnline map[int]*UserProcess
}

//完成对userManger的初始化工作
func init() {
	userManger = &UserManger{UsersOnline: make(map[int]*UserProcess, 1024)}
}

//完成对UsersOnline添加
func (this *UserManger) AddUserOnline(up *UserProcess) {
	this.UsersOnline[up.UserId] = up
}

//删除
func (this *UserManger) DeleteUserOnline(userId int) {
	delete(this.UsersOnline, userId)
}

//返回当前所有在线的用户
func (this *UserManger) GetUserOnlineAll() map[int]*UserProcess {
	return this.UsersOnline
}

//返回id对应的在线用户
func (this *UserManger) GetUserOnline(userId int) (up *UserProcess, err error) {
	//如何从map取出一个值，带检测方式
	up, ok := this.UsersOnline[userId]
	if !ok { //说明你要找的用户当前不在线
		err = fmt.Errorf("用户%d 当前不在线", userId)
		return
	}
	return
}
