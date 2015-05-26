package algorithms

func QuickSort(set []int) []int {
	quickRecursive(set, 0, len(set))
	return set
}

func quickRecursive(set []int, left int, right int) {
	if (right - left) <= 1 {
		return
	}
	pivotIndex := int((right - left) / 2)
	pivotValue := set[pivotIndex]

	var i, j, tmp int = left, right - 1, 0

	for i <= j {
		switch {
		case set[i] >= pivotValue && pivotValue >= set[j]:
			tmp = set[i]
			set[i] = set[j]
			set[j] = tmp
		}

		i++
		j--
	}

	quickRecursive(set, left, i-1)
	quickRecursive(set, i, right)
}

/*
func QuickSort_V1(set []int) []int {
	if len(set) <= 1 {
		return set
	}
	pivot := set[int(len(set)/2)]

	leftValues := make([]int, 0)
	pivotValues := make([]int, 0)
	rightValues := make([]int, 0)

	for i := 0; i < len(set); i++ {
		switch {
		case set[i] < pivot:
			leftValues = append(leftValues, set[i])
		case set[i] > pivot:
			rightValues = append(rightValues, set[i])
		case set[i] == pivot:
			pivotValues = append(pivotValues, set[i])
		}
	}

	leftValues = QuickSort(leftValues)
	rightValues = QuickSort(rightValues)

	leftValues = append(leftValues, pivotValues...)
	leftValues = append(leftValues, rightValues...)

	return leftValues
}

func QuickSortV2(set []int) []int {
	if len(set) <= 1 {
		return set
	}
	pivotIndex := int(len(set) / 2)
	pivotValue := set[pivotIndex]

	var tmp int
	var i, j int = 0

	for i, j := 0, len(set)-1; i <= j; i++ {
		switch {
		case set[i] >= pivotValue && pivotValue >= set[j]:
			tmp = set[i]
			set[i] = set[j]
			set[j] = tmp
		}

		j--
	}

	QuickSort(set[pivotIndex:])
	QuickSort(set[:pivotIndex])

	return set
}
*/
