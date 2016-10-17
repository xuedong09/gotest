package main

import "fmt"
import "time"

var c int

func counter() int {

	c++
	return c
}

func main() {

	exit := make(chan struct{})
	a := 100

	go func(x, y int) {

		time.Sleep(time.Second)
		fmt.Println("go: ", x, y)
		close(exit)
		time.Sleep(time.Second)
		fmt.Println("finish go")
	}(a, counter())
	<-exit
	fmt.Println("main exit")
}
