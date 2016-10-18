package main

import (. "fmt")
//import "time"
import "runtime"

var c int

func main() {
	//runtime.GOMAXPROCS(10)
	defer Println("defer==========")
	Println("0000000", c)
	Println(runtime.GOMAXPROCS(10), runtime.NumCPU())
	Println(runtime.GOMAXPROCS(0), runtime.NumCPU())

}

func init() {
	c = 100

}
