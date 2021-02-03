package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}

func copyFile(fileName string, srcName string) (written int64, error error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//通过srcFile，获取到Reader
	reader := bufio.NewReader(srcFile)
	//打开fileName
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//通过file，获取到writer
	writer := bufio.NewWriter(file)
	defer file.Close()

	return io.Copy(writer, reader)
}

func test1() {
	//copy 图片到另外一个目录

	//调用copyFile()
	srcFile := "e:/1.jpg"
	dstFIle := "f:/2.jpg"
	_, err := copyFile(dstFIle, srcFile)
	if err == nil {
		fmt.Println("copy success")
	} else {
		fmt.Println("copy file err=", err)
	}
}

//定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int //记录英文个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其他字符个数
}

func test2() {
	//思路：打开一个文件，创建一个Reader
	//每读取一行，就去统计改行有多少个 英文，数字，空格和其他字符
	//然后将结果保存到一个结构体
	fileName := "E:\\Go\\src\\go_code\\Chapter09\\main\\message.go"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()

	//定义个CharCount实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'A' && v <= 'z':
				count.ChCount++
			case v == ' ':
				count.SpaceCount++
			case v >= 0 && v <= 9:
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	//开始循环读取fileName的内容
	//输出统计结果
	fmt.Println("ChCount=", count.ChCount, ",NumCount=", count.NumCount, ",SpaceCount=", count.SpaceCount, ",OtherCount=", count.OtherCount)
}

func test3() {
	fmt.Println("命令行的参数有", len(os.Args))
	//遍历os.Args切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}

func test4() {
	//定义几个变量，用户接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	//"u":就是 -u 指定参数
	//"":默认值
	//"用户名，默认为空":说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "p", "", "密码，默认为空")
	flag.StringVar(&host, "h", "", "主机，默认为空")
	flag.IntVar(&port, "port", 3306, "端口号，默认为空")

	//这里有一个非常重要的操作，转换， 必须调用该方法
	flag.Parse()
	//输出结果
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v", user, pwd, host, port)
}
