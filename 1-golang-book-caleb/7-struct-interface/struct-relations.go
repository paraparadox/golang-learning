package main

import "fmt"

// Person struct
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// -----

// AndroidHasPerson struct
type AndroidHasPerson struct {
	Person Person
	Model  string
}

// -----

// AndroidIsPerson struct
type AndroidIsPerson struct {
	Person // inherits Person struct fields and methods
	Model  string
	Name   string
}

func (a *AndroidIsPerson) SelfTalk() {
	fmt.Println("Yo! My name is", a.Name)
}

// -----

func main() {
	person := Person{"Bob"}
	person.Talk()

	androidHasPerson := AndroidHasPerson{Person: Person{"Steve"}, Model: "XH-12"}
	androidHasPerson.Person.Talk()

	androidIsPerson := AndroidIsPerson{Model: "XH-13", Name: "Alan"}
	androidIsPerson.Talk()               // prints android's person-name which is empty yet
	androidIsPerson.SelfTalk()           // prints android's name which is Alan in initialization
	androidIsPerson.Person.Name = "Mark" // sets android's person-name to Mark

	// ways to call Talk() method of androidIsPerson:
	androidIsPerson.Talk()        // prints android's person-name which is Mark
	androidIsPerson.Person.Talk() // prints android's person-name which is Mark
}
