package main

import (
	"fmt"
	"reflect"
)

//通过反射修改 值
func reflectTest01(b interface{}) {
	rVal := reflect.ValueOf(b)
	rVal.SetInt(20)
}

func main() {
	var num int = 1
	reflectTest01(&num)
	fmt.Println("num=", num)
}
