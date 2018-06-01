package main

import "fmt"

type Ninja struct {
  Name string
  Power int
  Friends map[int]string
}

func main() {
  team := make( map[string]int)

  team["spark"] = 9
  team["hulk"] = 5
  team["superman"] = 20

  fmt.Println(team) // map[spark:9 hulk:5 superman:20]

  name, exists := team["batman"]
  fmt.Println(name)
  fmt.Println(exists)

  delete(team, "batman")
  fmt.Println(team)

  hulk := &Ninja{
    Name: "Hulk",
    Power: 2500,
    Friends: make(map[int]string),
  }

  hulk.Friends[0] = "Spark"
  hulk.Friends[1] = "Batman"

  fmt.Println(*hulk)
}
