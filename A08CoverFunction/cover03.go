package main

import "fmt"

func f4(x, y int) (func(int) int, func(int) int) {
	return func(int) int {
		return x+1
	}, func(int) int {
			return y-1
		}
}

func main(){
	t1, t2 := f4(1,1)
	fmt.Println(t1(10), t2(10))
}
