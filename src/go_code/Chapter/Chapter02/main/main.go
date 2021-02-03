package main

import (
	"fmt"
)

func init() {
	fmt.Println("---------begin---------")
}

type AInterface interface {
	Say()
}

type BInterface interface {
	Work()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type Monster struct {
}

func (m Monster) Work() {
	fmt.Println("monster working...")
}

func (m Monster) Say() {
	fmt.Println("monster saying...")
}

func main() {
	m := Monster{}
	var a AInterface = m
	a.Say()
	var b BInterface = m
	b.Work()
}
