package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if f(a[j], a[j+1]) == 1 {
				x := a[j]
				a[j] = a[j+1]
				a[j+1] = x
			}
		}
	}
}
