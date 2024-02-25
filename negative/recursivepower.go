package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	var total int = 1
	if power > 0 {
		total = nb * RecursivePower(nb, power-1)
	}
	return total
}
