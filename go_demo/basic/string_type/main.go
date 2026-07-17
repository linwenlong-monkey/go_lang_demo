package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "Hello"
	var str2 = "World"
	str3 := "Hello Golang"

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)

	//字符串默认是空
	var str4 string
	fmt.Printf("值: %v, 类型: %T", str4, str4)


	//多行字符串
	var str5 = `
	   	<html>
	   		<body>
	   			<h1>Hello World</h1>
	   		</body>
	   	</html>
	`
	fmt.Println(str5, len(str5))

	//拼接字符串
	var str6 = str1 + str2
	fmt.Println(str6)
	//拼接字符串第二种方式
	var str7 = fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str7)

	//分割字符串
	var str8 = "hello,world,golang,imooc"
	var slice1 = strings.Split(str8, ",")
	fmt.Println(slice1)
}
