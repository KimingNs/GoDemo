package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func init() {
	fmt.Println("---------begin---------")
}

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var intSlice = []int{0, -1, 19, 5, 100, 20}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	for _, v := range heros {
		fmt.Println(v)
	}

	sort.Sort(heros)
	fmt.Println()
	for _, v := range heros {
		fmt.Println(v)
	}
}
