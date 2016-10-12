package main

import "fmt"

func main() {
	var x int = 100
	var str string = "hello world"
	//声明初始化多个变量
	var i, j, k int = 1, 2, 3

	//不用指明类型，通过初始化值来推导
	var b = true
	// 等价与 var y int = 100
	y := 100

	const s string = "hello world"
	const s1 = "test const 类型推导"
	const pi float32 = 3.14
	fmt.Printf("x: %d\n", x)
	fmt.Printf("str: %s\n", str)
	fmt.Printf("i: %d, j: %d, k: %d\n", i, j, k)
	fmt.Printf("b: %t\n", b)
	fmt.Printf("y: %d\n", y)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("s1: %s\n", s1)
	fmt.Printf("pi: %f\n", pi)
	fmt.Printf("test end!\n")
}
