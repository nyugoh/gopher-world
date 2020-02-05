package main

import "fmt"

func main() {
  // Making slices
  scores := make([]int, 10) // this will make a slice of length 10 and capacity 10
  marks := make([]int, 5, 10) // will make a slice of length 5 and capacity of 10
  goals := make([]int, 0, 5)

  // Initialization
  names := []string{"Jane", "Doe", "John", "Steve"} //length is 4 and capacity is 4
  checks := make([]bool, 10) // All are set to false
  fmt.Println(len(names))
  fmt.Println(cap(names))
  fmt.Println(checks)

  fmt.Println(names)
  names = append(names, "Joe", "Hugo")
  fmt.Println(names)

  fmt.Printf("Scores has a length of %d and capacity of %d \n", len(scores), cap(scores))
  fmt.Printf("Marks has a length of %d and capacity of %d \n", len(marks), cap(marks))

  // marks[7] = 50 // will crush the program coz the length is 5
  // scores[2] = 50 // will work coz the array has a length of 5

  // append will increase the slice and return the new it
  goals = append(goals, 5)
  fmt.Println(goals)
  // add another item(s) to slice
  goals = append(goals, 2, 3, 6, 7) // append more than one item
  fmt.Println(goals)
}
