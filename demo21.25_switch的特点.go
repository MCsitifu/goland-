package main

import "fmt"

func main() {
	var num int
	fmt.Println("请输入1-4之间的一个数值；")
	fmt.Scanln(&num)
	switch num {
	case 1 :
		fmt.Println("第一季度")
	case 2 :
		fmt.Println("第二季度")
	case 3 :
		fmt.Println("第三季度")
	case 4 :
		fmt.Println("第四季度")
	}
	fmt.Println("程序执行结束")
	//default可有可无，没有default不会报错
	//case的值不能重复，否则程序会报错：Duplicate case重复错误
}
