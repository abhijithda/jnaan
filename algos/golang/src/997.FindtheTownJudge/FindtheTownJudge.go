package findthetownjudge

import "log"

func findJudge(N int, trust [][]int) int {
	log.Printf("\n\n\nGiven N: %d, trust: %v", N, trust)
	if N == 1 {
		if len(trust) == 0 {
			return 1
		}
		if len(trust[0]) == 0 {
			return 1
		}
	}

	judge := -1

	type trusted struct {
		by []int
		// reverse trusted by count!
		rcount int
	}
	p := map[int]*trusted{}
	for i := range trust {
		per := trust[i][0]
		t := trust[i][1]
		if _, ok := p[t]; !ok {
			p[t] = &trusted{}
		}
		p[t].by = append(p[t].by, per)
		if _, ok := p[per]; !ok {
			p[per] = &trusted{}
		}
		p[per].rcount++
	}
	log.Printf("Trusted: %+v", p)

	for k, v := range p {
		log.Printf("%d trusted by %d: %v", k, v.rcount, v.by)
		if len(v.by) == N-1 {
			if v.rcount == 0 {
				if judge == -1 {
					judge = k
				} else {
					// As there can't be two judges return -1
					return -1
				}
			}
		}
	}
	return judge
}
