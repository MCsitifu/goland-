package main

import "fmt"

func main() {
	//if  {
	//
	//}else {
	//
	//}
	//结婚年龄：国家法定结婚年龄男生不得低于22周岁，女生不得低于20周岁
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age) //接受键盘输入，并把输入的值赋值给变量age
	var sex int      //1表示男生2表示女生
	fmt.Println("请输入你的性别（男生输入1，女生输入2）")
	fmt.Scanln(&sex) //接受键盘输入，并把输入的值赋值给变量sex

	//先判断性别
	if sex == 1 { //男的
		if age >= 22 {
			fmt.Println("可以结婚了")
		}else {
			fmt.Println("对不起，还不能结婚")
		}
	} else {//女的
		if age >= 20{
			fmt.Println("可以结婚了")
		}
		fmt.Println("对不起，还不能结婚")
	}
}
