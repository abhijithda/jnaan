package ClimbingStairs

func climbStairs(n int) int {
	var pn1Cost, pn2Cost, curCost int
	if n > 1 {
		pn1Cost = 1
	}

	for i := n; i > 0; i-- {
		curCost = pn1Cost + pn2Cost
		pn2Cost = pn1Cost
		pn1Cost = curCost
	}
	return curCost
}

func climbStairsRecursion(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}
