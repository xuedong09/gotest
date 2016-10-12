package main

import (
	"fmt"
	"./mypkg"
//	"./mypkg/test"
)

func main() {
	fmt.Println("testmypkg: I will test my package")
	mypkg.Test()
	mypkg.Test2()
}
