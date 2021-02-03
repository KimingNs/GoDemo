package main

import "fmt"

func main() {

	numChan := make(chan int, 2000)
	resChan := make(chan int, 2000)
	boolChan := make(chan bool, 2000)
	n := 2000

	go channelWrite(numChan)
	for i := 0; i < n; i++ {
		go channelRead(numChan, resChan, boolChan)
	}

	for {
		if len(boolChan) == n {
			close(resChan)
			break
		}
	}
	i := 1
	for v := range resChan {
		fmt.Printf("res[%v]=%v\n", i, v)
		i++
	}
}

func channelWrite(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		//放入数据
		numChan <- i
	}
	close(numChan)
}

func channelRead(numChan chan int, resChan chan int, boolChan chan bool) {
	sum := 0
	num := <-numChan
	for i := 1; i <= num; i++ {
		sum += i
	}
	resChan <- sum
	boolChan <- true
}
