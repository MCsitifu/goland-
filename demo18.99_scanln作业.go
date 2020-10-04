package main

import "fmt"

func main() {
	//var username string		作业1
	//var password string
	//fmt.Println("您正在使用千锋教辅系统")
	//fmt.Println("请输入正确的账号：")
	//fmt.Scanln(&username)//账号：js223499848	密码：123456..
	//fmt.Println("请输入正确的密码：")
	//fmt.Scanln(&password)
	//if username=="js2234399848" && password == "123456.." {
	//	fmt.Println("正在登录,请等待...")
	//}else {
	//	fmt.Println("账号或密码错误，请重新输入")
	//}


	var height int
	var weight int
	fmt.Println("请输入您的身高：")
	fmt.Scanln(&height)
	fmt.Println("请输入您的体重：")
	fmt.Scanln(&weight)
	fmt.Printf("您的身高为：%d，您的体重为：%d\n",height,weight)
	var contrast float32
	contrast = float32(weight) / float32(height)
	if contrast <= 0.675 && contrast>=0.525{
		fmt.Println("哇！您的身高体重属于标准的黄金比例，请保持呀！")
	}else if contrast > 0.675{
		fmt.Println("您的体重过高，请少吃高脂类、高糖类食物，多注意饮食卫生！！")
	}else if contrast < 0.525{
		fmt.Println("您的体重过低，请合理搭配饮食，千万别挑食哦！！")
	}
	fmt.Printf("您的体重身高比为：%f",contrast)





}