package main

import "fmt"

func main() {
	num1 := 5//一个变量num1 值5
	fmt.Printf("num1的值是：%d\n",num1)
	num2 := num1//值拷贝：将原来的数据拷贝一份，赋值给新的变量num2
	num1 = 8
	fmt.Printf("修改后的num1的值是：%d\n",num1)
	fmt.Printf("num2的值是：%d\n",num2)

	arr1 := [3]int{17,18,19}
	arr2 := arr1
	arr1[0]=3
	fmt.Println(arr1)
	fmt.Println(arr2)
	//结论：数组的拷贝和赋值是一种值拷贝，当发生赋值操作的时候，会拷贝内容给新的变量到内存空间


	//数据类型：
	//基本数据类型：数值类型：int、float布尔类型：bool字符串类型：string
	//复合数据类型：数组：arr  ...

	//按照计算机的存储方式来分：
	//值类型：当数据发生拷贝或发生传递的时候，传递数据的内容，也就是说将数据复制一份，到新的变量和内存
	//值类型的特点：对于原来的数据，没有任何影响

	//引用类型：......


	num1Arr := [3]int{2,5,1}

	num2Arr := [4]int{2,5,1,6}
	fmt.Println(num1Arr)
	fmt.Println(num2Arr)
	fmt.Printf("num1Arr的数据类型是：%T\n",num1Arr)
	fmt.Printf("num2Arr的数据类型是：%T\n",num2Arr)
//关于数据的数据类型：数组的长度是数组数据类型的一部分，两个长度不相等的数组，其数据类型也不同
//数组的类型是：[size]类型
}
