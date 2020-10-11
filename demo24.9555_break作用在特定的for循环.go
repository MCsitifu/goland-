package main

import "fmt"

	func main() {
		//break,和continue在双层for循环中的使用
	OUT : for i := 1;i <4 ;i++  {
			for j:=1 ; j<3  ; j++  {
				if j == 2 {
					//break OUT//结束内层循环
					continue OUT//跳出外层的当前循环，继续下一次执行
				}
				fmt.Println("i",i,"j",j)
			}
		}
		fmt.Println("程序结束")
	}
