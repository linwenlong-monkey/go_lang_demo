package main

import "fmt"

func main() {
	
	a := 'a'
	println(a)

	fmt.Printf("a = %v, 类型 = %T, 原样输出字符 = %c",a,a,a);


	//获取字符串的某个字符
	b := "hello world"
	fmt.Printf("b[0] = %v, 类型 = %T, 原样输出字符 = %c",b[0],b[0],b[0]);

	//一个汉字占3个字节，一个字母占一个字节 unsafe.Sizeof() 获取变量占用的字节数
	c := "中国"
	fmt.Printf("c[0] = %v, 类型 = %T, 原样输出字符 = %c",c[0],c[0],c[0]);

	//rune类型是int32的别名，表示一个Unicode码点
	var d rune = '中'
	fmt.Printf("d = %v, 类型 = %T, 原样输出字符 = %c",d,d,d);

	//rune类型是int32的别名，表示一个Unicode码点
	var e rune = 'a'
	fmt.Printf("e = %v, 类型 = %T, 原样输出字符 = %c",e,e,e);

	//for循环获取字符串里面的字符
	for i := 0; i < len(b); i++ {
		fmt.Printf("b[%d] = %v, 类型 = %T, 原样输出字符 = %c\n",i,b[i],b[i],b[i]);
	}
	//i、v不需要的时候可以下划线
	for i,v := range b {
		fmt.Printf("b[%d] = %v, 类型 = %T, 原样输出字符 = %c\n",i,v,v,v);
	}

	//初始化语句写在外面
	j := 0
	for ; j < len(c); j++ {
		fmt.Printf("c[%d] = %v, 类型 = %T, 原样输出字符 = %c\n",j,c[j],c[j],c[j]);
	}


}
