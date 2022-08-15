package math

import "testing"

type testpair struct {
	values  []float64
	average float64
	min     float64
	max     float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 1, 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 1, 1},
	{[]float64{-1, 1}, 0, -1, 1},
	{[]float64{}, 0, 0, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		result := Average(pair.values)
		if result != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", result,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range tests {
		result := Min(pair.values)
		if result != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", result,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range tests {
		result := Max(pair.values)
		if result != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", result,
			)
		}
	}
}
