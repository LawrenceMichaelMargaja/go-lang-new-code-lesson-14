package utils

// in:  3,2,1
// out: 1,2,3
func BubbleSort(elements []int) []int {
	keepRunning := true

	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements) - 1; i++ {
			if elements[i] > elements[i + 1] {
				elements[i], elements[i + 1] = elements[i + 1], elements[i]
				keepRunning = true
			}
		}
	}

	return elements
}
