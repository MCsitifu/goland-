package main

import "fmt"

func main()  {
	//变量的默认值
	var a int
	fmt.Println(a)

	var b string
	fmt.Println(a)
	fmt.Println(b)
	//int：整数类型		string：字符串，默认值""		float：浮点数，带有小数点的数据，1.2

	var grade float32
	fmt.Println(grade)//float默认值也是0，go语言中有float32位和float64位两种

	var st1,st2 = "jin","sheng"
	fmt.Println(st1,st2)

	//变量的集合
	var (
		name = "jinsheng"
		sex = "male"
		address = "nanchang"
		age = 21
	)
	fmt.Println(name,sex,address,age)

	//舍弃变量
	num1,_ :=1,2//_下划线表示将该位置上的变量进行舍弃
	fmt.Println(num1)
}
