package math

// Find the average value of an array of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Find the minimum element of an array of numbers
func Min(xs []float64) float64 {
	min := xs[0]
	for _, x := range xs {
		if min > x {
			min = x
		}
	}
	return min
}

// Find the maximum of an array of numbers
func Max(xs []float64) float64 {
	max := xs[0]
	for _, x := range xs {
		if max < x {
			max = x
		}
	}
	return max
}
