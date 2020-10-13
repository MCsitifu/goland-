package main

import "fmt"

func main() {
	//选择排序：特点：无论数组多大或者多小，数组当中肯定有一个最大值（最小值）
	//问题：只是这个最大值或者最小值我们不知道在那个元素上而已
	//[5]int{9,6,5,3,1}
	//目标结果：[5]int{1,3,5,6,9}
	//设置一个最小数，拿最小数与元素挨个进行比较
	arr := [5]int{20,18,15,10,5}
	//min := 0//定义一个变量，作为整个数组中的最小数
	for index:=0 ; index<len(arr) ; index++  {
		//根据index，获取到arr数组的元素，赋值给min变量，作为可能的最小值
		min := arr[index]
		//index+1 是index的下一位
		for i:=index+1; i<len(arr);i++  {//内层for循环：
			//如果下一个位置上的元素  比  我们的min 要小，则要交换两者的位置
			if arr[i]<min {
				arr[index],arr[i]=arr[i],arr[index]//操作多元素的交换
				min=arr[i]//把最小值设置成新找到的这个值
			}
		}
	}
	fmt.Println(arr)
}
