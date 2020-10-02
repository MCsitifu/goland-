package main

import "fmt"

func main() {
	//位运算符：计算机中的二进制表示只有0和1：int8，int16：位数
	//十进制10：--->转换成二进制
	//&（按位与）
	//|（按位或）
	//^（按位异或）
	var a = 2//二进制： 10			5   101		5	101
	var b = 7//二进制：111			7	111		9	1001
	a = 5
	b = 9
	res1 := a & b//都是1则为1			101=5		0001=1
	res2 := a | b//有一个1则为1			111=7		1101=13
	res3 := a ^ b//不同则为1				010=2		1100=12
	//把两个变量a、b转换为二进制数
	//每一位进行计算
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	num1 := 3//		11
	num2 := 9//		1001
	res4 := num1 << 2//按位向左移动2位-->3*2^2
	res5 := num2 << 2
	res6 := num1 << 4//按位向左移动4位-->3*2^4
	fmt.Println(res4)
	fmt.Println(res5)
	fmt.Println(res6)
	res7 := num2 >>2//按位向右移动2位-->9*2^-2取小数点前的整数
	fmt.Println(res7)//1001区小数点前的数


}
