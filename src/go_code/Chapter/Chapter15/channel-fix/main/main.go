package main

import "fmt"

func main() {
	//channel 可以声明为只读或只写

	//1.在默认情况下，管道为双向
	//var chan1 chan int//可读可写

	//2.声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 10)
	chan2 <- 100
	close(chan2)
	//num := <-chan2

	fmt.Println(chan2)

	//3.声明为只读
	var chan3 <-chan int

	_, num2 := <-chan3
	fmt.Println(num2)

}

func test() {
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go write(ch, exitChan)
	go read(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束")

}

func write(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

func read(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}
