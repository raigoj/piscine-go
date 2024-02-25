package piscine

func AppendRange(min, max int) []int {
	var slc []int
	for n := min; n < max; n++ {
		slc = append(slc, n)
	}
	return slc
}
