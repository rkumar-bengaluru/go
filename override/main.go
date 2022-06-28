package main

import "fmt"

type Animal struct {
	color string
}

type Dog struct {
	Animal
}

// overridden behavior
func (d *Dog) doSomething() {
	// field promotion.
	fmt.Println("Dog Color is ....", d.color)
}

// default behaviour
func (a *Animal) doSomething() {
	fmt.Println("color of animal ", a.color)
}

func main() {
	var a Animal = Animal{color: "red"}
	d := Dog{a}
	// default behavior
	d.Animal.doSomething()
	// overridden method behavior
	d.doSomething()
}
