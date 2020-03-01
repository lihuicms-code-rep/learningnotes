package main

import (
	"fmt"
	"math"
	"strconv"
)

//基本类型

func main() {
	//8,16进制,及格式化输出
	a, b, c := 100, 0244, 0x1af
	fmt.Println(a, b, c)
	fmt.Printf("0b%b, %#o, %#X\n", a, b, c)

	fmt.Println(math.MaxInt8, math.MinInt8)

	//标准库strconv可以在不同进制的字符串进行转换

	//parseInt,  str  ===> num
	a1, _ := strconv.ParseInt("0111101010", 2, 64)
	a2, _ := strconv.ParseInt("04567", 8, 64)
	a3, _ := strconv.ParseInt("12f", 16, 64)
	fmt.Println(a1, a2, a3)

	//formatInt,  num ===> str
	fmt.Println("0b" + strconv.FormatInt(480, 2))
	fmt.Println("0" + strconv.FormatInt(2423, 8))
	fmt.Println("0x" + strconv.FormatInt(303, 16))

	//有关go中浮点数精度的问题,以及0.1+0.2==0.3的问题

	//go语言中的两个类型别名
	var b1 byte = 123
	var b2 uint8 = 123
	var b3 = b1 + b2

	fmt.Printf("b1 01type:%T, vlaue:%d\n", b1, b1)     //打印值都是uint8
	fmt.Printf("b2 01type:%T, value:%d\n", b2, b2)
	fmt.Printf("b3 01type:%T, value:%d\n", b3, b3)

}