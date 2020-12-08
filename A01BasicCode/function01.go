package main

import(
	"fmt"
	"math"
)

func func01(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))
	fmt.Println(func01(hypot))
	fmt.Println(func01(math.Pow))
}
