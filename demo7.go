package main

import "fmt"
//全局变量
var num = 24//声明了一个全局变量，作用范围大，可以在该程序中的任何地方使用
func main()  {
	var num = 20
	var age = 18
	//变量作用域，就是作用范围，从声明处开始，到该变量所处的{}的结尾处
	fmt.Println(age)
	fmt.Println(num)

}

func hello()  {

	fmt.Println(num)
}