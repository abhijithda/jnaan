package minimumheighttrees

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func init() {
	const logFile = "log.txt"
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()
	log.SetOutput(f)
}

func removeEdgeNode(nodes []int, node int) []int {
	tnodes := []int{}
	for i, n := range nodes {
		if n == node {
			if i != 0 {
				tnodes = nodes[:i]
			}
			if i+1 < len(nodes) {
				tnodes = append(tnodes, nodes[i+1:]...)
			}
			break
		}
	}

	return tnodes
}

func addEleIfNotPresent(A []int, ele int) []int {
	for _, a := range A {
		if a == ele {
			return A
		}
	}
	return append(A, ele)
}

func findMinHeightTrees(n int, edges [][]int) []int {
	log.Printf("\n\n\nGiven n: %d; edges(%d): %v", n, len(edges), edges)
	// root nodes
	root := []int{}

	if len(edges) == 0 || len(edges) == 1 {
		if n == 1 {
			root = []int{0}
		}
		if n == 2 {
			root = []int{0, 1}
		}
		return root
	}
	// count nodes
	g := map[int][]int{}
	for e := range edges {
		n1 := edges[e][0]
		n2 := edges[e][1]
		g[n1] = append(g[n1], n2)
		g[n2] = append(g[n2], n1)
	}
	log.Println("Graph:", g)

	for len(g) > 2 {
		lns := map[int]int{} // leaf nodes
		for n, e := range g {
			log.Println("Node:", n)
			if len(e) == 1 {
				lns[n] = e[0]
				delete(g, n)
				log.Printf("Removed leaf node: %d from graph = %v",
					n, g)
			}
		}
		// Remove the edges formed with this node.
		log.Println("Edges that need to be removed:", lns)
		for ln, e := range lns {
			g[e] = removeEdgeNode(g[e], ln)
			log.Printf("After removing %d from graph[%d] = %v",
				ln, e, g[e])
		}
	}

	for n := range g {
		root = append(root, n)
	}
	sort.Ints(root)
	return root
}
