package main

import(
	"fmt"
	"io"
	"strings"
)

func test01() {
	r := strings.NewReader("hello, reader!")
	//每次 8 字节的速度读取它的输出
	b := make([]byte, 8)

	for{
		n, err := r.Read(b)
		fmt.Printf("%v, %v, %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF{
			break
		}
	}
}

func main() {
	test01()
}
