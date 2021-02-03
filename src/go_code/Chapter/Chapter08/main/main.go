package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	test1()
}

func test1() {
	file, err := os.Open("E:/Go/src/go_code/Chapter01/main/message.go")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//文件是指针
	defer file.Close()
	//有缓冲，用于文件大的情况
	// 创建一个 *Reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
}

func test2() {
	//无缓冲，用于文件不大的情况
	content, err := ioutil.ReadFile("E:/Go/src/go_code/Chapter01/main/message.go")
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content))
}

func test() {
	filePath := "e:/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file=%v", file)

	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}

func test3() {
	filePath := "e:/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	fmt.Println("文件读取结束")
}

func test4() {
	filePath := "e:/test.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print("read file err=%v", err)
	}
	fmt.Printf("%v", string(content))
}
