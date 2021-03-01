package main

import "fmt"

func main() {
	arr := [10]int{6, 3, 2, 5, 10, 1, 4, 8, 7, 9}
	arrNew := [10]int{}

	index := 0
	arrNew[index] = arr[index]
	for {
		if index >= len(arr) {
			break
		}
		fmt.Println(index)
		next := index - 1
		for ; next >= 0; next-- {
			if arr[index] < arrNew[next] {
				arrNew[next], arrNew[next+1] = arr[index], arrNew[next]
			} else {
				if arrNew[next] < arrNew[next+1] {
					break
				}
				arrNew[index] = arr[index]
				break
			}
			fmt.Println(arrNew)
		}
		index++
	}
	fmt.Println(arrNew)
}
