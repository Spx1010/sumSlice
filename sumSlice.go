package sumslice

func Sum(x ...int) int {
	sum := 0
	for i := range x {
		sum += x[i]
	}
	return sum
}
