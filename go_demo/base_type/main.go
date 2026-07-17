package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var a = 10
	var b = 20

	num := 30

	fmt.Printf("%d + %d = %d", a, b, a+b)

	fmt.Printf("%d", num)

	num1 := 40

	fmt.Println(num1)

	num2 := 90
	fmt.Println(num2)
	//var 和 :=，二选一，别混用。
	//报错的原因是 var 和 := 不能同时使用，这是两种互斥的变量声明方式：
	var num3 int8 = 88

	fmt.Println(num3)

}

func add(a int, b int) int {
	return a + b
}
