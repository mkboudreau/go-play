package helpers

import "fmt"

func OutputCase(zeroBasedIndex int, result string) {
	fmt.Printf("Case #%v: %v", zeroBasedIndex+1, result)
	fmt.Println("")
}
