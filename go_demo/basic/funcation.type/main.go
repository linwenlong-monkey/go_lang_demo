package main

import (
	"fmt"
)

//定义函数
// func <funcName>(<params>) <returnType> {

// }

func add(a int, b int) int {
	sum := a + b
	return sum
}

func add2(a, b int) int {
	sum := a + b
	return sum
}

func add3(a, b int) (int, int) {
	sum := a + b
	return sum, sum
}

//这行代码啥意思 
func add4(a, b int) (sum int, sum2 int) {
	sum = a + b
	sum2 = a + b
	return
}

func add5(a, b int) (sum int) {
	sum = a + b
	return
}

func add6(a, b int) (sum int) {
	sum = a + b
	return sum
}

//求两个数的差
func subFn(x, y int) (sub int) {
	if x > y {
		sub = x - y
	} else {
		sub = y - x
	}
	return
}

//函数的可变参数
func add7(args ...int) (sum int) {
	for _, v := range args {
		sum += v
	}
	return
}

//可变参数跟固定参数混合使用，可变参数放最后
func add8(a int, args ...int) (sum int) {
	for _, v := range args {
		sum += v
	}
	sum += a
	return
}




//定义函数类型
type calc func(int, int) int//定义了一个函数类型

func addCalc(a, b int) int {
	return a + b
}

func subCalc(a, b int) int {
	return a - b
}


//函数作为返回值
func calcFn(op string) calc {
	if op == "+" {
		return addCalc
	} else if op == "-" {
		return subCalc
	}
	return nil
}

//闭包,函数里面嵌套一个函数，然后返回函数
func adder(a, b int) func() int {
	return func() int {
		return a - b
	}
}

func main() {
	sum := add(1, 2)
	fmt.Println(sum)


	sub := subFn(1, 2)
	fmt.Println(sub);

	var sumdb int = add7(1, 2, 3, 4, 5)
	fmt.Println(sumdb)

	//函数类型
	var calc1 calc
	calc1 = addCalc
	reuslt := calc1(1, 2)
	fmt.Println(reuslt)
	calc2 := subCalc
	reuslt = calc2(1, 2)
	fmt.Println(reuslt)


	//匿名函数
	calc3 := func(a, b int) int {
		return a + b
	}
	reuslt = calc3(1, 2)
	fmt.Println(reuslt)

	//闭包


	
}