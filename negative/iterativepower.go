package piscine

func IterativePower(num, exp int) int {
	if exp < 0 {
		return 0
	}
	var total int = 1
	for i := 0; i < exp; i++ {
		total *= num
	}
	return total
}
