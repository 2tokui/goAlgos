package algo

import "fmt"

// avg case: O(log(log(n))) uniformly distributed data
// 1, 2, 4, 8, 16, 32, 64...
// worst case: O(n)
// this algo is used in ordered arrays
func InterpolationSearch(array []int, value int) int {
	high := len(array) - 1
	low := 0

	for value >= array[low] && value <= array[high] && low <= high {
		probe := low + (high - low) * (value - array[low]) / (array[high] - array[low])
		fmt.Println("Probe:", probe)
		if array[probe] == value {
			return probe
		} else if array[probe] < value {
			low = probe + 1
		} else {
			high = probe - 1
		}
	}

	return -1
}
 
