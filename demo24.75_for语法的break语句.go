package main

import "fmt"

func main() {
	//break语句：switch语法当中，结束当前case的执行：
	//num :=18//10
	//for ;num<25 ;num++  {
	//
	//	fmt.Println("当前年龄是",num)
	//	if num>=22 {
	//		fmt.Println("以达到法定结婚年龄，终止循环")
	//		break
	//	}

		age := 18
		for ;age <25 ;age++  {
			if age<20 {
				fmt.Println("对不起，你还不能结婚")
				//continue//作用：for循环内的循环体，continue以后的内容不执行了，进入下一次循环
			}
			fmt.Println("你的年龄是：",age)
			if age>22 {
				fmt.Println("恭喜你，可以结婚了")
			}
		}
		fmt.Println("程序结束")
	}
