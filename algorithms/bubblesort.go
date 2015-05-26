package algorithms

/*
func BubbleSort(set []int) []int {

	var tmp int
	for i := 0; i < len(set); i++ {
		// 10, 4, 1 , 2, 3, 4 , 5, 1, 3

		for j := len(set) - 1; j > i; j-- {
			if set[i] > set[j] {
				tmp = set[j]
				set[j] = set[i]
				set[i] = tmp
			}
		}
	}

	return set
}
*/

func BubbleSort(set []int) []int {
	var tmp int
	var done bool

	for {
		done = true
		for i := 0; i < len(set)-1; i++ {
			if set[i] > set[i+1] {
				done = false
				tmp = set[i+1]
				set[i+1] = set[i]
				set[i] = tmp
			}
		}
		if done {
			break
		}
	}

	return set
}
