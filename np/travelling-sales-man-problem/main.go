package main

import (
	"fmt"
	"math/rand"
)

func tsp(cities []int, paths [][]int, dist int) bool {
	found := false
	routes := permutations(cities)
	for _, route := range routes {
		totalDistance := 0
		/*for j:=0;j< len(route);j++ {
			if j+1 < len(route) {
				p := route[j]
				k := route[j+1]
				totalDistance += paths[p][k]
			}
		}*/
		for i := range route {
			if i != 0 {
				p := route[i-1]
				k := route[i]
				totalDistance += paths[p][k]
			}
		}
		if totalDistance < dist {
			found = true
		}
	}
	return found
}

// Don't edit below this line

func main() {
	rand.Seed(0)
	for numCities := 2; numCities < 10; numCities++ {
		paths := [][]int{}
		cities := []int{}
		for i := 0; i < numCities; i++ {
			path := []int{}
			for j := 0; j < numCities; j++ {
				if i == j {
					path = append(path, 0)
				} else if j < i {
					path = append(path, paths[j][i])
				} else {
					path = append(path, rand.Intn(1000))
				}
			}
			paths = append(paths, path)
			cities = append(cities, i)
		}
		dist := rand.Intn(3000)
		pathExists := tsp(cities, paths, dist)
		fmt.Println("Paths:")
		printMatrix(paths)
		fmt.Printf("Path shorter than %v exists: %v\n", dist, pathExists)
		fmt.Println("---")
	}
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func printMatrix(s2 [][]int) {
	n := len(s2)

	// number of columns in s2
	m := len(s2[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(s2[i][j], " ")
		}
		fmt.Print("\n")
	}
}
