package main

import "fmt"

func main() {
	//多维数组：多个维度的数据
	//多维数组中最常见的是二维数组
	//一维数组：[5]int{9,6,5,3,1}

	//两个班级：
	//区块链三班：a:80,b:85,c:90,d:95
	//区块链四班：e:82,f:88,g:80,h:83

	//二维数组的格式：[size1][size2]datetype{}
	//二维数组：
	arr := [2][4]int{
		{80,85,90,95},
		{82,88,80,83},
	}

	fmt.Println(arr[0][2])
//	fmt.Println(arr)
////数组的长度
//	fmt.Println(len(arr))
//	fmt.Println(len(arr[0]))

	for _,arr := range arr {
		for _,value := range arr {
			fmt.Printf("%d\t",value)
		}
		fmt.Println()
	}

	Arr := [2][2][2]int{//三维数组
		{{1,1},{1,1}},
		{{1,1},{1,1}},
	}
	fmt.Println(Arr)

}
