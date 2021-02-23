package main

import "fmt"

//定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode //表示指向下一个结点
}

/*
	给链表插入一个结点
	编写第一种插入方式,在单链表的最后加入.
*/
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1.先找到该链表的最后这个结点
	//2.创建一个辅助结点，【帮忙】
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next //让temp不断的指向下一个结点
	}
	//3.现在一定是最后一个结点，将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}

/*
	给链表插入一个结点
	编写第二种插入方式，根据no的编号从小到大插入
*/
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1.找到适当的结点
	//2.创建一个辅助结点
	temp := head
	flag := true
	//让插入的结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil {
			//说明到链表的最后
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode 就应该插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			//说明我们的链表中已经有这个no，就不让插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

//删除一个结点
func DelHeroNode(head *HeroNode, no int) {
	temp := head
	flag := true
	//找到要删除结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil {
			break
		}
		if temp.next.no == no {
			flag = false
			break
		}
		temp = temp.next
	}
	if flag {
		fmt.Println("该链表没有这个结点")
	} else {
		temp.next = temp.next.next
	}
}

//显示链表的所有结点信息
func ListHeroNode(head *HeroNode) {
	//1.创建一个辅助结点
	temp := head

	//先判断该链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("empty link")
		return
	}

	//2.遍历这个链表
	for {
		fmt.Printf("[%d,%s,%s] -->", temp.next.no, temp.next.name, temp.next.nickname)
		//判断是否到链表尾
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	//1.先创建一个头结点
	head := &HeroNode{}

	//2.创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "songjiang",
		nickname: "jishiuyu",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "lujunyi",
		nickname: "yuqilin",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "wusong",
		nickname: "dahu",
	}

	//3.加入
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)

	DelHeroNode(head, 4)
	//4.显示
	ListHeroNode(head)
}
