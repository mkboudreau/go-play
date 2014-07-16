package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func RequireInt(line string) int {
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(fmt.Errorf("Could not read int: %v", err))
	}
	return i
}

func RequireIntArray(line string) []int {
	tokenized := strings.Split(strings.TrimSpace(line), " ")
	intArray := make([]int, len(tokenized))
	for i := 0; i < len(intArray); i++ {
		intArray[i] = RequireInt(tokenized[i])
	}
	return intArray
}

func RequireStringArray(line string) []string {
	return strings.Split(strings.TrimSpace(line), " ")
}
