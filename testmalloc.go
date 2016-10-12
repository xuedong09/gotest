package main

import "fmt"

func main() {

	var p *[]int = new([]int)
	var v []int = make([]int, 10)

	fmt.Println("p, *p: ", p, *p)
	fmt.Println("v: ", v)

	*p = make([]int , 5 , 10)// 初始长度为10并预留十个长度
	fmt.Println("p, len(p), cap(p): ", p, len(*p), cap(*p))
	*p = (*p)[:cap(*p)]
	fmt.Println("扩容：p, len(p), cap(p): ", p, len(*p), cap(*p))
	fmt.Println("p[2]: ", (*p)[2])
	(*p)[5] = 1
	fmt.Println("p: ", *p)

	//习惯用法
	v = make([]int, 10)
	fmt.Println(v)
}
