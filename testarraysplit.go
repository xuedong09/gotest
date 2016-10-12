package main

import "fmt"

func main() {
	//test array split
	a := [5] int {1, 2, 3, 4, 5}
	b := a[2: 4] // [2, 4)
	fmt.Println("a: ", a)
	fmt.Println("a[2: 4]: ", b)

	b = a[:4] // [0,4)
	fmt.Println("a[:4]: ",b)

	b = a[2:] // [2,]
	fmt.Println("a[2:]: ", b)
}
