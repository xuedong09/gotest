package main

import "fmt"

func main() {
	var i int = 1
	var pint *int = &i

	fmt.Printf("i=%d\tpint=%p\t*pint=%d\n", i, pint, *pint)

	*pint = 2
	fmt.Printf("i=%d\tpint=%p\t*pint=%d\n", i, pint, *pint)

	i = 3
	fmt.Printf("i=%d\tpint=%p\t*pint=%d\n", i, pint, *pint)
}
