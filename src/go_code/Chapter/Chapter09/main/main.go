package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
}

func test1() {
	//创建一个新文件，写入内容
	//1.打开文件
	filePath := "e:/test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
	}

	//及时关闭file句柄
	defer file.Close()

	//准备写入语句
	str := "hello world\n你好啊！\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString(str)
	}

	//因为writer是带缓存，因此在调用writerString方法时，其实内容是先写入到缓存的,所以需要调用flush方法将缓存的内容真正写入到文件中
	_ = writer.Flush()
}

func test2() {
	filePath := "e:/test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer file.Close()

	str := "你好啊！My dear friend! \n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, _ = writer.WriteString(str)
	}
	_ = writer.Flush()

}

func test3() {
	filePath := "e:/test.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf(str)
	}

	str := "插入语句 Insert！\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString(str)
	}
	_ = writer.Flush()
}

func test4() {
	//将e:/test.txt 导入到 f:/test.txt

	//1.首先将 d:/test.txt 内容读取到内存
	//2.将读取到的内容写入 e:/kkk.txt

	file1Path := "e:/test.txt"
	file2Path := "f:/test.txt"
	content, err := ioutil.ReadFile(file1Path)

	if err != nil {
		fmt.Printf("read file1 err=%v", err)
	}

	err = ioutil.WriteFile(file2Path, content, 0666)
	if err != nil {
		fmt.Printf("write file2 err=%v", err)
	}

}

func test5() {
	filePath := "e:/test.txt"
	_, err := os.Stat(filePath)
	if err == nil {
		fmt.Println("文件存在")
		return
	}
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return
	}
	fmt.Println("文件不确定是否存在")
}
