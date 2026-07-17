package main

import "fmt"

func main() {
	a := 3.12
	fmt.Println(a)
	fmt.Printf("值:%v---%f,类型%T\n", a, a, a)

	var b float32 = 3.1415926
	fmt.Println(b)
	fmt.Printf("值:%v---%.02f,类型%T\n", b, b, b)

	//科学计数法表示浮点类型
	var c float64 = 3.14e2
	fmt.Println(c)

	//int转float
	d := 10
	var e float64 = float64(d)
	fmt.Println(e)

}
