package main

import "fmt"

//自定义类型
type myInt int
type myFloat float64
type myString string
type myBool bool
type myArray [10]int

type Person struct {
	name string
	age int
	sex string
	// hobby []string
	// address string
	// map1 map[string]any
}

//结构体的嵌套
type Student struct {
	person Person
	id int
}

//结构体的继承
type Teacher struct {
	Person
	id int
}
type Employee struct {
	Person
	id int
}
type Worker struct {
	Person
	id int
}


//结构体定义方法
func (p Person) desc() string {
	return fmt.Sprintf("Person{name:%s, age:%d, sex:%s}", p.name, p.age, p.sex)
}

//自定义类型定义方法。给myInt类型自定义方法
func (i myInt) desc() string {
	return fmt.Sprintf("myInt{%d}", i)
}

//自带的类型不能自定义方法
// func (i int) desc1() string {
// 	return fmt.Sprintf("myInt{%d}", i)
// }




func main() {
	person := Person{"张三", 18, "男"}
	fmt.Println(person)

	//new方法实例
	person1 := new(Person)
	person1.name = "张三"
	person1.age = 18
	person1.sex = "男"
	fmt.Println(person1)

	//&实例对象
	person2 := &Person{"张三", 18, "男"}
	fmt.Println(person2)

	//键值对实例化
	person3 := Person{name: "张三", age: 18, sex: "男"}
	fmt.Println(person3)

	fmt.Println(person.desc())

	student := Student{person: Person{name: "张三", age: 18, sex: "男"}, id: 1}
	fmt.Println(student)
}
