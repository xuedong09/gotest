package main

import (
	"fmt"
	"time"
)

func comsumer(id int, data chan int , done chan bool){

	for x := range data {

		fmt.Println(id, "recv:", x)
	}
	done <- true
}

func producer(data chan int) {

	for i := 0; i<100000; i++ {

		data<-i
	}
	close(data)
}


func main() {

	done := make(chan bool)
	data := make(chan int)

	go comsumer(1, data, done)
	go comsumer(2,data, done)
	go producer(data)
	<-done
	fmt.Println(time.Second)
}
