package main

import "fmt"

//初始化,主要是数组,切片,字典,结构体这些类型在初始化时要指明具体标签
func main() {
	type data struct {
		name string
		perm byte
	}

	var a data = data {
		name : "filename",
		perm : 7,
	}

	b := data {
		"hao",
		1,
	}

	c := []int {
		1,
		2,
		3,
	}

	fmt.Println(a, b, c)

	d := map[string]int {
		"a":1,
		"b":2,
	}

	d["c"] = 1
	fmt.Println(d)


}