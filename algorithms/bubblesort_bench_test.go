package algorithms

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSortWithNoSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[0].test)
	}
}
func BenchmarkBubbleSortWithSizeOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[1].test)
	}
}
func BenchmarkBubbleSortWithSizeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[2].test)
	}
}
func BenchmarkBubbleSortWithSizeThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[3].test)
	}
}
func BenchmarkBubbleSortWithSizeFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[4].test)
	}
}
func BenchmarkBubbleSortWithSizeFive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[5].test)
	}
}
func BenchmarkBubbleSortWithSizeSix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[6].test)
	}
}
func BenchmarkBubbleSortWithSizeSeven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[7].test)
	}
}
func BenchmarkBubbleSortWithSizeEight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[8].test)
	}
}
func BenchmarkBubbleSortWithSizeNine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[9].test)
	}
}
func BenchmarkBubbleSortWithSizeTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testCases[10].test)
	}
}

func BenchmarkBubbleSortWithDynamicSizeOneHundred(b *testing.B) {
	testcase := make([]int, 100)
	for i := range testcase {
		testcase[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(testcase)
	}
}
func BenchmarkBubbleSortWithDynamicSizeOneThousand(b *testing.B) {
	testcase := make([]int, 1000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(testcase)
	}
}

/*
func BenchmarkChannelBubbleSortWithNoSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[0].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[1].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[2].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[3].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[4].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeFive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[5].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeSix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[6].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeSeven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[7].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeEight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[8].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeNine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[9].test)
	}
}
func BenchmarkChannelBubbleSortWithSizeTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testCases[10].test)
	}
}
func BenchmarkChannelBubbleSortWithDynamicSizeOneHundred(b *testing.B) {
	testcase := make([]int, 100)
	for i := range testcase {
		testcase[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testcase)
	}
}
func BenchmarkChannelBubbleSortWithDynamicSizeOneThousand(b *testing.B) {
	testcase := make([]int, 1000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testcase)
	}
}
func BenchmarkChannelBubbleSortWithDynamicSizeOneMillion(b *testing.B) {
	testcase := make([]int, 1000000)
	for i := range testcase {
		testcase[i] = rand.Int()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelBubbleSort(testcase)
	}
}
*/
