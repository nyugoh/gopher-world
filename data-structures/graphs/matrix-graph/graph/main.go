package graph

import "fmt"

type Graph struct {
	V int
	E int
	Adj [][]int
}

func (g Graph) CreateAdjMatrix() {
	fmt.Print("Enter no. of vertices::")
	_, err := fmt.Scanf("%d", &g.V)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Vertices::", g.V)

	fmt.Print("Enter no. of edges::")
	_, err = fmt.Scanf("%d", &g.E)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Edges::", g.E)

	g.Adj = make([][]int, g.V, g.V)

	for i:=0; i < g.V; i++ {
		// Declare the other inner array
		g.Adj[i] = make([]int, g.V, g.V)
		for j:=0; j < g.V; j++ {
			g.Adj[i][j] = 0
		}
	}

	// Read the edjes
	u, v := 0, 0
	for j:=0; j < g.E; j++ {
		fmt.Println("Reading Edge::")
		_, err = fmt.Scanf("%d %d", &u, &v)
		g.Adj[u][v] = 1
		g.Adj[v][u] = 1 // if directed set both to 1
	}
	fmt.Println(g.Adj)
}
