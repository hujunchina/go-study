package main

import "fmt"

func fx(x int) func(int){
	return func(int){
		fmt.Println(x)
	}
}

func main(){
	t := fx(10)
	t(1)
}