package algorithms

func ChannelMergeSort(set []int) []int {
	if len(set) > 1 {
		splitIndex := int(len(set) / 2)

		ch := make(chan []int)
		defer close(ch)
		go func() {
			ch <- ChannelMergeSort(set[:splitIndex])
		}()
		go func() {
			ch <- ChannelMergeSort(set[splitIndex:])
		}()

		return Merge(<-ch, <-ch)
	}
	return set
}
