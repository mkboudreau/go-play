package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Challenge interface {
	ExpectedLinesFromInput() int
	InputLine(index int, line string)
	Valid() bool
	Solve() error
	Answer() string
}

type ChallengeBuilder interface {
	NewChallenge() Challenge
}

type ChallengeRunner struct {
	ChallengeBuilder ChallengeBuilder
}

func NewChallengeRunner(builder ChallengeBuilder) *ChallengeRunner {
	return &ChallengeRunner{
		ChallengeBuilder: builder,
	}
}

func (runner *ChallengeRunner) BuildChallengesFromFile(inputFile string) []Challenge {
	if len(inputFile) <= 0 || strings.TrimSpace(inputFile) == "" {
		panic(fmt.Errorf("Did not specify input file"))
	}

	return runner.convertInputFileToChallenges(inputFile)
}

func (runner *ChallengeRunner) BuildChallengesFromString(inputData string) []Challenge {
	reader := strings.NewReader(inputData)
	return runner.readInAllChallenges(bufio.NewReader(reader))
}

func (runner *ChallengeRunner) SolveAllChallenges(challenges []Challenge) {
	for i, c := range challenges {
		err := c.Solve()
		if err != nil {
			fmt.Println("ERROR: ", err)
		} else {
			OutputCase(i, c.Answer())
		}
	}
}

func (runner *ChallengeRunner) convertInputFileToChallenges(inputFile string) []Challenge {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("Could open file to read, err: %v", err))
	}

	reader := bufio.NewReader(file)
	defer file.Close()

	return runner.readInAllChallenges(reader)
}

func (runner *ChallengeRunner) readInAllChallenges(r *bufio.Reader) []Challenge {
	line, err := r.ReadString('\n')
	if err != nil {
		panic(fmt.Errorf("Could not read in all challenges, err: %v", err))
	}
	line = strings.TrimSuffix(line, "\n")
	challengeCount := RequireInt(line)
	challenges := make([]Challenge, challengeCount)
	for i := 0; i < challengeCount; i++ {
		challenges[i] = runner.readInSingleChallenge(r)
		if !challenges[i].Valid() {
			panic(fmt.Errorf("Invalid Challenge %v", challenges[i]))
		}
	}
	return challenges
}

func (runner *ChallengeRunner) readInSingleChallenge(r *bufio.Reader) Challenge {
	challenge := runner.ChallengeBuilder.NewChallenge()

	for i := 0; i < challenge.ExpectedLinesFromInput(); i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSpace(line)
		challenge.InputLine(i, line)
	}
	return challenge
}
