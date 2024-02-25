package piscine

func ForEach(f func(int), a []int) {
	for _, s := range a {
		f(s)
	}
}
