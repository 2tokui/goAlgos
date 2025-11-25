package algo

func SelectionSort(array []int) {
	for k := range array {
		minIdx := k

		for idx := k; idx < len(array); idx++ {
			if array[idx] < array[minIdx] {
				minIdx = idx
			}
		}

		tmp := array[k]
		array[k] = array[minIdx]
		array[minIdx] = tmp
	}
}
