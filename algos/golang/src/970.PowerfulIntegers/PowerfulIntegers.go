package powerfulintegers

import "math"

func powerfulIntegers(x int, y int, bound int) []int {
	res := map[int]bool{}

	val := 0
	for i := 0; i <= int(math.Pow(float64(x), float64(i))); i++ {
		for j := 0; j <= int(math.Pow(float64(y), float64(j))); j++ {
			val = int(math.Pow(float64(x), float64(i)) + math.Pow(float64(y), float64(j)))
			if val > bound {
				break
			}
			res[val] = true
		}
	}

	powints := []int{}
	for k := range res {
		powints = append(powints, k)
	}
	return powints
}
