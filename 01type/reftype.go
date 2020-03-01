package main

import "fmt"

//引用类型(reference 01type)特指slice,map,channel这三种预定义类型

//引用类型必须用make来创建

func mkslice() []int {
	s := make([]int, 0, 10)
	fmt.Printf("s addr:%p\n", s)
	s = append(s, 100)
	fmt.Printf("then s addr:%p\n", s)

	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	fmt.Printf("m addr:%p\n", m)
	m["a"] = 1
	return m
}

func main() {
	s := mkslice()
	fmt.Printf("now s addr:%p\n", s)

	m := mkmap()
	fmt.Printf("now m addr:%p\n", m)

	p := new(map[string]int)      //new只是分配该类型指针的内存,没有分配它自身存储的内存
	fmt.Printf("p addr:%p, value:%v\n", p, p)
	//q := *p
	//q["a"] = 1    // 不可以panic: assignment to entry in nil map
}