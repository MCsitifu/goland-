package main

import "fmt"

func main()  {
	//格式一：const 常量名字 数据类型 = 内容值
	//格式二：const 常量名字 = 数值
	const PAI float32 = 3.14//定义一个常量
	fmt.Println(PAI)
	//PAI = 3.1415926  在定义的时候const表明，该变量是一个常量，常量是不可更改的
	const BAIDU = "http://www.baidu.com"
	var  address = "nanchang"
	fmt.Println(BAIDU)//全部大写的通常是常量
	fmt.Println(address)//小写通常是变量

	//常量的集合
	const (
		num = 10
		name = "jinsheng"
		sex = "male"
	)
	//常量举例
	const (
		Monday = "星期一"
		Tuesday = "星期二"
		Wednesday = "星期三"
		Thursday = "星期四"
	)
	//常量组
	const (
		x int = 10
		//y int = 10
		//z int = 10
		y
		z
		name1 string = "jinsheng"
		sex1 string = "male"

	)
	fmt.Println(x,y,z)
}
