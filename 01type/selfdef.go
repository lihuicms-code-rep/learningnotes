package main

import (
	"fmt"
)

//自定义类型
//使用关键字type定义用户自定义类型,包括基于现有基础类型创建,或者结构体，函数等

type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func main() {
	f := read | write
	fmt.Printf("%b\n", f)

	//与var,const类似，多个type可以定义成组
	type (
		user struct {
			Name string
			Age int
		}

		event func(string) bool
	)

	u := user{"lihui", 11}
	fmt.Println(u)

	var f1 event = func(s string) bool {   //event函数类型的具体实现
		if s == "ok" {
			return true
		}
		return false
	}

	fmt.Println(f1("ok" ))

	//即使是指定了基础类型,也只是表明他们具有相同底层数据结构,是两种不同的类型
	type myint int
	var a myint = 10
	var b int = 10
	fmt.Printf("a 01type %T\n", a)    //main.myint
	fmt.Printf("b 01type %T\n", b)    //int


	//未命名类型
	//数组,切片,字典,通道这些类型与具体的元素类型以及长度等有关,称为未命名类型
	//具有相同声明的未命名类型视为同一类型



}
