package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//2.输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3.转成稀疏数组
	//思路
	//(1).遍历 chessMap，如果我们发现有一个元素的值不为0，创建一个node结构体
	//(2).将其放入到对应的切片即可

	var sparseArr []ValNode

	//标准的一个稀疏数组应该还有一个 记录元素的二维数组的规模（行和列，默认值）
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode值结点
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d:%d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
	//将稀疏数组恢复为原始数组
	var chessMap2 [11][11]int

	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	//看看chessMap2 是不是恢复
	fmt.Println("恢复后的原始数据......")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
