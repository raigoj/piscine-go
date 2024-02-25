package piscine

func SortIntegerTable(table []int) {
	for {
		ready := true
		for i := 0; i < len(table)-1; i++ {
			if table[i] > table[i+1] {
				ready = false
				buffer := table[i]
				table[i] = table[i+1]
				table[i+1] = buffer
			}
		}
		if ready {
			break
		}
	}
}
