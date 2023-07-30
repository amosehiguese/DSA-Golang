package binary_search

type IntStr interface {
	int | string
}

func BinarySearch[T IntStr](list []T, item T) int {
	low := 0
	high := len(list)

	var mid int
	for low <= high {
		mid = (high + low) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return mid
}

