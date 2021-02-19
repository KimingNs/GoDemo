package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	maxSize int
	array   [5]int //数组 => 模拟队列
	front   int    //表示指向队列首部
	rear    int    // 表示指向队列尾部
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

//从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	val = this.array[this.rear]
	this.rear--
	return val, err
}

//显示队列，找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	//this.front不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, this.array[i])
	}
}

func main() {
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 表示退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数：")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err == nil {
				fmt.Println("入列成功！")
			} else {
				fmt.Println(err.Error())
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("当前队列队尾数字为：", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			break
		}
	}
}
