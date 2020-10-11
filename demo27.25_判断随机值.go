package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1、用户输入一个数
	//2、计算机随机生成一个数
	//3、比较这两个值的大小，当输入的数 > 计算机生成的数的时候，程序结束，否则一直输入和判断比较
//读取用户输入
	var num int
LOOP:fmt.Println("请输入一个整数")
	fmt.Scanln(&num)
	//计算机随机生成
	rand.Seed(time.Now().UnixNano())
num1 := rand.Intn(num*2)
fmt.Println("计算机生成的数：",num1)

//判断
	if num > num1 {
		fmt.Println("输入的数比生成的数值大，程序结束")
	}else {
		fmt.Println("不满足条件，继续比较判断")
		goto LOOP
	}

}
