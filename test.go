package main

import (. "fmt")
//import "time"
//import "runtime"

var c int

func test(a, b int) {

	Println(a/b)
}

func main() {
	//runtime.GOMAXPROCS(10)
	defer Println("defer==========")
	Println("0000000", c)
	test(1,0)
	//Println(runtime.GOMAXPROCS(10), runtime .NumCPU())
	//Println(runtime.GOMAXPROCS(0), runtime.NumCPU())

}

func init() {
	c = 100

}
