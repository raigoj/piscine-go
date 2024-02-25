package piscine

func MakeRange(min, max int) []int {
	len := max - min
	if len <= 0 {
		var slc []int
		return slc
	}
	slc := make([]int, len)
	for i := 0; i < len; i++ {
		slc[i] = min + i
	}
	return slc
}
