package main

import "fmt"

//强调：go语言程序入口是main函数，go语言规定，main函数所在的软件包包名必须是main
func main()  {
	//变量的声明和定义
//var age int
//age=18
//fmt.Println(age)
	var a,b,c int//3,4,5
	a = 3
	b = 4
	c = 5
	fmt.Println(a,b,c)
	//对声明的变量进行修改
	a = 6
	b = 8
	c = 10
	fmt.Println(a,b,c)
	//a = "金圣"//红色波浪线表示语法错误
	//强类型语言：要保证变量的类型和赋值的类型一致，比如：go语言，java语言
	var s1,s2,s3 string = "jin","sheng","tan"
	fmt.Println(s1,s2,s3)

	var name , age ,address  = "金圣叹",18,"南昌"//不同的数据类型的多变量声明方式
	//格式化打印：%s是代表获取字符串的内容
	fmt.Println(name,age,address)
	fmt.Printf("name的数据类型:%T\n",name)
	fmt.Printf("age的数据类型:%T\n",age)
	fmt.Printf("address的数据类型:%T\n",address)

	fmt.Printf("name的内容是:%s,其数据类型是:%T\n",name,name)
	fmt.Printf("age的值是:%d,其数据类型是:%T\n",age,age)

	num1 , num2 := 1 , 3
	fmt.Println(num1+num2)

	num1 , num2 = 5 , 8
	fmt.Println(num1,num2)

	num1 , num3 := 10 , 20
	fmt.Println(num1 , num3)
}
