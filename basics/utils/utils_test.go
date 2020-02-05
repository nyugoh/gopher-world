package utils

import (
	"testing"
)

type testPair struct {
	values []float64
	average float64
}

var pairs = []testPair {
	{ []float64{1, 2}, 1.5},
	{ []float64{1, 1, 1, 1}, 1},
	{ []float64{}, 0},
	{ []float64{-12.5, -10}, -11.25},
}

func TestAverage(t *testing.T) {
	for _, pair := range pairs {
		avg := Average(pair.values)
		if pair.average != avg {
			t.Error(
				"For ", pair.values,
				" expected", pair.average,
				"but got", avg,
			)
		}
	}
}

func TestCurrentDate(t *testing.T) {
	return
}