package main

import "fmt"

func main()  {
	//圆形：半径是a=5，求圆的周长d？
	//圆的周长：d=2*3.14*a
	var a float32//先声明
	a = 5
	var d float32
	d = a * 3.14 * 2
	fmt.Println("圆形的周长",d)
	fmt.Printf("周长d的数据类型:%T",d)
}
