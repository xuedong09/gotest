package main

import "fmt"


func max(a int, b int) int {

	if a>b {
		return a
	}
	return b
}

func multi_ret() (int, int) {

	return 1, 2
}

func sum(nums ... int) {
	fmt.Println("multi params: ", nums)

}

func nextNum() func() int {

	i, j := 1, 1
	return func() int {
		var tmp = i+j
		i, j = j, tmp
		return tmp
	}
}

func main() {

	fmt.Println("max of 1 and 2: ", max(1, 2))
	a,b := multi_ret()
	fmt.Println("multi return: ", a, b)
	sum(1, 2, 3, 4 ,5)
	nextNumFunc := nextNum()
	for i:=0; i<10; i++ {
		fmt.Println("nextnum: ", nextNumFunc())
	}

}
