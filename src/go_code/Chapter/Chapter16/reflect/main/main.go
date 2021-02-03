package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量的type，kind，值
	//1。先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//n2 := 2 + rVal.Int()
	//fmt.Println(n2)
	//fmt.Printf("rValue=%v,type=%T\n", rVal, rVal)

	//下面将rVal转成 interface{}
	iv := rVal.Interface()
	fmt.Printf("iv = %v,iv = %T\n", iv, iv)
	//将interface{} 通过断言转成需要的类型
	//这里，我们就简单实用了一带检测的类型断言
	//同学们可以使用 switch 的断言形式来做的更加的灵活
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
	//num2 := iv.(int)
	//fmt.Println("num2=", num2)

}

//专门演示反射[对结构体的反射]
func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的 type ， kind ， 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//3.获取 变量对应的Kind
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	//(2) rTyp.Kind() ==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind = %v kind = %v\n", kind1, kind2)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//请编写一个案例：
	//演示对（基本数据类型，Interface{},reflect.Value)进行反射的基本操作

	//1.先定义一个int
	//var num int = 100

	//2.定义一个Student的实例
	stu := Student{Name: "Tom", Age: 10}
	//reflectTest01(stu)
	reflectTest02(stu)
}
