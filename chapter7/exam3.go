package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, My name is", p.Name)
}

func main() {
	p := Person{Name: "Chase"}
	a := Android{Name: "Suri", Model: "R2D2"}
	p.Talk()
	// a
	a.Person.Talk()
	a.Talk()
	fmt.Println(a.Name, a.Model)
}
