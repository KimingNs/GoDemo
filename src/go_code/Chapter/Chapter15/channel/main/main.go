package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func test1() {
	//演示一下channel的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2.看看intChan是什么
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	//3.向channel写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 5

	//4.看看channel的长度和cap容量
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	fmt.Printf("channel len=%v,cap=%v\n", len(intChan), cap(intChan))

	num2 := <-intChan
	fmt.Printf("num2=%v\n", num2)
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	fmt.Printf("channel len=%v,cap=%v\n", len(intChan), cap(intChan))

	fmt.Println("----------------------------------")

	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string)
	m1["city1"] = "beijing"
	m1["city2"] = "wuhan"
	m1["city3"] = "shanghai"

	m2 := make(map[string]string)
	m2["human1"] = "zhangsan"
	m2["human2"] = "lisi"
	m2["human3"] = "wangwu"

	fmt.Printf("channel len=%v,cap=%v\n", len(mapChan), cap(mapChan))
	mapChan <- m1
	mapChan <- m2
	fmt.Printf("channel len=%v,cap=%v\n", len(mapChan), cap(mapChan))
	test := <-mapChan
	fmt.Println(test)
}

func test2() {
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"cat", 10}
	allChan <- cat

	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("newCat=%T,newCat=%v\n", newCat, newCat)
	//类型断言
	value, ok := newCat.(Cat)
	if ok == true {
		fmt.Printf("newCat.Name= %v\n", value.Name)
	}

}

func test3() {
	//channel关闭
	//intChan := make(chan int, 100)
	//for i := 0; i < 100; i++ {
	//	if i == 10 {
	//		close(intChan)
	//	}
	//	intChan <- i
	//}
	//fmt.Printf("channel len=%v,cap=%v\n", len(intChan), cap(intChan))

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i
	}
	close(intChan2)
	//不关闭channel会报错 dead lock
	for v := range intChan2 {
		fmt.Println("v=", v)
	}

}

func work1() {

}

//writeData
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan <- i
		fmt.Printf("writeData 写入数据=%v\n", i)
	}
	close(intChan)
}

//readData
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)

}

func main() {
	//work1
	//创建两个channel
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
