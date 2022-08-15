package main

import "fmt"
import m "1-golang-book-caleb/9-packages/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
	fmt.Println(m.Min([]float64{1, 2, 3, 4}))
	fmt.Println(m.Max([]float64{1, 2, 3, 4}))
}
