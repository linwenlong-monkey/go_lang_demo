package main

import "fmt"

func main() {

	var a int = 10
	var p *int = &a

	fmt.Println(p)

	//new函数分配内存。
	var p1 *int = new(int) // *int
	*p1 = 10
	fmt.Println(p1)
	fmt.Println(*p1)

	//make和new函数的区别
	//new的是指针类型
	var m map[string]int = make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)

	var m1 *map[string]int = new(map[string]int)
	*m1 = make(map[string]int)
	(*m1)["a"] = 1
	(*m1)["b"] = 2
	fmt.Println(m1)
	
	
}
