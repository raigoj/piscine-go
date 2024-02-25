package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var mult int
	var i int
	for ; i < len(a)-1; i++ {
		comp := f(a[i], a[i+1])
		if comp > 0 {
			mult = -1
			break
		}
		if comp < 0 {
			mult = 1
			break
		}
	}
	for ; i < len(a)-1; i++ {
		if f(a[i]*mult, a[i+1]*mult) > 0 {
			return false
		}
	}
	return true
}
