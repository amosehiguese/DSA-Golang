package queue

import (
	"reflect"
	"testing"
)

func TestQueueInt(t *testing.T) {
	data := []struct{
		arr []int
		expected []int
	}{
		{
			arr:[]int{1, 2, 3, 4, 5},
			expected:[]int{3, 4, 5,7,8,9},
		},
		{
			arr: []int{2,3,4,7,5,1,6},
			expected: []int{4,7,5,1,6,7,8,9},
		},
		{
			arr: []int{2,7,7,4,7},
			expected: []int{7,4,7,7,8,9},
		},

	}

	arr := []int{7,8,9}
	queue := Queue[int]{}
	for _, d := range data {
		queue.vals = d.arr
		queue.Dequeue()
		queue.Dequeue()
		queue.Enqueue(arr...)
		expected := d.expected

		if !reflect.DeepEqual(queue.vals, expected) {
			t.Errorf("Got %v but expected %v", queue.vals, expected)
		}
	}
}

func TestQueueString(t *testing.T) {
	data := []struct{
		arr []string
		expected []string
	}{
		{
			arr:[]string{"my", "in", "who", "what", "now"},
			expected:[]string{ "who", "what", "now", "signed", "amos"},
		},
		{
			arr: []string{"ant", "bird", "cat", "dog"},
			expected: []string{"cat", "dog", "signed", "amos"},
		},
		{
			arr: []string{"react", "tailwind", "css", "kubernetes", "golang", "typescript"},
			expected: []string{"css", "kubernetes", "golang", "typescript", "signed", "amos"},
		},

	}

	arr := []string{"signed", "amos"}
	queue := Queue[string]{}
	for _, d := range data {
		queue.vals = d.arr
		queue.Dequeue()
		queue.Dequeue()
		queue.Enqueue(arr...)
		expected := d.expected

		if !reflect.DeepEqual(queue.vals, expected) {
			t.Errorf("Got %v but expected %v", queue.vals, expected)
		}
	}
}

