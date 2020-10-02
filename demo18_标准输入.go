package main

import "fmt"

func main() {

	var x int
	var y int
	fmt.Println("请输入两个整数，分别作为x变量和y变量的值")
	fmt.Scanln(&x,&y)//读取键盘输入的值，并赋值给对应的变量
	fmt.Printf("x变量的值是：%d，y的值是：%d\n",x,y)

	fmt.Println("x的内存地址是：",&x)//通过&符号可以查看变量的内存地址
	fmt.Println("y的内存地址是：",&y)//内存地址：相当于家庭住址、门牌号
	//fmt.printf表示格式化输入
	//fmt.scanf表示输出
	fmt.Scanf("%d,%d",&x,&y)
	fmt.Printf("x变量的值是：%d,y的值是：%d\n",x,y)

}
