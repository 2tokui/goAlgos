package algo

// I had difficulties even when copying from the guy writing
// in Java
func merge(leftArray, rightArray, array []int) {
	leftSize := len(leftArray)
	rightSize := len(rightArray)
	var i, l, r int
	for l < leftSize && r < rightSize {
		if leftArray[l] < rightArray[r] {
			array[i] = leftArray[l]
			i++
			l++
		} else {
			array[i] = rightArray[r]
			i++
			r++
		}
	}
	for l < leftSize {
		array[i] = leftArray[l]
		i++
		l++
	}
	for r < rightSize {
		array[i] = rightArray[r]
		i++
		r++
	}
}

func MergeSort(array []int) {
	length := len(array)

	// base case
	if length <= 1 { return }

	middle := length/2

	left := []int{}
	right := []int{}

	for i := range length {
		if i < middle {
			left = append(left, array[i])
		} else {
			right = append(right, array[i])
		}
	}

	MergeSort(left)
	MergeSort(right)

	merge(left, right, array)
}

// I tried for a long time, but I couldn't get my brain to
// find a solution
//func merge(arrayA, arrayB []int) []int {
//	newArray := []int{}
//	size := len(arrayA)+len(arrayB)
//	for range size {
//		for _, v := range arrayA {
//			for _, v2 := range arrayB {
//				if v < v2 {
//					newArray = append(newArray, v)
//					break
//				} else {
//					newArray = append(newArray, v2)
//					break
//				}
//			}
//		}
//	}
//	if arrayA[0] > arrayB[0] {
//		return []int{arrayB[0], arrayA[0]}
//	}
//	return []int{arrayA[0], arrayB[0]}
//}
//
//func MergeSort(array []int) []int {
//	if len(array) == 1 {
//		return array
//	}
//	var middle int = len(array) / 2
//	left := MergeSort(array[:middle])
//	right := MergeSort(array[middle:])
//	return(merge(left, right))
//}
