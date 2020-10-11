package main

import "fmt"

func main() {

	age1 := 18
	age2 := 20
	fmt.Println(age1 > age2)//比较大小
	fmt.Println(age1 == age2)//判断是否相等

	num1 := [3]int{18,19,20}
	num2 := [3]int{18,19,20}
	fmt.Println(num1 == num2)//说明：两个数组是可以判断是否相等，会挨个元素进行判断，判断数组相等的时候，两个数组长度要一致

}
