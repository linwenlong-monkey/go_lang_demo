package main

import "fmt"
func main() {

	//初始化一个5个int的数组
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	fmt.Println(a)

	var str string = "Hello"
	fmt.Println(str)

	
	var arr = []string{"Hello", "World", "Golang"}
	fmt.Println(arr)

	//三个点初始化数组
	//数组长度会自动计算
	var arr2 = [...]string{"Hello", "World", "Golang"}
	fmt.Println(arr2)

	//二维数组

	var arr3 = [3][2]string{
		{"Hello", "World"},
		{"Golang", "Programming"},
		{"Go", "Language"},
	}
	fmt.Println(arr3)
	//便利二维数组
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Println(arr3[i][j])
		}
	}
	//for range便利二维数组
	for i, v := range arr3 {
		for j, v2 := range v {
			fmt.Println(i, j, v2)
		}
	}

	//声明一个切片，没有指定长度
	var slice []string
	slice = append(slice, "Hello")
	slice = append(slice, "World")
	fmt.Println(slice)

	//声明一个切片，指定长度
	var slice2 = make([]string, 5)
	slice2[0] = "Hello"
	slice2[1] = "World"
	fmt.Println(slice2)

	//声明一个切片，指定长度和容量
	var slice3 = make([]string, 5, 10)
	slice3[0] = "Hello"
	slice3[1] = "World"
	fmt.Println(slice3)

	//切片的切片
	var slice4 = slice3[0:2]
	fmt.Println(slice4)

	//切片的切片，省略下标，默认从0开始
	var slice5 = slice3[:2]
	fmt.Println(slice5)

	//切片的切片，省略上标，默认到最后
	var slice6 = slice3[2:]
	fmt.Println(slice6)

	//切片的切片，省略上下标，默认从0到最后
	var slice7 = slice3[:]
	fmt.Println(slice7)

	//切片的切片，省略上下标，默认从0到最后
	var slice8 = slice3[1:4]
	fmt.Println(slice8)

	//切片的切片，省略上下标，默认从0到最后
	var slice9 = slice3[1:4:5]
	fmt.Println(slice9)

	//声明切片，并给定特定的个数值
	var slice10 = []string{1: "Hello", 2: "World"}
	println(slice10)

	//声明切片，并给定特定的个数值
	var slice11 = []string{"Hello", "World"}
	println(slice11)

	//切片的默认值是nil
	var slice12 []string
	fmt.Println("slice12 = ", slice12)

	//将一个数组的值用一个切片接收
	var arr4 = [5]string{"Hello", "World", "Golang", "Programming", "Go"}
	var slice13 = arr4[0:3]//从第0个开始，获取数组里面3个值
	fmt.Println(slice13)


	//根据一个切片定义切片
	var slice14 = []string{"Hello", "World", "Golang", "Programming", "Go"}
	var slice15 = slice14[0:3]
	fmt.Println(slice15)

	//根据一个切片定义切片
	var slice16 = []string{"Hello", "World", "Golang", "Programming", "Go"}
	var slice17 = slice16[:3]
	fmt.Println(slice17)

	//根据一个切片定义切片
	var slice18 = []string{"Hello", "World", "Golang", "Programming", "Go"}
	var slice19 = slice18[2:]
	fmt.Println(slice19)

	//切片的长度和容量
	var slice20 = []string{"Hello", "World", "Golang", "Programming", "Go"}
	fmt.Println("len = ",len(slice20), "cap=",cap(slice20))
}