package main

import "fmt"

func main() {
	//打印1000以内的水仙花数
	//水仙花数：比如153 = 1*1*1+5*5*5+3*3*3
	var i int
	for i = 1;i<=1000 ;i++  {

		if i <=9 && i==i*i*i {
		fmt.Println(i)
		}
		if i>=10 && i<=99 && i == (i%10)*(i%10)*(i%10)+(i/10)*(i/10)*(i/10) {
			fmt.Println(i)
		}
		if i>=100 && i<=999 && i == (i%10)*(i%10)*(i%10)+(i/10%10)*(i/10%10)*(i/10%10)+(i/100)*(i/100)*(i/100) {
			fmt.Println(i)
		}
	}
}
