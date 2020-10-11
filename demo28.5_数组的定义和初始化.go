package main

import "fmt"

func main() {
	//定义一个数组，用于存放数值
	var ageArr = [3]int{18,19,20}
	fmt.Println(ageArr)
	ageArr[1]=20
	fmt.Println(ageArr)//通过下标来修改数值

	//ageArr1 := [4]int{}//申请了长度为4个内存的一个数组空间，默认值都是0
	ageArr1 := [4]int{18,19,20,20}
	fmt.Println(ageArr1)

	numArr := [5]int{19,1,2,3,10}//{}内的元素个数不能大于[]中的数字
	fmt.Println(numArr)

	//不设置数组的大小，此时Go语言会更具元素的个数来设置数组的大小
	num1Arr := []int{80,90,70,100,85}
	fmt.Println(num1Arr)
	//查看数据的长度
	fmt.Println(len(num1Arr))
	//查看数据类型
	fmt.Printf("num1Arr的数据类型是：%T\n",num1Arr)

	//省略写法
	AgeArr := [...]int{15,18,20,15}
	fmt.Printf("AgeArr的数据类型是：%T\n",AgeArr)
}
