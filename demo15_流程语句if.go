package main

import "fmt"

func main() {
	//if语句：在程序当中，负责做出选择
	var grade int
	//大于60分算及格
	grade = 60
	if grade >= 60{
		fmt.Printf("成绩是：%d,成绩合格\n",grade)
	}
	fmt.Println("程序执行结束")



	fmt.Println("程序开始执行")
	age := 28//年龄
	//中国法定结婚年龄：男22岁，女20岁
	if age >=22 {
		fmt.Println("该男子满足法定结婚年龄，可以准备彩礼了")
	}
	fmt.Println("程序结束")


}
