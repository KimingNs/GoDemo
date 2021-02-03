package main

import (
	"fmt"
)

func init() {
	fmt.Println("---------begin---------")
}
func main() {
	var x interface{}
	var b2 float32 = 2.1
	x = b2

	y, ok := x.(float64)
	if ok {
		fmt.Println("y:", y)
	} else {
		fmt.Println("false")
	}
	fmt.Println("继续执行")
}
