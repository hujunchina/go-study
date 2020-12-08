package main

import (
	"fmt"
)

type Cortex struct{
	x int
	y int
}

var mm = map[string]Cortex{
	"H": {1,2},
	"P": {3,4},
}

func map01() {
	m := make(map [string]Cortex)
	m["h"]= Cortex{1,2}
	fmt.Println(m)
	fmt.Println(mm)
}

func map02() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	map01()
	map02()
}
