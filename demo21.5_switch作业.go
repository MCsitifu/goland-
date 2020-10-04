package main

import (
	"fmt"
)

func main() {
	var num  int
	var num1  int
	fmt.Println("请输入你想要计算的一个值或者两个值")
	fmt.Scanln(&num,&num1)
	var str string
	fmt.Println("请输入你想要用于计算的符号(++、--、%)：")
	fmt.Scanln(&str)
	//if str == "++" && str== "--" {
		switch str {
			case "++" :
			fmt.Println("++后的值为：", num+1,num1+1)
			case "--" :
			fmt.Println("--后的值为：", num-1,num1+1)
			case "%":
			fmt.Println("两数取得的余数为：",num % num1)
		default:
			fmt.Println("请输入正确的符号")
			//}
			//}else if str=="%" {
			//fmt.Println("两数取得的余数为：",num % num1)
			}

}