package main

import "fmt"

func main() {
	for i:=1;i<=5 ;i++  {
		for o:=1;o<=5 ;o++  {
			fmt.Print("hello\t")
		}
		fmt.Println()
	}

	for p:=1;p<=9 ;p++  {//99乘法表
		for l:=1;l<=p ;l++  {
			fmt.Printf("%d * %d = %d\t",l,p,p*l)
		}
		fmt.Println()
	}



	for num4:=1;num4<=5 ;num4++  {

		for num5:=1;num5<=5-num4 ;num5++  {
			fmt.Print(" ")
		}

		for num1 := 1; num1 <= num4*2-1; num1++ {
			fmt.Print("*")
		}
		fmt.Println()
	}


	for num:=1;num<=4 ;num++ {
		for num5:=1;num5<=num ;num5++  {
			fmt.Print(" ")
		}
		for num2 := 1; num2 <= 9-num*2; num2++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
