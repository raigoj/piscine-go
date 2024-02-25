package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	ans := 0
	for {
		pow := ans * ans
		if pow == nb {
			return ans
		}
		if pow > nb {
			return 0
		}
		ans++
	}
}
