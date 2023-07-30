package selection

// Find smallest
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIdx := 0

	for idx := range arr {
		if arr[idx] < smallest {
			smallest = arr[idx]
			smallestIdx = idx
		}
	}

	return smallestIdx
}

func SelectionSort(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)

	for idx := range arr {
		smallestIdx := findSmallest(arr)
		newArr[idx] = arr[smallestIdx]
		arr = append(arr[:smallestIdx], arr[smallestIdx + 1:]...)
	}

	return newArr
}
