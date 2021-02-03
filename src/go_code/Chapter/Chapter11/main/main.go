package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"monster_name"`
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string  `json:"monster_skill"`
}

func main() {
	testStruct()
}

func testMap() {
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 100
	a["address"] = "洪崖洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败 err=%v\n", err)
	}
	fmt.Printf("monster序列化后=%v\n", string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "Jack"
	m1["age"] = 7
	m1["address"] = "beijing"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "Tom"
	m2["age"] = 200
	m2["address"] = "MoxiMoxi"
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

func testStruct() {
	monster := Monster{
		Name:     "niuniu",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      5000.0,
		Skill:    "quanquan",
	}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("Json序列化失败", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}

//将json字符串转换为结构体
func unmarshalStruct() {
	str := "{\"monster_name\":\"CowMoMo\"" +
		",\"monster_age\":500" +
		",\"monster_birthday\":\"2020-1-5\"" +
		",\"monster_sal\":4000" +
		",\"monster_skill\":\"HaSaKi\"}"

	//定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Println("反序列化后 monster=", monster)
}
