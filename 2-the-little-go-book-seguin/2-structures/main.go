package main

import "fmt"

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

// Super - a method of Saiyan struct
func (s *Saiyan) Super() {
	s.Power += 10000
}

// Super - function that accepts a pointer to an object of type Saiyan
func Super(s *Saiyan) {
	s.Power += 10000
}

// NewSaiyan - factory (or constructor-like) that creates and returns an instance of Saiyan struct
func NewSaiyan(name string, power int, father *Saiyan) Saiyan {
	return Saiyan{
		Name:   name,
		Power:  power,
		Father: father,
	}
}

func main() {
	goku := &Saiyan{"Goku", 9000, nil}
	Super(goku)
	fmt.Println(goku.Power)
	goku.Super()
	fmt.Println(goku.Power)
	gohan := NewSaiyan("Gohan", 1000, goku)
	fmt.Println(goku, gohan, gohan.Father)
	// these both are equal:
	// tofu := new(Saiyan)
	// tofu := &Saiyan{}

}
