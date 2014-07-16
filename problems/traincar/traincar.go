package main


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

	runner := helpers.NewChallengeRunner(new(TrainCarChallengeBuilder))
	challenges := runner.BuildChallengesFromFile(inputFile)
	runner.SolveAllChallenges(challenges)
}

type TrainCarChallenge struct {
	TrainCount int
	Trains     []string
	FinalResult int
}

func (c *TrainCarChallenge) Valid() bool {
	return c.TrainCount > 0 && c.TrainCount == len(c.Trains)
}
func (c *TrainCarChallenge) String() string {
	return fmt.Sprint("TrainCount [", c.TrainCount, "]; Trains", c.Trains)
}
func (c *TrainCarChallenge) Answer() string {
	return fmt.Sprintf("%v", c.FinalResult)
}

func (c *TrainCarChallenge) Solve() error {
	//TODO: IMPLEMENT SOLVE
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
func (c *TrainCarChallenge) ExpectedLinesFromInput() int {
	return 2
}
func (c *TrainCarChallenge) InputLine(index int, line string) {
	switch index {
	case 0:
		c.TrainCount = helpers.RequireInt(line)
	case 1:
		c.Trains = helpers.RequireStringArray(line)
	}
}

type TrainCarChallengeBuilder struct {
}

func (builder *TrainCarChallengeBuilder) NewChallenge() helpers.Challenge {
	return &TrainCarChallenge{}
}