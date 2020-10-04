package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)//特点：当程序真正执行时，程序走到scanln这一行，会一直等待用户输入，指导用户敲入了键盘的回车键，才会把对应的内容读取出来
	fmt.Printf("你的年龄是：%d\n",age)//第九行不会主动执行
}
