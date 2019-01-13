package kclosestpointstoorigin

import (
	"log"
	"math"
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	ans := [][]int{}
	// minD := int(^uint(0) >> 1)
	dist := []int{}
	distInd := map[int][]int{}
	for p := range points {
		// Optimization: No need to calculate sqrt as all values would be
		// compared without sqrt!
		val := math.Pow(float64(points[p][0]), float64(2)) +
			math.Pow(float64(points[p][1]), float64(2))

		log.Println("Distance:", val)
		dist = append(dist, int(val))
		distInd[int(val)] = append(distInd[int(val)], p)
	}

	sort.Ints(dist)
	log.Println("After sorting by distance:", dist)
	for d := range dist {
		if K == 0 {
			break
		}
		for _, indexes := range distInd[dist[d]] {
			ans = append(ans, points[indexes])
			log.Println("Added to ans:", ans)
			K--
			if K == 0 {
				break
			}
		}

	}
	return ans
}
