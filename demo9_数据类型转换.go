package main

import "fmt"

//基本数据类型：数值型（整形，浮点型，复数）、布尔类型（bool）、字符串类型（string）
//复合数据类型（派生类形）：
//
func main()  {
//数据类型转换：在需要的时候，将数据类型进行强制转换，数据类型转换语法：Type(value)
	var num int8
	num = 100

	var num1 int16
	num1 = 20

	//sum := int16(num) + num1//解决方案一
	sum := num +int8(num1)//解决方案二
	fmt.Println(sum)
	//数据类型的转换只能在同类型中进行转换
	//age := 18
	//result := int8(age)
	//fmt.Println(result)

	grade := 86.54
	result := int8(grade)
	fmt.Printf("数据类型是:%T\n",grade)
	fmt.Printf("result的数据类型是:%T\n",result)

	//别名：int8、int16、int32、int64
	//uint8别名是byte：字节		1byte = 8个bit
	var a uint8 = 3//大名
	var b byte = 5//小名
	sum1 := a + b
	fmt.Println(sum1)

	//关于float的小数点位数
	var pai = 3.1415926
	fmt.Printf("pai的数值是:%f\n",pai)

	//保留两位小数
	fmt.Printf("保留两位小数的值是:%.2f\n",pai)
	//保留三位小数
	fmt.Printf("保留三位小数：%.3f",pai)

	//字符串 string
	name := "jinsheng"//看一下字符串的长度
	fmt.Println("字符串name的长度是:",len(name))

	address := "南昌省上饶市"
	length := len(address)
	fmt.Println("地址字符串的长度是",length)

	//截取字符串
	name1 := "tiechuimeimei"//铁锤妹妹
	//如何得到tiechui这个部分字符串
	tiechui := name1[0:7]//前闭后开区间[0：7）
	fmt.Println(tiechui)

	//截取meimei这个部分区间
	xiaojiejie := name1[7:len(name1)]
	fmt.Println(xiaojiejie)

	//如果是切到字符串末尾，则可以省略不写
	xiaomeimei := name1[7:]
	fmt.Println(xiaomeimei)

	//如果是从头开始切，也可以省略不写
	tiechui1 := name1[:7]
	fmt.Println(tiechui1)

	chuimei := name1[3:10]
	fmt.Println(chuimei)
}
//格式化打印：%d是整数，%T是数据类型，%f是浮点数，%.nf保留浮点数的小数点后n位小数