package algo

// i did this first please don't kill me
//func BubbleSort(array []int) {
//	isOrdered := func() bool {
//		for k := range array {
//			if k == len(array) - 1 {
//				break
//			}
//			a := array[k]
//			b := array[k+1]
//			if a > b {
//				return false
//			}
//		}
//		return true
//	}
//	for k := 0; k < len(array); k++ {
//		if k == len(array) - 1 {
//			k = 0
//			if isOrdered() {
//				return
//			}
//		}
//		a := array[k]
//		b := array[k+1]
//		if a > b {
//			array[k] = b
//			array[k+1] = a
//		}
//	}
//}

func BubbleSort(array []int) {
	for range len(array) {
		for k := range array {
			if k == len(array) - 1 {
				break
			}
			if array[k] > array[k+1] {
				temp := array[k]
				array[k] = array[k+1]
				array[k+1] = temp
			}
		}
	}
}

