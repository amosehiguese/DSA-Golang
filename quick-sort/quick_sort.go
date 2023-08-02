package quick

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int
	for _, item := range arr[1:] {
		if item < pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}

	less = append(QuickSort(less), pivot)
	greater = QuickSort(greater)

	sortedArr := append(less, greater...)
	return sortedArr
}
