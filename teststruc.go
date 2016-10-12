package main

import "fmt"


type Person struct{

	name string
	age int
}

func (p *Person) interview() string {

	s := fmt.Sprintf("Hello! I am %s and I am %d years old.\n", p.name, p.age)
	return s
}

func main() {

	person1 := Person{"Tom", 30}
	person2 := Person{name: "xd", age: 30}
	fmt.Println("person1: ", person1)
	fmt.Println("person2: ", person2)

	p3 := &person1
	fmt.Println("p3: ", p3)

	fmt.Println("name, age:  ", p3.name, p3.age)
  fmt.Println(p3.interview())
}
