package main

import "fmt"

//快速排序
//说明
//1.left 表示数组左边的下标
//2.right 表示数组右边的下标
//3.array 表示要排序的数组
func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right
	//pivot 是中轴,支点
	pivot := (left + right) / 2
	temp := 0
	for l < r {
		for array[l] < array[pivot] {
			l++
		}
		for array[r] > array[pivot] {
			r--
		}
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		if array[l] == array[pivot] {
			r--
		}
		if array[r] == array[pivot] {
			l++
		}
	}
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}
}

func QuickSortMiddle(left int, right int, arr *[10]int) {
	//从左边找到比中间值大的放到右区间，
	//从右边找到比中间值小的，放到左区间
	l := left
	r := right
	m := (l + r) / 2
	for l < r {
		fmt.Println("before : ", arr)
		for ; arr[l] < arr[m]; l++ {
		}
		for ; arr[r] > arr[m]; r-- {
		}
		arr[l], arr[r] = arr[r], arr[l]
		fmt.Println("after : ", arr)
	}
}

func QuickSortFirst(left int, right int, arr *[10]int) {
	//从数组首个元素开始
	l := left
	r := right
	key := arr[left]
	for l < r {
		fmt.Println("before : ", arr)
		for {
			if l == r {
				arr[l] = key
				break
			}
			if arr[r] <= key {
				arr[l] = arr[r]
				l++
				break
			}
			r--
		}
		for {
			if l == r {
				arr[l] = key
				break
			}
			if arr[l] >= key {
				arr[r] = arr[l]
				r--
				break
			}
			l++
		}
		if l == r {
			arr[l] = key
		}
		fmt.Println("after : ", arr)
	}
	if left < l {
		QuickSortFirst(left, l-1, arr)
	}
	if right > r {
		QuickSortFirst(r+1, right, arr)
	}
}

func main() {
	arr := [10]int{6, 5, 2, 10, 1, 8, 8, 9, 10, 4}
	QuickSortFirst(0, 9, &arr)
	fmt.Println(arr)
}
