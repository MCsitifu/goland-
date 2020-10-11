package main

import "fmt"

func main() {
	////1、for循环实现的数组的遍历

	////var name int
	//ageArr := [3]int{18,20,22}
	////fmt.Println(ageArr[0])//少数可以打印
	////fmt.Println(ageArr[1])
	////fmt.Println(ageArr[2])
	//for index := 0;index<=len(ageArr)-1 ;index++  {
	//	fmt.Printf("第%d个元素的值是：%d\n",index+1,ageArr[index])
	//	//name += ageArr[index]//数组之和
	//}
	////fmt.Println(name)

	//2、range：范围、区域
	//go语言当中有一个新语法：for....range xxx通过这个语法可以遍历容器
	arr11 := []int{5,1,4,9,20,14,31}
	//a,b := 1,5
	for index,value := range arr11{//index,value 这两个变量的名字，能不能叫其他的？当然可以
		fmt.Println(index,value)

	}
	var num int
	for _,b := range arr11{//变量舍弃：_下划线
		fmt.Println(b)
		num += b
	}
	fmt.Println(num)

}
