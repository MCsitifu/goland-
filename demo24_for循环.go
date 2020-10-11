package main

import "fmt"

func main() {
	//循环语句：条件满足的时候，一直执行
	//for 表达式1； 表达式2 ；表达式3{
	//要循环执行的程序体
	//}
	//打印自己的名字三遍
	for num :=0 ; num < 3;num++{
		fmt.Println("name")
	}
	fmt.Println("程序结束")


	sum := 0;//1-10之和
	num1 := 0
	for num1=1 ; num1 <11 ; num1++  {
		sum += num1
	}
	fmt.Println(sum)

	for i := 1;i<=20;i++{//大于0的，20以内的3的倍数
		if i%3==0 {
			fmt.Println(i)
		}
	}

	for n:=1;n<=10;n++{//输出10以内的奇数
		if n%2!=0{
			fmt.Println(n)
		}
	}

	num2:=1
	for o:=1;o<=5;o++{//5的阶乘
		num2*=o
		fmt.Println(num2)
	}
}


