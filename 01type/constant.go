package main

import "fmt"

//常量,运行时不会改变的值,就是在编译器就可以确定的字符,字符串,数字,布尔值等
//可以在初始化时指定类型或者由编译器进行推断

func main() {
	//未被使用的常量不会引发编译错误
	const x, y int = 1, 2
	const z = false
	const c = '中'       //rune
	
	const (
		i, f = 1, 3.14   //int,float64   默认的
		b = false
	)

	{
		const x = "ok"        //在不同作用域中可以定义同名的常量
		fmt.Println(x)
	}


    //如果显示说明了常量类型，需要注意
    //1.左右类型一致 2.不可以超过类型的限制
    const (
    	x1, y1 int = 99, -128
    	b1 byte = byte(x1)             //没有超过byte的限制
    	//z1 uint8 = uint8(y1)         //不可以,已经超过uint8的0~128的限制
	)

	fmt.Println(x1, y1, b1)

	//枚举,借助iota实现一组自增常量值来实现枚举类型
	const (
		a1 = iota
		a2
		a3
	)
	fmt.Println(a1, a2, a3)

	const (
		_ = iota
		KB = 1 << (10 * iota)    // 1 << (10 * 1)
		MB
		GB
	)

	fmt.Println(KB, MB, GB)

	//可以使用多个常量定义中使用iota,他们各自单独计数,只要确保后续的
	//每行的变量中列的数量相同
	const (
		h1, h2 = iota, iota * 10
		h3, h4
		h5, h6
	)

	fmt.Println(h1, h2)
	fmt.Println(h3, h4)
	fmt.Println(h5, h6)

	//中断iota的情况
	//中断后从恢复的那个值开始是按当前是第几行-1开始的
	const (
		g1 = iota + 1      //1
		g2             //2
		g3             //3
		g4 = 100       //100
		g5 = 200       //200
		g6             //200
		g7 = iota      //显示恢复iota,此时是6
		g8             //7
	)

	fmt.Println(g1, g2, g3, g4, g5, g6, g7, g8)

	//指定类型
	const (
		k1 = iota
		k2 float32 = iota
		k3 = iota           //不显示的指定，就是int
	)

	fmt.Printf("k1 01type:%T, value:%v\n", k1, k1)
	fmt.Printf("k2 01type:%T, value:%v\n", k2, k2)
	fmt.Printf("k3 01type:%T, value:%v\n", k3, k3)

	const t = "hello"
	const o = 0x200
	//fmt.Println(t, &t)       //不可以取常量的地址

}