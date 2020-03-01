package main

import "fmt"

//流程控制

func main() {
	//if ... else ...

	x := 3
	if x > 5 {
		fmt.Println("x > 5")
	} else if x < 5 && x > 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("ojbk")
	}

	//支持初始化语句
	n := 5
	var ninit = func(num int) {
		n += num
	}

	if ninit(n); n >= 10 {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}

	//尽量减少代码块嵌套

	//switch, 还支持初始化语句

	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)

		fallthrough //继续执行下一个case语句块的内容,直接执行，不用判断
	case 6:
		x += 20
		fmt.Println(x)
	}

	//for

	//goto
	//使用goto之前,需要先定义标签,未被使用的标签会错误

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i > 5 {
			goto exit
		}
	}
exit:
	fmt.Println("exit")

	//break和continue

}
