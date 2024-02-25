package piscine

func Join(elems []string, sep string) string {
	var notFirst bool
	var res string
	for _, v := range elems {
		if notFirst {
			res += sep
		} else {
			notFirst = true
		}
		res += v
	}
	return res
}
