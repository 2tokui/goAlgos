package algo

import "math"

// finds position of target
// collection must be _sorted_
// divide and conquer

// O(1) for 1st and last elements
// O(log n) for elements in the middle
func BinarySearch(collection []int, targetElement int) int {
	var start int = 0
	var end int = len(collection)-1

	if targetElement == collection[start] {
		return start
	}

	if targetElement == collection[end] {
		return end
	}

	for {
		middle := int(math.Floor(float64((start+end)/2)))
		middleGuy := collection[middle]
		if middleGuy == targetElement {
			return middle
		}
		if targetElement < middleGuy {
			end = middle
		} else {
			start = middle
			if middle == end -1 {
				return -1 // not found
			}
		}
	}
}

