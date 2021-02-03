package main

import (
	"fmt"
	"sync"
	"time"
)

/**
需求：计算1-200 各个数的阶乘，并且把各个数的阶乘放入到map中。
		最后显示出来

思路：1.编写一个函数，来计算各个数的阶乘，并让如到map中
		2.我们启动的协程多个，统计的将结果放入到map中
			3.map应该做一个全局的
*/

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁
	//sync 是包：  synchronized 同步
	lock sync.Mutex
)

func test(num int) {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	//这里我们将res放入到myMap
	//加锁
	lock.Lock()
	myMap[num] = res
	//解锁
	lock.Unlock()
}

//多进程
func main() {
	//开启多个协程完成这个任务
	for i := 1; i <= 10; i++ {
		go test(i)
	}

	time.Sleep(time.Second)
	//这里我们输出结果,
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
