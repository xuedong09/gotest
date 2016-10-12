package main

import "fmt"
import "reflect"

func test_if(i int) {
	if i %2 == 0 {
		fmt.Println("even: ", i)
	} else if i %2 == 1{
		fmt.Println("odd: ", i)
	} else {
		fmt.Println("never ")
	}
}

func test_t(i interface {}) {
	i_t := reflect.TypeOf(i).Name()
	if i_t == "string" {
		fmt.Printf("%s is a string not int\n", i)
		return
	} else if i_t == "int" {
		fmt.Printf("%d is a int\n", i)
	}else {
		fmt.Printf("input is type %s not int: ",i_t)
		fmt.Println(i)
		return
	}

}

func test_t2(i interface {}) {
	switch i.(type) {
	case string:
		fmt.Printf("%s is a string not int\n", i)
	case int:
		fmt.Printf("%d is a int\n", i)
	default:
		fmt.Printf("input is type not int\n")
	}
}
func main() {


	a := [10]int {1, 2, 3, 4 , 5, 6, 7 , 8, 9, 10}
	s := "abc"
	f := 0.12
	d := 10
	// 判断奇偶打印
	for i, _ := range a {
		fmt.Println(i)
		test_if(i)
	}

	fmt.Println("for 2")
	for i:=0; i<10; i++ {

		test_if(i)
	}

	fmt.Println("for 3")
	x := 1
	for x != 0 {
		if x >10 {
			break
		}
		x++
	}
	fmt.Println(">>>test if else if ")
	test_t(s)
	test_t(d)
	test_t(f)
	test_t(a)
	fmt.Println(">>>test switch ")
	test_t2(s)
	test_t2(d)
	test_t2(f)
	test_t2(a)
	fmt.Println("for ever")
	for {
		//fmt.Println("1")
	}
}
