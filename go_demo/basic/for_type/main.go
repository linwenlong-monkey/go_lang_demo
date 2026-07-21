package main

import (
	"fmt"
	
)

func main() {

	var str string = "Hello"
	
	for key, value := range str {
		fmt.Println(key, value)
	}


	var arr = []string{"Hello", "World", "Golang"}
	for key, value := range arr {
		fmt.Println(key, value)
	}

	var m = map[string]string{"name": "张三", "age": "18", "sex": "男"}
	for key, value := range m {
		fmt.Println(key, value)
	}


	
}
