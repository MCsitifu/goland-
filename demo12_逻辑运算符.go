package main

import "fmt"

func main() {

	//数学里的逻辑符号：
	//最基本的操作符号
	//“与”操作：^
	//“或”操作:V
	//“非”操作:

	//逻辑运算符：与、或、非
	//与:&&(并且，同时的意思)
	//或：||（或者）

	b1 := false
	b2 := false
	b3 := true
	b4 := true
	//&&全真则真，一假则假
	result :=b1 && b2
	fmt.Println(result)
	result1 :=b1 && b3
	fmt.Println(result1)
	result2 := b3 && b4
	fmt.Println(result2)
	//||全假则假，一真则真
	result3 := b1 || b2
	fmt.Println(result3)
	result4 := b1 || b3
	fmt.Println(result4)
	result5 := b3 || b4
	fmt.Println(result5)//!ture -->flase
}
