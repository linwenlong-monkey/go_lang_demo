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

	//切片转字符串

	// // ❌ 错误写法
	// // s := []string("a", "b", "c")

	// // ✅ 正确写法（切片字面量）
	// s := []string{"a", "b", "c"}

	aa := []string {"hello golang","php","java"}
	var str9 = strings.Join(aa, ",")
	fmt.Println(str9)

	//字符串包含
	var str10 = "hello golang"
	var b1 = strings.Contains(str10, "golang")
	fmt.Println(b1)

	//字符串相等
	var b2 = strings.EqualFold("hello golang", "hello golang")
	fmt.Println(b2)

	//前缀
	var b3 = strings.HasPrefix("hello golang", "hello")
	fmt.Println(b3)

	//后缀
	var b4 = strings.HasSuffix("hello golang", "golang")
	fmt.Println(b4)

	//索引
	var index = strings.Index("hello golang", "golang")
	fmt.Println(index)

	//最后索引
	var lastIndex = strings.LastIndex("hello golang", "golang")
	fmt.Println(lastIndex)

	//替换
	var str11 = strings.Replace("hello golang", "golang", "world", 1)
	fmt.Println(str11)

	//大小写转换
	var str12 = strings.ToUpper("hello golang")
	fmt.Println(str12)

	var str13 = strings.ToLower("HELLO GOLANG")
	fmt.Println(str13)

	//去除前后空格
	var str14 = strings.TrimSpace("  hello golang  ")
	fmt.Println(str14)

	//去除两边指定字符
	var str15 = strings.Trim("hello golang", "h")
	fmt.Println(str15)

	//去除左边指定字符
	var str16 = strings.TrimLeft("hello golang", "h")
	fmt.Println(str16)

	//去除右边指定字符
	var str17 = strings.TrimRight("hello golang", "g")
	fmt.Println(str17)

}
