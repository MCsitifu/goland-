package main

import "fmt"

func main() {
	//打印100以内能被4和6整除的数
	//count :=0
	for i:=1;i<=100 ;i++  {

		if i >= 90 || i <=30{
			continue
		}//contiue是结束本次循环，继续执行下一次循环

		if i%4 == 0 && i%6 == 0 {
			//fmt.Println(i)
			//count++
			//if count==6 {
			//	break
			//}//只输出前6个

			//if i>=30 && i<=90{
			//fmt.Println(i)
			//}//输出数字在30-90之间
			fmt.Println(i)
		}

	}
}