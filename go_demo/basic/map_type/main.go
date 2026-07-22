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

	var m3 = make(map[string]int)
	m3["a"] = 1
	m3["b"] = 2
	println(m3)

	delete(m3, "a")
	println(m3)

	for k, v := range m3 {
		println(k, v)
	}

	for k := range m3 {
		println(k)
	}

	for _, v := range m3 {
		println(v)
	}

	for k := range m3 {
		delete(m3, k)
	}

	println(m3)

	m4 := map[string]int{"a": 1, "b": 2}

	for k, v := range m4 {
		println(k, v)
	}

	for k := range m4 {
		println(k)
	}

	for _, v := range m4 {
		println(v)
	}

	for k := range m4 {
		delete(m4, k)
	}
	println(m4)

	//保存的值是string
	m5 := map[string]string{"a": "zhangsan", "b": "lisi"}
	for k, v := range m5 {
		println(k, v)
	}
	//判断是否存在a键，ok为true就是存在a的key
	if _, ok := m5["a"]; ok {
		println("存在a键")
	} else {
		println("不存在a键")
	}

	//保存的值是any
	m6 := map[string]any{"a": 1, "b": "lisi"}
	for k, v := range m6 {
		println(k, v)
	}

	//make创建map
	m7 := make(map[string]int)
	m7["a"] = 1
	m7["b"] = 2
	println(m7)

	for k, v := range m7 {
		println(k, v)
	}

	
	//map与切片结合
	var userinfo = make([]map[string]string,3,3);
	userinfo[0] = make(map[string]string)
	userinfo[0]["name"] = "zhangsan"

	userinfo[1] = make(map[string]string)
	userinfo[1]["name"] = "lisi"

	userinfo[2] = make(map[string]string)
	userinfo[2]["name"] = "wangwu"

	for i := 0; i < len(userinfo); i++ {
		println(userinfo[i]["name"])
	}
	
	//map对象的值是切片
	var userinfo1 = make(map[string][]string,3)
	userinfo1["a"] = []string{"zhangsan","lisi"}

}
