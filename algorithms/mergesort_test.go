package algorithms

import (
	"testing"
)

var testCases = []struct {
	test     []int
	expected []int
}{
	{[]int{}, []int{}},
	{[]int{2}, []int{2}},
	{[]int{5, 7}, []int{5, 7}},
	{[]int{7, 1, 4}, []int{1, 4, 7}},
	{[]int{2, 1, 6, 8}, []int{1, 2, 6, 8}},
	{[]int{2, 1, 6, 4, 8}, []int{1, 2, 4, 6, 8}},
	{[]int{5, 7, 1, 4, 57, 2}, []int{1, 2, 4, 5, 7, 57}},
	{[]int{5, 7, 1, 4, 0, 57, 2}, []int{0, 1, 2, 4, 5, 7, 57}},
	{[]int{5, 7, 1, 1, 1, 4, 57, 2}, []int{1, 1, 1, 2, 4, 5, 7, 57}},
	{[]int{5, 7, 1, 102, 1, 1, 4, 57, 2}, []int{1, 1, 1, 2, 4, 5, 7, 57, 102}},
	{[]int{5, 7, 1, 1, 200, 105, 1, 4, 57, 2}, []int{1, 1, 1, 2, 4, 5, 7, 57, 105, 200}},
}

func TestMergeSort(t *testing.T) {
	for _, testCase := range testCases {
		actual := MergeSort(testCase.test)
		if IsNotEqual(actual, testCase.expected) {
			t.Errorf("actual %v does not equal expected %v", actual, testCase.expected)
		}
	}
}

func TestChannelMergeSort(t *testing.T) {
	for _, testCase := range testCases {
		actual := ChannelMergeSort(testCase.test)
		if IsNotEqual(actual, testCase.expected) {
			t.Errorf("actual %v does not equal expected %v", actual, testCase.expected)
		}
	}
}


func IsNotEqual(a, b []int) bool {
	return !IsEqual(a, b)
}
func IsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
