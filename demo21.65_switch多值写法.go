package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	switch num {
	case 1,2,3 :
		fmt.Println("多值写法完成")
	}
}
