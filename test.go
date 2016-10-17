package main

import (. "fmt")
import "time"

func main() {

	defer Println("defer==========")
	Println("0000000")
	time.Sleep(5000000000000000)
}
