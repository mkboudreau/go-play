package algorithms

func MergeSort(set []int) []int {
	if len(set) > 1 {
		splitIndex := int(len(set) / 2)
		return Merge(MergeSort(set[:splitIndex]), MergeSort(set[splitIndex:]))
	}
	return set
}

func Merge(iSet, jSet []int) []int {
	i, j, k := 0, 0, 0

	kSet := make([]int, len(iSet)+len(jSet))

	for ; i < len(iSet) && j < len(jSet); k++ {
		if iSet[i] < jSet[j] {
			kSet[k] = iSet[i]
			i++
		} else {
			kSet[k] = jSet[j]
			j++
		}
	}
	for ; i < len(iSet); k++ {
		kSet[k] = iSet[i]
		i++
	}
	for ; j < len(jSet); k++ {
		kSet[k] = jSet[j]
		j++
	}

	return kSet
}
