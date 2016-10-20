package main

import (. "fmt")

type user struct {

	name string
	age int
}

func (u user) ToString() string {

	return Sprintf("%+v", u)
}

type manager struct {

	user
	title string
}

func main() {

	var m manager
	m.name = "Tom"
	m.age = 30
	//m.title = "pm"

	Println(m.ToString())
}
