package main

import "fmt"

func main() {
	//关系运算符：>,<.>=.<=,==,!=,运算符运算的结果是ture/flase,只能是这两个结果值当中的一个，go语言中的一种
	//
	age1 := 18
	age2 := 19

	result := age1 > age2
	fmt.Println(result)
	fmt.Printf("result的类型：%T\n",result)

	result1 := age1 <= age2
	fmt.Println(result1)

	num1 := 100
	num2 := 100
	result2 := num1 >= num2
	fmt.Println(result2)
	result3 :=num1 <= num2
	fmt.Println(result3)
	result4 := num1 == num2//==判断两个值是否相等
	fmt.Println(result4)
	result5 := num1 != num2
	fmt.Println(result5)


}