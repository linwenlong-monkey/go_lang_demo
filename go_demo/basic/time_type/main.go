package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	fmt.Println(timeObj)

	//格式化日期解释 
	//2006 
	//01
	//02
	//03 12小时制，15 24小时制
	//04
	//05
	fmt.Println(timeObj.Format("2006-01-02"))
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))//格式化输出

	ear := timeObj.Year()
	fmt.Println(ear)
	month := timeObj.Month()
	fmt.Println(month)
	day := timeObj.Day()
	fmt.Println(day)
	hour := timeObj.Hour()
	fmt.Println(hour)
	minute := timeObj.Minute()
	fmt.Println(minute)
	second := timeObj.Second()
	fmt.Println(second)

	//当前时间戳
	unix := timeObj.Unix()
	fmt.Println(unix)

	//时间戳转日期对象 
	//前面是毫秒时间戳，后面是纳秒时间戳。
	timeObj2 := time.Unix(unix,0)
	fmt.Println("日期对象 ",timeObj2)

	unixNano := timeObj.UnixNano()
	fmt.Println(unixNano)

	week := timeObj.Weekday()
	fmt.Println("week = ",week)

	fmt.Println(time.Now().Add(time.Hour * 24 * 7))



	//字符串转时间戳
	var str = "2020-01-01 23:13:05"
	timeObj3, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		fmt.Println("字符串转时间戳失败")
	} else {
		fmt.Println(timeObj3)
	}
	//ParseInLocation转换
	timeObj4, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		fmt.Println("字符串转时间戳失败")
	} else {
		fmt.Println(timeObj4)
	}
	

	//定时器
	time.AfterFunc(time.Second * 5, func() {
		fmt.Println("定时器执行了")
	})

	time.Sleep(time.Second * 10)
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("定时器执行了")
	}


	time.Tick(time.Second * 5)
	for {
		<-time.Tick(time.Second * 5)
		fmt.Println("定时器执行了")
	}

	time.After(time.Second * 5)
	for {
		<-time.After(time.Second * 5)
		fmt.Println("定时器执行了")
	}

	ticker :=  time.NewTicker(time.Second)
	for t := range ticker.C {
		fmt.Println(t)
	}

	
}
