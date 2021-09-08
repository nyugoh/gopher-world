package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)

	for i := 0; i < 100; i++ {
		scores = append(scores, int(rand.Int31n(1000)))
	}

	// sort
	sort.Ints(scores)
	fmt.Println(scores)

	worst := make([]int, 10)
	best := make([]int, 10)
	copy(worst, scores[:10])            // least 10
	copy(best, scores[len(scores)-10:]) // Top 10

	fmt.Println(worst)
	fmt.Println(best)
}
