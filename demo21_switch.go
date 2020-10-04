package main

import "fmt"

func main() {
	//程序结构：顺序结构、选择结构
	//选择结构：if分支语句、switch语句
	//switch：选择
	//case：情况、事例
	//语法结构：
	//	switch  变量  {
	//		case   值1 ：
	//		case   值2 ：
	//		case   值3 ：
	//		case   值4 ：
	//		case   值5 ：
	//		......
	//		default
	//			}
	var age int
	fmt.Println("请输入您的年龄:")
	fmt.Scanln(&age)
	switch age {
	case 18 :
		fmt.Println("恭喜你成年了")
	case 20 :
		fmt.Println("如果你是女生，恭喜你可以结婚了，如果是男生，对不起还要再等等")
	case 22 :
		fmt.Println("不论你是男是女，都达到了法定结婚年龄")
	default:
		fmt.Printf("你的年龄是：%d\n",age)



	}
}