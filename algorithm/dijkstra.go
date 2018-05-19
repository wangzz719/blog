package main

import "fmt"

func dijkstra(graph [][]int, target int) int {
	result := make(map[int]int)
	result[0] = 0
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph); j++ {
			if graph[i][j] != -1 {
				if result[j] != 0 {
					if result[j] > result[i]+graph[i][j] {
						result[j] = result[i] + graph[i][j]

					} else {
						result[j] = result[i] + graph[i][j]
					}
				} else {
					result[j] = result[i] + graph[i][j]
				}
			}
		}
	}
	return result[target]
}

func main() {
	graph := [][]int{{-1, 1, 12}, {-1, -1, 9}, {-1, -1, -1}}
	fmt.Println(dijkstra(graph, 2))
}
