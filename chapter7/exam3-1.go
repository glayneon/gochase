package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Android struct {
	Person
	Model   string
	Company string
}

func (p *Person) Talk() {
	fmt.Printf("Hi, I'm %s %s\n", p.firstName, p.lastName)
}

func main() {
	a := Android{Model: "R2", Company: "Invest"}
	a.firstName, a.lastName = "Ruri", "Web"
	a.Talk()
	a.Company = "Bolt"
	fmt.Println(a.Company)
}
