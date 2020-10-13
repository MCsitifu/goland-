package main

import "fmt"

func main() {
	//数组的排序：按照一定的规则，对数组当中的元素进行有规则的排序
//	排序分类：
//	升序：从小到大
//	降序：从大到小
//	例如：[5]int{2,21,34,18,92}
//	升序：[5]int{5,18,21,34,92}
//	降序：[5]int{92,34,21,18,5}
//目标：通过所学的go语言语法知识，编写程序，让数组变成有序的

//var str = [3]string{"aaa","aa","a"}//打印string类型
//fmt.Println(str)
//fmt.Printf("%T\n",str)

//	var INT [10]int//输入10个整数值，并打印
//fmt.Println("请输入10个整数类型的数值")
//fmt.Println(INT)
//	for index:=0;index<=len(INT)-1;index++  {
//		fmt.Printf("请输入第%d个元素：\n",index+1)
//		fmt.Scanln(&INT[index])
//	}
//	for _,value:=range INT{
//		fmt.Printf("%d\t",value)
//	}

	//排序算符：冒泡排序、选择排序、快速排序、希尔排序、归并排序
	//冒泡排序规则：比较两个相邻的数，小的在前，大的在后，进行排序
	arr := [5]int{23,15,8,32,9}
	//目标：从小到大[8,9,15,23,32]
	fmt.Println(len(arr))
	for i:=1 ; i<=len(arr) ; i++ {
		for index:=0; index < len(arr)-i;index++{

			if arr[index] > arr[index+1] {
				//num := arr[index]
				//arr[index] = arr[index+1]
				//arr[index+1] = num

				arr[index],arr[index+1]=arr[index+1],arr[index]
			}
		}
	}
		fmt.Println(arr)
}
