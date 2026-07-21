package main

func main() {
	m := map[string]int{"a": 1, "b": 2}
	println(m)

	var m1 = map[string]int{"a": 2, "b": 3}
	println(m1)

	//定义一个空的map，然后添加值
	var m2 = map[string]int{}
	m2["a"] = 1
	m2["b"] = 2
	println(m2)
}