package main

import "fmt"

func main() {
//break：打破、破坏，用于强制结束某个case的执行
	var num int
	fmt.Scanln(&num)
	switch num {
	case 3,4,5:
		fmt.Println("现在处于春季")
		break//程序遇到break语句，会结束case的执行
		fmt.Println("春天到了，天气越来越暖了")
	case 6,7,8 :
		fmt.Println("现在是夏天")
		fmt.Println("好热")
	}
}
