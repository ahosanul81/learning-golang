package main

import "fmt"

type Student struct {
	id      int
	address string
	age     int
}

func displayInfo(s Student) {
	fmt.Println(s.id)
	fmt.Println(s.address)
	fmt.Println(s.age)
}

func (x *Student) increaseAgeal(val int) {
	x.age = x.age + val
}
func structure() {
	rahim := Student{101, "Cumilla", 20}
	karim := Student{102, "Jashore", 24}
	rahim.increaseAgeal(2)
	displayInfo(rahim)
	displayInfo(karim)

}
