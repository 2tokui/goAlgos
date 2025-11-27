package algo

// i have no clue why this worked
// i just tried to fix bugs and it worked
//func InsertionSort(array []int) {
//	for idx := 1; idx < len(array); idx++ {
//		tmp := array[idx]
//		for j := idx-1; j >= 0; j-- {
//			if array[j] > tmp {
//				array[j+1] = array[j]
//			} else if array[j] < tmp {
//				array[j+1] = tmp
//				break
//			} else if array[j] == tmp {
//				array[j+1] = tmp
//			}
//			if j == 0 {
//				array[0] = tmp
//			}
//		}
//	}
//}

// The smart way that I couldn't have thought about myself.
func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		tmp := array[i]
		j := i - 1
		for j >= 0 && array[j] > tmp {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = tmp
	}
}

