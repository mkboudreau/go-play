package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mkboudreau/problems/helpers"
	"os"
	"strings"
)

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "input", "input.in", "the input file used for processing")
	flag.Parse()

	if strings.TrimSpace(inputFile) == "" {
		panic(fmt.Errorf("Did not specify input file"))
	}

	cases := convertInputFileToCases(inputFile)

	SolveAllCases(cases)
}

type Case struct {
	Credit     int
	ItemCount  int
	ItemPrices []int
}

func (c *Case) Valid() bool {
	return c.Credit > 0 && c.ItemCount > 0 && c.ItemCount == len(c.ItemPrices)
}

func (c *Case) String() string {
	return fmt.Sprint("Credit [", c.Credit, "]; ItemCount [", c.ItemCount, "]; ItemPrices", c.ItemPrices)
}

func SolveAllCases(cases []*Case) {
	for i, c := range cases {
		answer, err := SolveCase(c)
		if err != nil {
			fmt.Println("ERROR: ", err)
		} else {
			helpers.OutputCase(i, fmt.Sprintf("%v %v", answer[0], answer[1]))
		}
	}
}

func SolveCase(c *Case) ([]int, error) {
	for i := 0; i < len(c.ItemPrices)-1; i++ {
		for j := i + 1; j < len(c.ItemPrices); j++ {
			if c.ItemPrices[i]+c.ItemPrices[j] == c.Credit {
				i++
				j++
				if i < j {
					return []int{i, j}, nil
				} else {
					return []int{j, i}, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Could not find two values that add up to %v", c.Credit)
}

func convertInputFileToCases(inputFile string) []*Case {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("Could open file to read, err: %v", err))
	}

	reader := bufio.NewReader(file)
	defer file.Close()

	return readInAllCases(reader)
}

func readInAllCases(r *bufio.Reader) []*Case {
	line, err := r.ReadString('\n')
	if err != nil {
		panic(fmt.Errorf("Could not read in all cases, err: %v", err))
	}
	line = strings.TrimSuffix(line, "\n")
	numOfTestCases := helpers.RequireInt(line)
	cases := make([]*Case, numOfTestCases)
	for i := 0; i < numOfTestCases; i++ {
		cases[i] = readInSingleCase(r)
		if !cases[i].Valid() {
			panic(fmt.Errorf("Invalid Test Case %v", cases[i]))
		}
	}
	return cases
}

func readInSingleCase(r *bufio.Reader) *Case {
	inputCase := new(Case)

	for i := 0; i < 3; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSpace(line)
		switch i {
		case 0:
			inputCase.Credit = helpers.RequireInt(line)
		case 1:
			inputCase.ItemCount = helpers.RequireInt(line)
		case 2:
			inputCase.ItemPrices = helpers.RequireIntArray(line)
		}
	}
	return inputCase
}
