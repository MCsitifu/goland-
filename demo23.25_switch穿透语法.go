package main

import "fmt"

func main() {

	var num int
	fmt.Println("请输入月份：（1-12之间）")
	fmt.Scanln(&num)

	switch num {

	case 1:
		fmt.Println("冬季")
	case 2:
		fmt.Println("冬季")
	case 3:
		//fmt.Println("春季")
		fallthrough//从当前case中，强制穿透到下一个case中，执行程序
	case 4:
		//fmt.Println("春季")
		fallthrough
	case 5:
		fmt.Println("春季")


	}

}
