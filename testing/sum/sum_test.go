package sum_test

import (
	"github.com/nyugoh/gopher-world/testing/sum"
	"testing"
)

func TestSum(t *testing.T) {
	// For simple testing
	/*s := Sum(1, 2, 3, 4, 5)
	if s != 15 {
		t.Fatalf("Expected 15 got %v", s)
	}
	*/

	// Table testing
	tt := []struct{
		name string
		numbers []int
		sum int
	}{
		{"1-5", []int{1, 2, 3, 4, 5}, 15},
		{"1 & -1", []int{1, -1}, 0},
		{"nil slice", []int{}, 0},
		{"one", []int{1}, 1},
	}
	for _ , tc := range tt {
		// Run inidividual test cases
		t.Run(tc.name, func(t *testing.T) {
			if s := sum.Sum(tc.numbers...); s != tc.sum {
				t.Errorf("%s expected %v got %v", tc.name, tc.sum, s)
			}
		})
	}
}
