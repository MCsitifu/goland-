package main

import "fmt"

func main() {
	//平行四边形
	//for i := 1 ; i<=6 ; i++ {
	//	for j:=1 ; j<=6-i ; j++  {
	//		fmt.Print(" ")
	//	}
	//	//打印*号
	//	for o:=1 ; o<=9  ;o++  {
	//		fmt.Print("*")
	//	}
	//	fmt.Print("\n")
		//fmt.Println()//自带换行
	//}

	//空心三角形
	//1.打印实行三角形2.把中间星星换成空格
	for i := 1 ; i <= 6 ; i++ {
		for o:=1 ; o <= i ; o++  {
			//fmt.Println("*")//等腰直角三角形

			//最后一行比较特殊，要全部打印星星，因为最后一行是三角形的边
			if i==6 {
				fmt.Print("* ")
				continue
			}
			if o==1 || o==i{
				fmt.Print("* ")
			}else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

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



}
