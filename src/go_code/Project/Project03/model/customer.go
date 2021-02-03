package model

import "fmt"

//声明一个Customer结构体
type CustomerModelStruct struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Mobile string
	Email  string
}

//编写工厂模式
func NewCustomer(id int, name string, gender string, age int, mobile string, email string) CustomerModelStruct {
	return CustomerModelStruct{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Mobile: mobile,
		Email:  email,
	}
}

//返回用户的信息
func (this CustomerModelStruct) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Mobile, this.Email)
	return info
}
