package algorithms

import (
	"math/rand"
	"testing"
)

func BenchmarkQuickSortWithNoSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[0].test)
	}
}
func BenchmarkQuickSortWithSizeOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[1].test)
	}
}
func BenchmarkQuickSortWithSizeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[2].test)
	}
}
func BenchmarkQuickSortWithSizeThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[3].test)
	}
}
func BenchmarkQuickSortWithSizeFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[4].test)
	}
}
func BenchmarkQuickSortWithSizeFive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[5].test)
	}
}
func BenchmarkQuickSortWithSizeSix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[6].test)
	}
}
func BenchmarkQuickSortWithSizeSeven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[7].test)
	}
}
func BenchmarkQuickSortWithSizeEight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[8].test)
	}
}
func BenchmarkQuickSortWithSizeNine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[9].test)
	}
}
func BenchmarkQuickSortWithSizeTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(testCases[10].test)
	}
}

func BenchmarkQuickSortWithDynamicSizeOneHundred(b *testing.B) {
	testcase := make([]int, 100)
	for i := range testcase {
		testcase[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(testcase)
	}
}
func BenchmarkQuickSortWithDynamicSizeOneThousand(b *testing.B) {
	testcase := make([]int, 1000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(testcase)
	}
}
func BenchmarkQuickSortWithDynamicSizeOneMillion(b *testing.B) {
	testcase := make([]int, 1000000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(testcase)
	}
}

/*
func BenchmarkChannelQuickSortWithNoSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[0].test)
	}
}
func BenchmarkChannelQuickSortWithSizeOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[1].test)
	}
}
func BenchmarkChannelQuickSortWithSizeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[2].test)
	}
}
func BenchmarkChannelQuickSortWithSizeThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[3].test)
	}
}
func BenchmarkChannelQuickSortWithSizeFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[4].test)
	}
}
func BenchmarkChannelQuickSortWithSizeFive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[5].test)
	}
}
func BenchmarkChannelQuickSortWithSizeSix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[6].test)
	}
}
func BenchmarkChannelQuickSortWithSizeSeven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[7].test)
	}
}
func BenchmarkChannelQuickSortWithSizeEight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[8].test)
	}
}
func BenchmarkChannelQuickSortWithSizeNine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[9].test)
	}
}
func BenchmarkChannelQuickSortWithSizeTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testCases[10].test)
	}
}
func BenchmarkChannelQuickSortWithDynamicSizeOneHundred(b *testing.B) {
	testcase := make([]int, 100)
	for i := range testcase {
		testcase[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testcase)
	}
}
func BenchmarkChannelQuickSortWithDynamicSizeOneThousand(b *testing.B) {
	testcase := make([]int, 1000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testcase)
	}
}
func BenchmarkChannelQuickSortWithDynamicSizeOneMillion(b *testing.B) {
	testcase := make([]int, 1000000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelQuickSort(testcase)
	}
}
*/
