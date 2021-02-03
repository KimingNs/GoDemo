package main

import (
	"fmt"
)

func init() {
	fmt.Println("---------begin---------")
}

type AInterface interface {
	test01()
	BInterface
}

type BInterface interface {
	test02()
}

type CInterface interface {
	AInterface
	BInterface
	test03()
}

type C struct {
}

func (c C) test01() {
	fmt.Println("test01")
}

func (c C) test02() {
	fmt.Println("test02")
}

func (c C) test03() {
	fmt.Println("test03")
}

func main() {
	stu := C{}
	var a AInterface = stu
	var b BInterface = stu
	var c CInterface = stu
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)
	a.test01()
	a.test02()
	b.test02()
	c.test01()
	c.test02()
	c.test03()

}
