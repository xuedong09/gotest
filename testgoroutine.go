package main

import (
	"fmt"
	"time"
)

func task(id int) {

	for i := 0; i<5; i++ {

		fmt.Printf("task id: %d: %d\n", id, i)
		time.Sleep(time.Second)
	}
}

func main() {

	defer fmt.Println("=====")
	go task(1)
	go task(2)
	time.Sleep(time.Second * 10)
}
