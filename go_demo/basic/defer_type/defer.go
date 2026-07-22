package main

//匿名返回值和命名返回值配合defer修复修改

func add(a, b int) (result int) {
	defer func() {
		result += 10
	}()
	return a + b
}

func add1(a, b int) (int) {
	var result int
	defer func() {
		result += 10
	}()
	return a + b
}





func main() { 
	//defer函数 延迟执行代码，然后逆序执行defer修饰的语句
	defer println("defer")
	println("main")

	defer func() {
		println("defer2")
	}()
	//先打印main，然后defer2，最后defer

	// //panic函数
	// panic("panic")

	//recover函数处理异常
	defer func() {
		if err := recover(); err != nil {
			println("recover:", err)
		}
	}()
}
