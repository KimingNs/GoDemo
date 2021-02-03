package main

import (
	"fmt"
)

func init() {
	fmt.Println("---------begin---------")
}

type Monkey struct {
	Name string
}

type Flying interface {
	Flying()
}

type Swimming interface {
	Swimming()
}

func (m Monkey) climbing() {
	fmt.Println(m.Name, "climbing...")
}

type LittleMonkey struct {
	Monkey
}

func (lm LittleMonkey) Flying() {
	fmt.Println(lm.Name, "Flying...")
}

func (lm LittleMonkey) Swimming() {
	fmt.Println(lm.Name, "Swimming...")
}

func main() {
	m := Monkey{
		Name: "Test",
	}
	lm := LittleMonkey{
		Monkey{
			Name: "Tom",
		},
	}
	m.climbing()
	lm.climbing()
	//var fly Skill = lm
	//var swim Skill = lm
	lm.Flying()
	lm.Swimming()

}
