package sum_test

import (
	"fmt"
	"github.com/nyugoh/gopher-world/testing/sum"
)

func ExampleSum() {
	s := sum.Sum(1, 2, 3, 4, 5)
	fmt.Println("sum of one to five is", s)
	// Output:
	// sum of one to five is 15
}
