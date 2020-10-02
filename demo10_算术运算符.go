package main

import (
	"fmt"
)

func main() {
	//数学中的运算：加减乘除(+,-,*,/)
	//运算符概念：参与变量运算的符号，称之为运算符
	//运算符：
	//运算符分类：1.算术运算符+、-、*、/、++、--
	//%:取模
	//++自增
	//--自减

	a := 3//整数3
	b := 5//整数5
	sum := a+b
	fmt.Println(sum)
	sub := a - b
	fmt.Println(sub)
	mul := a*b
	fmt.Println(mul)
	div := a/b
	fmt.Println(div)
	mod := a%b
	fmt.Println(mod)

	age :=18
	fmt.Println("铁锤妹妹的年龄：",age)
	age++//自动加1
	fmt.Println("过年后，铁锤妹妹的年龄：",age)

	num :=2020
	fmt.Println("当前年份:",num)
	num--
	fmt.Println("穿越一年后的年份:",num)
}
