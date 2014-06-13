package algorithms

/*
type Comparable interface {
	Compare(c1 Comparable) int
}

func (i1 int) Compare(c2 Comparable) int {
	i2 := int(c2)
	if i1 < i2 {
		return -1
	} else if i1 > i2 {
		return 1
	} else {
		return 0
	}
}

func ComparableMergeSort(set []Comparable) []Comparable {
	if len(set) > 1 {
		splitIndex := int(len(set) / 2)
		return ComparableMerge(ComparableMergeSort(set[:splitIndex]), ComparableMergeSort(set[splitIndex:]))
	}
	return set
}

func ComparableMerge(iSet, jSet []Comparable) []Comparable {
	i, j, k := 0, 0, 0

	kSet := make([]Comparable, len(iSet)+len(jSet))

	for ; i < len(iSet) && j < len(jSet); k++ {
		if iSet[i].Compare(jSet[j]) < 0 {
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
*/
