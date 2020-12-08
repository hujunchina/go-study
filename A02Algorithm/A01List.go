package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	fmt.Println(l.Back().Value, l.Len())
	fmt.Println(l.Front().Value, l.Len())
}