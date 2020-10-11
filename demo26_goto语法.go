package main

import "fmt"

func main() {
	//DEMO:for num:=1;num<=5 ; num++ {
	//	for i:=1;i<=num ;i++  {
	//		if num==2 {
	//			break DEMO//默认跳出内层循环
	//			continue//默认跳出外层循环
	//		}
	//		fmt.Print("* ")
	//	}
	//	fmt.Println()
	//}

	//goto语法：可以无条件的强制转移到goto语句所指定行进行程序执行

	var a = 10
	LOOP:
	for a < 20 {
		if a==15 {
			//fmt.Print(a)//不断打印a的值
			goto LOOP//程序会跳转到20行继续执行
		}
		fmt.Printf("a的值：%d\n",a)
		a++
	}

}
