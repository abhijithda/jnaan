package nrepeatedelementinsize2narray

func repeatedNTimes(A []int) int {
	ele := -1
	dup := map[int]bool{}
	for _, a := range A {
		if _, ok := dup[a]; ok {
			ele = a
			break
		} else {
			dup[a] = true
		}
	}
	return ele
}
