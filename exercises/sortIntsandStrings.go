package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{8, 4, 6, 2, 7, 3, 9, 1}
	names := []string{"a", "B", "A", "z", "b"}

	fmt.Println(numbers)
	fmt.Println(names)

	sort.Ints(numbers)
	fmt.Println(numbers)

	sort.Strings(names)
	fmt.Println(names)

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers)

	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Println(names)
}
