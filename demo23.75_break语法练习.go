package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	switch num {
	case 3,4,5:
		if num ==3 {
			fmt.Println("大学没开学")
			break//case直接结束
		}
		fmt.Println("3月份过去了")
		if num ==4 {
			fmt.Println("你输入的是4")
		}
		fmt.Println("4月份过去了")
		if num ==5 {
			fmt.Println("你输入的是5")
		}
	case 6,7,8:
	default:
		fmt.Println("程序结束")
	}
}
