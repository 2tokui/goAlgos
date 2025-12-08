package algo

// speed:
// avg and best case = O(n log(n))
// worst case (already sorted) = O(n^2)
// space: O(log(n)) since it uses recursion

// after seeing other implementations
// i can see how shitty this is.
// i'm not sorting in-place and i'm creating
// copies recursively wasting precious memory
func QuickSort(sourceArray []int) []int {
	if len(sourceArray) <= 1 {
		return sourceArray
	}

	array := make([]int, len(sourceArray))
	copy(array, sourceArray)

	var pivotIndex int
	pivot := array[len(array)-1]
	j := 0
	i := -1
	var tmp int

	for range array {
		if array[j] < pivot {
			i++
			tmp = array[i]
			array[i] = array[j]
			array[j] = tmp
			j++
		} else {
			j++
		}
		if j == len(array)-1 {
			tmp = array[i+1]
			array[i+1] = pivot
			array[len(array)-1] = tmp
			pivotIndex = i + 1
		}
	}

	leftSide := array[:pivotIndex]
	rightSide := array[pivotIndex+1:]

	left := QuickSort(leftSide)
	right := QuickSort(rightSide)

	return append(append(left, pivot), right...)
}
