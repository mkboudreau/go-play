package algorithms

import (
	"math/rand"
	"testing"
)

func BenchmarkMergeSortWithNoSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[0].test)
	}
}
func BenchmarkMergeSortWithSizeOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[1].test)
	}
}
func BenchmarkMergeSortWithSizeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[2].test)
	}
}
func BenchmarkMergeSortWithSizeThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[3].test)
	}
}
func BenchmarkMergeSortWithSizeFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[4].test)
	}
}
func BenchmarkMergeSortWithSizeFive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[5].test)
	}
}
func BenchmarkMergeSortWithSizeSix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[6].test)
	}
}
func BenchmarkMergeSortWithSizeSeven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[7].test)
	}
}
func BenchmarkMergeSortWithSizeEight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[8].test)
	}
}
func BenchmarkMergeSortWithSizeNine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[9].test)
	}
}
func BenchmarkMergeSortWithSizeTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testCases[10].test)
	}
}

func BenchmarkMergeSortWithDynamicSizeOneHundred(b *testing.B) {
	testcase := make([]int, 100)
	for i := range testcase {
		testcase[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(testcase)
	}
}
func BenchmarkMergeSortWithDynamicSizeOneThousand(b *testing.B) {
	testcase := make([]int, 1000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(testcase)
	}
}
func BenchmarkMergeSortWithDynamicSizeOneMillion(b *testing.B) {
	testcase := make([]int, 1000000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(testcase)
	}
}

func BenchmarkChannelMergeSortWithNoSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[0].test)
	}
}
func BenchmarkChannelMergeSortWithSizeOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[1].test)
	}
}
func BenchmarkChannelMergeSortWithSizeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[2].test)
	}
}
func BenchmarkChannelMergeSortWithSizeThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[3].test)
	}
}
func BenchmarkChannelMergeSortWithSizeFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[4].test)
	}
}
func BenchmarkChannelMergeSortWithSizeFive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[5].test)
	}
}
func BenchmarkChannelMergeSortWithSizeSix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[6].test)
	}
}
func BenchmarkChannelMergeSortWithSizeSeven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[7].test)
	}
}
func BenchmarkChannelMergeSortWithSizeEight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[8].test)
	}
}
func BenchmarkChannelMergeSortWithSizeNine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[9].test)
	}
}
func BenchmarkChannelMergeSortWithSizeTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testCases[10].test)
	}
}
func BenchmarkChannelMergeSortWithDynamicSizeOneHundred(b *testing.B) {
	testcase := make([]int, 100)
	for i := range testcase {
		testcase[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testcase)
	}
}
func BenchmarkChannelMergeSortWithDynamicSizeOneThousand(b *testing.B) {
	testcase := make([]int, 1000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testcase)
	}
}
func BenchmarkChannelMergeSortWithDynamicSizeOneMillion(b *testing.B) {
	testcase := make([]int, 1000000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelMergeSort(testcase)
	}
}
