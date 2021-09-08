package main

import "fmt"

type Ninja struct {
	Name  string
	Power int
	Level float64
}

func totalPower(ninjas []*Ninja) []int {
	powers := make([]int, len(ninjas))

	for index, ninja := range ninjas {
		powers[index] = ninja.Power
	}
	return powers
}
func main() {
	spark := Ninja{"Spark", 9000, 5.56}
	crader := Ninja{"Crader", 10000, 4.56}
	slimmer := Ninja{"Slimmer", 7000, 7.924}
	ninjas := []*Ninja{&spark, &crader, &slimmer}
	fmt.Println(totalPower(ninjas))
}
