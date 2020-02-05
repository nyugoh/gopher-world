package main

import "fmt"

func main() {
  scores := []int{1, 2, 4, 6, 8, 3}

  // Slicing
  // [StartIndex:EndIndex] ~ EndIndex is not included
  fmt.Println(scores[2:5]) // returns 4, 6, 8
  fmt.Println(scores[:4]) // returns index 0-3 i.e. 1,2,4,6
  fmt.Println(scores[3:]) // returns index 3 to the end

  // Deleting element at a certain index
  
}
