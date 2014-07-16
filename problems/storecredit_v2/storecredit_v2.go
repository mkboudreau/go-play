package main

import (
	"flag"
	"fmt"
	"github.com/mkboudreau/problems/helpers"
)

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "input", "input.in", "the input file used for processing")
	flag.Parse()

	runner := helpers.NewChallengeRunner(new(StoreCreditChallengeBuilder))
	challenges := runner.BuildChallengesFromFile(inputFile)
	runner.SolveAllChallenges(challenges)
}

type StoreCreditChallenge struct {
	Credit      int
	ItemCount   int
	ItemPrices  []int
	FinalResult [2]int
}

func (c *StoreCreditChallenge) Valid() bool {
	return c.Credit > 0 && c.ItemCount > 0 && c.ItemCount == len(c.ItemPrices)
}

func (c *StoreCreditChallenge) String() string {
	return fmt.Sprint("Credit [", c.Credit, "]; ItemCount [", c.ItemCount, "]; ItemPrices", c.ItemPrices)
}
func (c *StoreCreditChallenge) Answer() string {
	return fmt.Sprintf("%v %v", c.FinalResult[0], c.FinalResult[1])
}

func (c *StoreCreditChallenge) Solve() error {
	for i := 0; i < len(c.ItemPrices)-1; i++ {
		for j := i + 1; j < len(c.ItemPrices); j++ {
			if c.ItemPrices[i]+c.ItemPrices[j] == c.Credit {
				i++
				j++
				if i < j {
					c.FinalResult = [2]int{i, j}
				} else {
					c.FinalResult = [2]int{j, i}
				}
				return nil
			}
		}
	}
	return fmt.Errorf("Could not find two values that add up to %v", c.Credit)
}
func (c *StoreCreditChallenge) ExpectedLinesFromInput() int {
	return 3
}
func (c *StoreCreditChallenge) InputLine(index int, line string) {
	switch index {
	case 0:
		c.Credit = helpers.RequireInt(line)
	case 1:
		c.ItemCount = helpers.RequireInt(line)
	case 2:
		c.ItemPrices = helpers.RequireIntArray(line)
	}
}

type StoreCreditChallengeBuilder struct {
}

func (builder *StoreCreditChallengeBuilder) NewChallenge() helpers.Challenge {
	return &StoreCreditChallenge{}
}
