package main

import "fmt"

func main() {
	a := 10
	if a > 0 {
		fmt.Println("a是个正数")
	}

	//go语言支持合并写法
	//if 条件表达式 表达式只有两种：true、flase
	if b := 6;b > 0{
		fmt.Printf("b的值是:%d\n",b)
		fmt.Println("b是个正数")
	}

	if a>5 {
		c := 85
		fmt.Println(c)
	}
	//fmt.Println(c)

}
