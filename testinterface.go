package main

import "fmt"
import "reflect"
import "math"

func t(i interface{}) {

	fmt.Println("i is : ", reflect.TypeOf(i).Name())
	switch i.(type) {
	case string:
		fmt.Printf("%s: is a string\n", i)
	case int:
		fmt.Printf("%d: is a int\n", i)
	}

}

type shape interface {

	area() float64
	perimeter() float64
}

type rect struct {

	width, height float64
}

func (r *rect) area() float64 {

	return r.width * r.height
}

func (r *rect) perimeter() float64 {

	return 2 * (r.width + r.height)
}

type circle struct {

	radius float64
}

func (c *circle) area() float64 {

	return 2 * math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {

	return 2 * math.Pi * c.radius
}

func main() {

	str := "abc"
	d := 10
	t(str)
	t(d)
	fmt.Println("type: ", reflect.TypeOf(str))

	r := rect {width: 2, height: 4}
	c := circle {radius: 2}
	s := []shape{&r, &c}

	fmt.Println("===========test interface=========")
	for _, sh := range s {

		fmt.Println("area: ", sh.area())
		fmt.Println("perimeter: ", sh.perimeter())
	}
}
