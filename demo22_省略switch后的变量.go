package main

import "fmt"

func main() {
	switch  {//如果switch后面省略了变量，则默认表示true
	case true:
		fmt.Println("case分支是true")//默认
	case false:
		fmt.Println("case分支是false")
	default:
		fmt.Println("case分支是default")
	}

}
