package main

import "fmt"

type Person struct{
	name string
	age int
}

type myInterface interface {
	getName() string
	getAge() int
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (p Person) getName() string{
	return p.name
}

func (p Person) getAge() int {
	return p.age
}

func inter01() {
	person := Person{"Hujun", 1}
	fmt.Println(person.getAge())
}

//指定了零个方法的接口值被称为[空接口]
func interface02(){
	var i interface{}
	fmt.Printf("%v, %T\n", i, i)

	i = "hello"
	fmt.Printf("%v, %T\n", i, i)

	//类型断言
	s, ok := i.(string)
	fmt.Println(s, ok)
}


func main() {
	inter01()
	interface02()
}
