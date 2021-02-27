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

//删除
func DeleteCatNode(head *CatNode, id int) *CatNode {
	//删除一个环形单向链表的思路
	//1.先让temp指向head
	//2.让helper指向环形链表最后
	//3.让temp和要删除的id进行比较，如果相同，则同helper完成删除[这里必须考虑]

	temp := head
	helper := head
	//判断空结点
	if temp.next == nil {
		fmt.Println("该循环链表为空")
		return head
	}

	//只有头结点时
	if temp.next == head && temp.no == id {
		head = &CatNode{}
		return head
	}

	//删除头结点时
	if head.no == id {
		for {
			if helper.next == head {
				break
			}
			helper = helper.next
		}
		head = temp.next
		helper.next = head
	} else {
		//不是头结点时
		for {
			if helper.next.no == id {
				helper = helper.next
				break
			}
			if helper.next == head {
				fmt.Println("该循环链表上没有这个结点")
				break
			}
			helper = helper.next
			temp = temp.next
		}
		temp.next = helper.next
	}
	return head
}

//输出这个环形的链表
func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[id=%d,name=%s] ->", temp.no, temp.name)
		if temp.next == head {
			fmt.Printf("[id=%d name=%s]", head.no, head.name)
			break
		}
		temp = temp.next
	}
}

func main() {
	////这里我们初始化一个环形链表的头结点
	//head := &CatNode{}
	//
	////创建一只猫
	//cat1 := &CatNode{
	//	no:   1,
	//	name: "tom",
	//}
	//
	//cat2 := &CatNode{
	//	no:   2,
	//	name: "tom2",
	//}
	//
	//cat3 := &CatNode{
	//	no:   3,
	//	name: "tom3",
	//}
	//InsertCatNode(head, cat1)
	//InsertCatNode(head, cat2)
	//InsertCatNode(head, cat3)
	//ListCircleLink(head)
	//fmt.Println()
	//head = DeleteCatNode(head, 1)
	//ListCircleLink(head)

	//这里我们初始化一个环形链表的头结点
	head := &CatNode{}

	for i := 1; i <= 17; i++ {
		cat := &CatNode{
			no: i,
		}
		InsertCatNode(head, cat)
	}

	Josephu(head, 4, 3)
}

//约瑟夫问题
//n个人 从k开始从1数数,数到m的人出列，从m的下一个人开始从1数数......
func Josephu(head *CatNode, k int, m int) {
	//代表人数
	temp := head
	//记录被删的
	no := 0
	//记录被删的下一个
	no2 := 0
	for {
		if temp.next.no == k {
			break
		}
		temp = temp.next
	}

	for {
		ListCircleLink(head)
		fmt.Println()
		if head.next == head {
			fmt.Println("约瑟夫问题循环结束")
			break
		}
		if no != 0 {
			temp = head
			for {
				if temp.next.no == no2 {
					break
				}
				temp = temp.next
			}
		}
		for i := 1; i < m; i++ {
			temp = temp.next
		}
		no = temp.next.no
		no2 = temp.next.next.no
		fmt.Println("该次去除的编号为：", no)
		head = DeleteCatNode(head, no)
	}
}
