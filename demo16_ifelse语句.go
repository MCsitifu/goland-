package main

import "fmt"

func main() {
	//if else语句
	//如果男士年龄大于22周岁,则满足国家法定结婚年龄，可以结婚，否则不能结婚
	age := 19
	if age>=22 {
		fmt.Println("该男士满足国家法定结婚年龄")

	}else {
		fmt.Println("该男子还没有达到国家法定结婚年龄")
	}

	//给定一个变量num，求输出该变量的奇偶性
	num := 99
	if num % 2 == 1 {
	fmt.Println("该数为奇数")
	}else {
		fmt.Println("该数为偶数")
	}
}
