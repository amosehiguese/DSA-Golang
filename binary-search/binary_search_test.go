package binary_search

import "testing"

func TestBinarySearchStr(t *testing.T){
	listStr := []string{"joy", "my", "example", "more", "as", "in"}
	got := BinarySearch(listStr, "more")
	expect := 3

	if got != expect {
		t.Errorf("Got %v but expected %v",got, expect)
	}
}
func TestBinarySearchInt(t *testing.T){
	listInt := []int{1, 2, 3, 4, 5}
	got := BinarySearch(listInt, 4)
	expect := 3

	if got != expect {
		t.Errorf("Got %v but expected %v",got, expect)
	}
}
