package main

import "fmt"

func main() {
  var scores[5] float64
  var total float64 = 0
  scores[0] = 70
  scores[1] = 60
  scores[2] = 80
  scores[3] = 90
  scores[4] = 100

  fmt.Println(scores)
  // len() function returns the length of the array
  for i := 0; i < len(scores); i++ {
    total += scores[i]
  }

  fmt.Println("Total ::", total)
  fmt.Println("Average ::", total/float64(len(scores)))
  // float64( other type here) converts the type to the new format we want
}
