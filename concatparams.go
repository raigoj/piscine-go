package piscine

func ConcatParams(args []string) string {
	var str string
	var notFirst bool
	for _, v := range args {
		if notFirst {
			str += "\n"
		} else {
			notFirst = true
		}
		str += v
	}
	return str
}
