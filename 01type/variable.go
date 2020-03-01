package main

import "fmt"

//变量,用来标识一段或多段用来存储数据的内存,变量的类型决定了变量内存长度和存储格式
//由于内存分配发生在运行期间,所以在编码阶段我们就用一个易于阅读的名字来标识这段内存。

func main() {

	//1.使用var定义

	var x int            //自动初始化为0
	var y = true         //显示提供初始化值可以省略变量类型,由编译器推断
	var a, b string      //可一次定义多个变量
	var c, d = 12, "iam" //并且可以初始化为不同的类型值

	var (                //以组的方式整理多行变量的定义
		e int
		f string
		g bool
	)

	fmt.Println(x, y, a, b, c, d, e, f, g)

	//2.简短方式进行变量定义和初始化

	//定义变量的同时,显式初始化,不用启动变量类型,只能在函数内部
	//同时要注意,局部变量和全局变量同名的问题
	x1 := 10
	x2, x3 := false, "ok"

	fmt.Println(x1, x2, x3)

	//简短模式并不是总是重新定义变量,也可以退化为赋值
	y1 := 10
	y1, y2 := 20, "ok"
	fmt.Println(y1, y2)

	//但是这种情况前提是:最少有一个新的变量被定义，以及在同一作用域
	//下面这种不在同一作用域的就不会被视为退化
	z1 := 10
	fmt.Printf("z1 addr:%p, value:%d\n", &z1, z1)
	{
		z1, z2 := 20, "ok"      //这里就是新的变量z1了
		fmt.Printf("z1 addr:%p, value:%d\n", &z1, z1)
		fmt.Printf("z2 addr:%p, value:%s\n", &z2, z2)
	}

	fmt.Printf("z1 addr:%p, value:%d\n", &z1, z1)

	//这个可以常用在err变量的使用上

	//3.多变量赋值

	p, q := 1, 2
	p, q = q + 2, p + 3          //先计算好右值,再赋值
	fmt.Println(p, q)            //4, 4

	//两值交换
	m, n := 1111, 2222
	m, n = n, m
	fmt.Println(m, n)

}
