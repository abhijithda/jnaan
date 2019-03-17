package pairsofsongswithtotaldurationsdivisibleby60

import "internal/mathutils"

func numPairsDivisibleBy60(time []int) int {
	res := 0
	unq := map[int]int{}
	visited := map[int]bool{}
	for _, t := range time {
		r := t % 60
		if _, ok := unq[r]; !ok {
			unq[r] = 0
			visited[r] = false
		}
		unq[r]++
	}

	for k, v := range unq {
		if visited[k] {
			continue
		}
		if k == 0 || k == 30 {
			visited[k] = true
			res += mathutils.NCR(v, 2)
		} else {
			visited[k] = true
			r := 60 - k
			if v1, ok := unq[r]; ok {
				visited[r] = true
				res += mathutils.NCR(v, 1) * mathutils.NCR(v1, 1)
			}
		}
	}
	return res
}
