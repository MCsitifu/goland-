package main

import "fmt"

func main() {
	//break,和continue在双层for循环中的使用
	for i := 1;i <5 ;i++  {
		for j:=1 ; j<5  ; j++  {
			if j == 2 {
				break//结束内层循环
				//continue跳出内层循环
			}
			fmt.Println("i",i,"j",j)
		}
	}
}
