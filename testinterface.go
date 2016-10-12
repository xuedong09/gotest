package main

import "fmt"
import "reflect"

func t(i interface{}) {

	fmt.Println("i is : ", reflect.TypeOf(i).Name())
	switch i.(type) {
	case string:
		fmt.Printf("%s: is a string\n", i)
	case int:
		fmt.Printf("%d: is a int\n", i)
	}

}

func main() {

	s := "abc"
	d := 10
	t(s)
	t(d)
	fmt.Println("type: ", reflect.TypeOf(s))
}
