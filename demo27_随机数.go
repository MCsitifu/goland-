package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机数：随机的给定义给定一个数
	//在go语言当中如何生成一个随机数：
	//math（数学）/rand（random随机数）包
	//随机数的产生：需要设置一个种子，每次生成的随机数的顺序都是相同的
	//Intn(n int):生成的随机数的范围，区间范围是:[0,n)
	num1 := rand.Intn(5)//int表示生成的随机数是一个整数，5表示生成的随机数的范围：0，1，2，3，4
	fmt.Println(num1)
	num1 = rand.Intn(5)
	fmt.Println(num1)
	num1 = rand.Intn(5)
	fmt.Println(num1)
	//rand.float32()生成一个浮点数的随机数

	//设置一个种子值、
	//rand.Seed()为随机数生成设置一个种子值
	//时间戳：当前或者某一个时刻的时间，时间戳不会重复,时间戳的格式往往是一个int64的整数
	//在设置生成随机数的种子时，往往会使用时间戳的格式
	//获取当前时间的时间戳
	time1 := time.Now()//time.Now()返回的是time类型
	//time1.Unix()//获取time1时间所对应的秒数//time1时间距离1970年1月1日，格林尼治时间
	//time1.UnixNano()//时间距离1970年1月1日格林尼治标准时间的纳秒数
	rand.Seed(time1.UnixNano())//设置了随机数的种子值
	num2 := rand.Intn(100)
	fmt.Println(num2)

}
