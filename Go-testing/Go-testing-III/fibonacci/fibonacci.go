package fibonacci

func fibonacci(iterations int) int {
	x1 := 0
	x2 := 1
	result := 0

	for i := 0; i < iterations-1; i++ {
		result = x1 + x2
		x1 = x2
		x2 = result
	}

	return result
}
