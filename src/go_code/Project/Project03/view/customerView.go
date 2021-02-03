package main

import (
	"fmt"
	"go_code/Project/Project03/controller"
	"go_code/Project/Project03/model"
)

type customerViewStruct struct {
	key  int
	loop bool
	id   int

	customerController *controller.CustomerControllerStruct
}

func (this *customerViewStruct) list() {
	customers := this.customerController.List()
	//显示
	fmt.Println("----------------客户列表----------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t 邮箱\t")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------------客户列表完成----------------")
	fmt.Println()
	fmt.Println()
}

//得到用户的输入信息，构建新的客户，并完成添加
func (this *customerViewStruct) add() {
	fmt.Println("----------------添加客户----------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	mobile := ""
	fmt.Scanln(&mobile)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	this.id = 1 + this.id
	id := this.id + 1
	//构建新的Customer实例
	customer := model.NewCustomer(id, name, gender, age, mobile, email)
	if this.customerController.Add(customer) {
		fmt.Println("添加完成")
	} else {
		fmt.Println("添加失败")
	}

}

func (this *customerViewStruct) showInfo() {

}

func (this *customerViewStruct) delete() {
	fmt.Println("----------------删除客户----------------")
	fmt.Println("请选择要删除的客户编号（-1退出）")
	id := -2
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认要删除吗（Y/N）")
	for {
		var res string
		fmt.Scanln(&res)
		if res == "N" {
			fmt.Println("已取消删除")
			break
		}
		if res == "Y" {
			if this.customerController.Delete(id) {
				fmt.Println("---------------删除完成----------------")
			} else {
				fmt.Println("客户编号输出错误，请重新输入")
			}
			break
		}
	}
}

func (this *customerViewStruct) change() {
	fmt.Println("---------------修改客户----------------")
	fmt.Println("请选择要修改的客户编号(-1退出)：")
	id := -2
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	test := this.customerController.FindById(id)
	if test == -1 {
		fmt.Println("无该客户信息")
		return
	}
	res := this.customerController.FindModelById(id)

	var name string
	fmt.Println("姓名(", res.Name, "):")
	_, _ = fmt.Scanln(&name)
	if name == "" {
		name = res.Name
	}
	var gender string
	fmt.Println("性别(", res.Gender, "):")
	_, _ = fmt.Scanln(&gender)
	if gender == "" {
		gender = res.Gender
	}
	var age int
	fmt.Println("年龄(", res.Age, "):")
	_, _ = fmt.Scanln(&age)
	if age == 0 {
		age = res.Age
	}
	var mobile string
	fmt.Println("手机号(", res.Mobile, "):")
	_, _ = fmt.Scanln(&mobile)
	if mobile == "" {
		mobile = res.Mobile
	}
	var email string
	fmt.Println("邮箱(", res.Email, "):")
	_, _ = fmt.Scanln(&email)
	if email == "" {
		email = res.Email
	}
	if !this.customerController.Change(id, name, gender, age, mobile, email) {
		fmt.Println("修改失败")
	}
}

//显示主菜单
func (cv *customerViewStruct) mainMenu() {
	for {
		fmt.Println("----------------客户信息管理软件----------------")
		fmt.Println("				1	添	加	客	户")
		fmt.Println("				2	修	改	客	户")
		fmt.Println("				3	删	除	客	户")
		fmt.Println("				4	客	户	列	表")
		fmt.Println("				5	退			出")
		fmt.Println("请选择（1-5）：")

		_, _ = fmt.Scanln(&cv.key)
		switch cv.key {
		case 1:
			cv.add()
		case 2:
			cv.change()
		case 3:
			cv.delete()
		case 4:
			cv.list()
		case 5:
			cv.loop = false
		default:
			fmt.Println("输入有误")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("已退出程序")
}

func main() {
	cv := customerViewStruct{
		key:  0,
		loop: true,
	}
	//完成对customerViewStruct结构体的初始化
	cv.customerController = controller.NewCustomController()
	cv.mainMenu()
}
