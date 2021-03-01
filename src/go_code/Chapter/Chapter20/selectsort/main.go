package main

import "fmt"

func main() {
	//定义一个数组
	arr := [20]int{9, 3, 0, 1, 6, 4, 5, 10, 19, 13, 15, 16, 17, 18, 2, 8, 7, 11, 12, 14}
	index := 0

	for {
		next := index
		min := next
		if index == len(arr)-1 {
			break
		}
		for i := index; i < len(arr); i++ {
			if arr[next] < arr[min] {
				min = next
			}
			next++
		}
		fmt.Println("arr[min]=", arr[min])
		arr[index], arr[min] = arr[min], arr[index]
		fmt.Println(arr)
		index++
	}
	fmt.Println(arr)
}
