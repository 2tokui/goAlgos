package algo

import "math/rand/v2"

// we need something to make it stop, a base case

type Vector struct {
	X int
	Y int
}

// My shitty randomwalker
func RandomWalk(root Vector, steps int) []Vector {
	if steps == 0 {
		return nil
	}
	// -2 -1 0 1 2
	offsetX := rand.IntN(4) - 2
	offsetY := rand.IntN(4) - 2
	vec := Vector{X: root.X + offsetX, Y: root.Y + offsetY}
	return append(RandomWalk(vec, steps-1), vec)
}

func Factorial(n int) int {
	if n < 1 {
		return 1
	}
	return n * Factorial(n - 1)
}

