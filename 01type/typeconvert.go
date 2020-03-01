package main

import (
	"fmt"
)

//类型转换,隐式转换造成的问题远大于它带来的好处。
//Go强制要求使用显示类型转换

func main() {
	x := 10
	p := &x
	fmt.Println(*p)
	q := (*int)(&x)    //如果转换的目标是指针,单向通道或者没有返回值的函数类型,必须使用括号
	fmt.Println(*q)
}