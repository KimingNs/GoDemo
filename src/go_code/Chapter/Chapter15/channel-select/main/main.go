package main

import (
	"fmt"
	"time"
)

func main() {

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second * 2)
	}

	return
	//使用select可以解决从管道取数据的阻塞问题

	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	//问题，在实际开发中，可能我们不好确定什么时候关闭该管道
	//可以使用select 方法
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据=%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取到的数据=%s\n", v)
		default:
			fmt.Println("都取不到")
		}
	}
}

func test() {

	//这里可以使用错误处理机制 defer + recover
	defer func() {
		//捕获抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误")
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang"
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}
