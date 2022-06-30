package main

import "fmt"

type AnimalInterface interface {
	say()
}

type Animal struct {
	color string
}

func (a *Animal) say() {
	fmt.Println("Hi from say ")
}

func (a *Dog) say() {
	fmt.Println("Hi from Dog ")
}

type Dog struct {
	Animal // embedding
	breed  string
}

func check(b AnimalInterface) {
	b.say()
}

func main() {
	a := Animal{color: "black"}
	d := &Dog{a, "Bulldog"}
	d.say()
	fmt.Println("color of dog is ", d.breed)
	check(d)
}
