package main

import "fmt"

func main() {
	//数组的特点：在计算机的内存当中，空间长度是固定的
	//优点：可以直接根据数组的下标进行元素的操作和访问num[0]..
	//缺点：在进行定义的时候，就必须要指定一个长度，而且这个长度不能改变
	//期望值：根据我们的实际需要，让计算机帮我们动态地去分配内存空间
	//切片：Slice：go语言当中有专门的一个用于动态的进行内存空间分配的数据类型


	//数组的定义：var arrName [size]type
	//切片的概念：是一种可变长的数组
	//切片的定义:var sliceName []type


	//长度：lenth，之切片当中，存放元素的个数
	//容量：cap，指切片当中可以容纳的元素的个数，容量参数默认可以不写，默认与len相同


	var slice1 []int//定义一个切片，名字是slice1，类型是int类型
	slice1 = []int{1,2,7,9}
	fmt.Println(slice1)
	fmt.Printf("slice1的长度是%d,slice1的容量是%d，slice1的类型是%T",len(slice1),cap(slice1),slice1)

	s2 := []int{4,6,2,8}
	fmt.Println(s2)

	//make函数创建切片
	s3 := make([]int,5,9)//创建的s3切片，类型是int，长度是5，容量是9
	fmt.Println(s3)
	fmt.Println(len(s3),cap(s3))
	//slice1 = make([]int,3,4)//3代表切片的长度，4代表切片的容量

	s4 := make([]string,3)//省略cap，cap和长度len相同
	fmt.Println(len(s4),cap(s4))

	//len()  cap()  make()都属于内置函数：
}
