package distringmatch

func diStringMatch(S string) []int {
	d := len(S)
	i := 0

	A := make([]int, d+1)
	for s := range S {
		if S[s] == 'I' {
			A[s] = i
			i++
		} else if S[s] == 'D' {
			A[s] = d
			d--
		}
	}
	A[len(S)] = i
	return A
}
