package main

import "fmt"

func main() {

	m := make(map[string] int) //使用make创建一个空的map
	fmt.Println(m)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)
	fmt.Println("len: ", len(m))

	v := m["two"]
	fmt.Println("v: ", v)

	delete(m, "two")
	fmt.Println("delete two: ", m)

	m1 := map[string] int {"a": 1, "b": 2, "c": 3}
	fmt.Println("m1: ", m1)

	for key, val := range m1 {

		fmt.Printf("key: %s val: %d\n", key, val)
	}
}
