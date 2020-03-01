package main

import "fmt"

//运算符

func main() {
	const v = 200     //没有显式说明v的常量类型,则在二元运算时保证两个操作数类型一致,v就会自动转换为那个类型，位移操作除外
	var b byte = 127
	c := v + b        //这里执行时,v和b都是uint8,但是和超出了uint8的限制,就变成了327-256=71
	fmt.Printf("c 01type %T, value:%d\n", c, c)  //uint8, 71

	var e float32 = 215
	d := v + e
	fmt.Printf("d 01type %T, value:%f\n", d, d)

	//位移操作的右操作数,必须是无符号整数或者可以转换的无显示说明的常量
	var t uint8 = 10
	x := 1 << t
	fmt.Printf("x 01type:%T, value:%d\n", x, x)

	//如果位移操作的左操作数没有显式声明类型
	a := 1.0 << 3     //这里会将1.0转为int
	fmt.Printf("a 01type:%T, value:%v\n", a, a)

	var u uint8 = 3   //由于u是非常量
	//m := 1.0 << u     //由于m未提供具体类型,就根据1.0推断成float64,就不对
	//fmt.Printf("m 01type:%T, value:%v\n", m, m)

	var m int = 1.0 << u  //必须要int写明
	fmt.Printf("m 01type:%T, value:%v\n", m, m)


	//位运算符
	bit1, bit2 := 5, 3      //0101, 0011
	fmt.Printf("%04b\n", bit1&bit2)
	fmt.Printf("%04b\n", bit1|bit2)
	fmt.Printf("%04b\n", bit1^bit2)   //异或运算
	fmt.Printf("%04b\n", ^bit1)
	fmt.Printf("%04b\n", bit1&^bit2)  //这个注意,bit clear,清除第二个数中的1，达到清除标记位的作用
	fmt.Printf("%04b\n", bit1<<2)
	fmt.Printf("%04b\n", bit2>>2)

	//自增, a++只能作为独立的语句,不能用于表达式

	//指针,不要将内存地址和指针混为一谈
	//内存地址是内存中每个字节单元的唯一编号,而指针是一个实体,指针会分配内存空间,用来专门保存内存地址的变量
	//&用于获取对象地址, *用于间接引用目标对象,二级指针**T



}