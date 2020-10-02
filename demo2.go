package main

import "fmt"

func main()  {
	//直角三角形两条直角边分别是a和b，a=3，b=4.那么斜边c=？
	//var:variable变量	int:表示整数类型
	//格式：var	变量的名字	变量类型
	var a int = 3//定义一个整数变量，名字为a，值是3
	var b int = 4//另外一条直角边
	var c int
	//通过数学公式计算斜边的值
	c =a*a +b*b
	fmt.Println("c的平方为：",c)
	//查看变量c的数据类型	format：格式
	fmt.Printf("变量c的数据类型为:%T",c)
}
