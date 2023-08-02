package quick

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []struct{
		arr []int
		expected []int
	}{
		{
			arr:[]int{1, 2, 3, 4, 5},
			expected:[]int{1, 2, 3, 4, 5},
		},
		{
			arr: []int{2,3,4,7,5,1,6},
			expected: []int{1,2,3,4,5,6,7},
		},
		{
			arr: []int{2,7,7,4,7},
			expected: []int{2,4,7,7,7},
		},

	}

	for _ , d := range data {
		sortedArr := QuickSort(d.arr)
		expected := d.expected

		if !reflect.DeepEqual(sortedArr, expected) {
			t.Errorf("Got %v but expected %v", sortedArr, expected)
		}
	}
}

