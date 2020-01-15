package MinCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	var minCost, p1cost, p2cost int

	for i := 0; i < len(cost); i++ {
		p2cost = p1cost
		p1cost = minCost
		if p1cost+cost[i] <= p2cost+cost[i] {
			minCost = p1cost + cost[i]
		} else {
			minCost = p2cost + cost[i]
		}
	}
	if p1cost < minCost {
		minCost = p1cost
	}
	return minCost
}
