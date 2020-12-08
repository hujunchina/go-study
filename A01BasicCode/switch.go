package main

import (
	"fmt"
	"runtime"
	"time"
)

func switch01()  {
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main()  {
	fmt.Println("go runs on")
	switch  os:= runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x")
	case "liux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	switch01()
}


