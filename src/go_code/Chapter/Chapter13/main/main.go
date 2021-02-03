package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	m := Monster{Name: "Test", Age: 100, Skill: "Fix"}
	m.Store()
}

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("Json序列化失败", err)
		return false
	}
	filePath := "e:/test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(string(data) + "\n")
	_ = writer.Flush()
	return true
}

func (this *Monster) Restore() bool {
	filePath := "e:/test.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Readfile err = ", err)
		return false
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("Unmarshal err = ", err)
		return false
	}
	return true
}
