package main

import (
	"fmt"
	"runtime"
	"time"
)

//多进程
func main() {
	//go test2()
	//test1()
	num := testCPU()
	fmt.Println(num)
}

func testCPU() int {
	num := runtime.NumCPU()
	return num
}

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1 hello world =", i)
		time.Sleep(time.Second)
	}
}

func test2() {
	for i := 0; i < 5; i++ {
		fmt.Println("test2 hello golang =", i)
		time.Sleep(time.Second)
	}
}
