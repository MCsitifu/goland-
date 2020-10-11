package main

import "fmt"

func main() {
	//for循环整个表达式，都可以省略不写
	var i int
	for ;i<10 ;  {
		fmt.Println(i)
		i++
	}
	//for ; ;  {//死循环
	//	fmt.println("123)
	//}
	//fmt.Println("程序结束")//执行不了

	//for {
	//	fmt.Println("123")
	//}

	//for true {
	//	fmt.Println("123")
	//}
}
