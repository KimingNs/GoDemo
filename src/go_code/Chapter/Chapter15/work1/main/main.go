package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//numChan := make(chan int, 10)
	numNewChan := make(chan int, 1000)
	writeChan := make(chan bool, 10)
	readChan := make(chan bool, 10)
	go writeDataToFile1(writeChan)
	go writeDataToFile2(writeChan)
	go writeDataToFile3(writeChan)
	go writeDataToFile4(writeChan)
	go writeDataToFile5(writeChan)
	go writeDataToFile6(writeChan)
	go writeDataToFile7(writeChan)
	go writeDataToFile8(writeChan)
	go writeDataToFile9(writeChan)
	go writeDataToFile10(writeChan)
	go readFileToData1(writeChan, readChan, numNewChan)
	go readFileToData2(writeChan, readChan, numNewChan)
	go readFileToData3(writeChan, readChan, numNewChan)
	go readFileToData4(writeChan, readChan, numNewChan)
	go readFileToData5(writeChan, readChan, numNewChan)
	go readFileToData6(writeChan, readChan, numNewChan)
	go readFileToData7(writeChan, readChan, numNewChan)
	go readFileToData8(writeChan, readChan, numNewChan)
	go readFileToData9(writeChan, readChan, numNewChan)
	go readFileToData10(writeChan, readChan, numNewChan)
	for {
		if len(readChan) == 10 {
			break
		}
	}

	fmt.Println("Channel运行完毕")
}

func writeDataToFile1(boolChan chan bool) {
	f, err := os.Create("e:/testChannel1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}
func writeDataToFile2(boolChan chan bool) {
	f, err := os.Create("e:/testChannel2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}
func writeDataToFile3(boolChan chan bool) {
	f, err := os.Create("e:/testChannel3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}
func writeDataToFile4(boolChan chan bool) {
	f, err := os.Create("e:/testChannel4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}
func writeDataToFile5(boolChan chan bool) {
	f, err := os.Create("e:/testChannel5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}
func writeDataToFile6(boolChan chan bool) {
	f, err := os.Create("e:/testChannel6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}
func writeDataToFile7(boolChan chan bool) {
	f, err := os.Create("e:/testChannel7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}
func writeDataToFile8(boolChan chan bool) {
	f, err := os.Create("e:/testChannel8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}
func writeDataToFile9(boolChan chan bool) {
	f, err := os.Create("e:/testChannel9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}
func writeDataToFile10(boolChan chan bool) {
	f, err := os.Create("e:/testChannel10.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(rand.Intn(200000)) + " ")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	boolChan <- true

}

func readFileToData1(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel1.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew1.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
func readFileToData2(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel2.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew2.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
func readFileToData3(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel3.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew3.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
func readFileToData4(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel4.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew4.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
func readFileToData5(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel5.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew5.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
func readFileToData6(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel6.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew6.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
func readFileToData7(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel7.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew7.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
func readFileToData8(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel8.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew8.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
func readFileToData9(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel9.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew9.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
func readFileToData10(writeChan chan bool, readChan chan bool, numNewChan chan int) {
	for {
		if len(writeChan) == 10 {
			break
		}
	}
	filePath := "e:/testChannel10.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		str = strings.Replace(str, " ", "", -1)
		int, _ := strconv.Atoi(str)
		numNewChan <- int
	}

	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = <-numNewChan
	}

	for i := 0; i < len(slice)-1; i++ {
		//遍历i位以后的所有元素，如果比i位元素小，就和i位元素互换位置
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	//接下来将排序后的slice写入新的文件
	f, err := os.Create("e:/testChannelNew10.txt")
	for i := 0; i < len(slice); i++ {
		f.WriteString(strconv.Itoa(slice[i]) + " ")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	readChan <- true
}
