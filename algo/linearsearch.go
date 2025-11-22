package algo

func LinearSearch(array []int, v int) int {
	for k, i := range array {
		if i == v {
			return k
		}
	}
	return -1
}

