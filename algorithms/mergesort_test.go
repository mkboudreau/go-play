package algorithms

import (
	"testing"
)

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
