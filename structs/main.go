package main

import "fmt"

type Ninja struct {
  Name string
  Power int
}

func main() {
  sparkachu := Ninja{"Sparkachu", 1000}

  fmt.Println(sparkachu.Name)
  fmt.Println(sparkachu.Power)
  sparkachu.boost()
  fmt.Println(sparkachu.Power)
}

func (n *Ninja) boost()  {
  n.Power += 1000
}
