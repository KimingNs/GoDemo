package controller

import (
	"go_code/Project/Project03/model"
)

//完成对Customer的操作
type CustomerControllerStruct struct {
	customers []model.CustomerModelStruct
	//声明一个字段，表示当前字段含有多少个客户
	customerNum int
}

func NewCustomController() *CustomerControllerStruct {
	//为了能够看到有客户在切片中，我们初始化一个客户
	cC := &CustomerControllerStruct{}
	cC.customerNum = 1
	customer := model.NewCustomer(1, "zhangsan", "nan", 18, "123456", "123456@qq.com")
	cC.customers = append(cC.customers, customer)
	return cC
}

//返回客户切片
func (this *CustomerControllerStruct) List() []model.CustomerModelStruct {
	return this.customers
}

func (this *CustomerControllerStruct) Add(customer model.CustomerModelStruct) bool {
	this.customers = append(this.customers, customer)
	this.customerNum++
	return true
}

func (this *CustomerControllerStruct) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)

	this.customerNum--
	return true
}

func (this *CustomerControllerStruct) Change(id int, name string, gender string, age int, mobile string, email string) bool {
	cm := this.FindModelById(id)
	cm.Name = name
	cm.Gender = gender
	cm.Age = age
	cm.Mobile = mobile
	cm.Email = email
	return true
}

func (this *CustomerControllerStruct) FindModelById(id int) *model.CustomerModelStruct {
	index := this.FindById(id)
	return &this.customers[index]
}

func (this *CustomerControllerStruct) FindById(id int) int {
	index := -1
	for i := 0; i < this.customerNum; i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}
