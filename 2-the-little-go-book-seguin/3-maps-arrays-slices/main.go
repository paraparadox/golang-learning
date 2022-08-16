package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func extractPowers(saiyans []*Saiyan) []int {
	powers := make([]int, len(saiyans))
	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}

func removeAtIndex(source []*Saiyan, index int) []*Saiyan {
	lastIndex := len(source) - 1
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

func main() {
	// Checking capacity updating
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		// if our capacity changed,
		// Go had to grow our array to accommodate the new data
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}

	// Creating slice with pre-defined capacity
	saiyans := []*Saiyan{
		{
			Name:   "Tofu",
			Power:  1000,
			Father: nil,
		},
		{
			Name:   "Goku",
			Power:  9000,
			Father: nil,
		},
		{
			Name:   "Gohu",
			Power:  10000,
			Father: nil,
		},
	}

	fmt.Println(extractPowers(saiyans))

	// Removing an element of a slice
	for _, saiyan := range saiyans {
		fmt.Println(saiyan)
	}
	fmt.Println()
	saiyans = removeAtIndex(saiyans, 0)
	for _, saiyan := range saiyans {
		fmt.Println(saiyan)
	}

	// Copying one slice into another
	scores = make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst[2:], scores[:10])
	fmt.Println(worst)
}
