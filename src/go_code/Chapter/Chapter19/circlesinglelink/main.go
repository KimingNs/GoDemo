package main

import "fmt"

type CatNode struct {
	no   int //猫的编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //形成一个环状
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	//先定义一个临时的变量，帮忙找到环形的最后一个结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("猫的信息为=[id=%d,name=%s] ->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	//这里我们初始化一个环形链表的头结点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	InsertCatNode(head, cat1)
	ListCircleLink(head)
}
