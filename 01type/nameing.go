package main

import (
	"fmt"
	"strconv"
)

//命名,对变量,常亮,函数,自定义类型进行命名

//建议
//区分大小写,使用驼峰(camel case),局部变量有限使用短名
//专有名词全部大写,escapeHTML,xxxTCP,

func main() {

	//1.局部变量简短为好

	var c int    //用c代替count
	for i := 0; i < 10; i++ {  //用i代替index
		c++
	}

	fmt.Println(c)

	//2.空标识符
	x, _ := strconv.Atoi("12")    //忽略返回值err
	fmt.Println(x)
}