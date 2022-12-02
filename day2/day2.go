// This one is solved using less logic and more hard coded maps to describe
// the game and strategy rules.
//
// The downside to this approach is that the actual rules of the game are
// not written here, and it is verbose.
//
// The upside is that the behaviour at each step is fairly easy to understand
// and potentially performant (though this is unmeasured)
//
// An interesting outcome of taking this approach is that there are no if statements.
package day2

import "strings"

const (
	opponentRock     = "A"
	opponentPaper    = "B"
	opponentScissors = "C"

	responseRock     = "X"
	responsePaper    = "Y"
	responseScissors = "Z"

	loseScore = 0
	drawScore = 3
	winScore  = 6
)

var shapeScores = map[string]int{
	responseRock:     1,
	responsePaper:    2,
	responseScissors: 3,
}

type round struct {
	Opponent string
	Response string
}

var outcomes = map[round]int{
	{opponentRock, responseRock}:     drawScore,
	{opponentRock, responsePaper}:    winScore,
	{opponentRock, responseScissors}: loseScore,

	{opponentPaper, responseRock}:     loseScore,
	{opponentPaper, responsePaper}:    drawScore,
	{opponentPaper, responseScissors}: winScore,

	{opponentScissors, responseRock}:     winScore,
	{opponentScissors, responsePaper}:    loseScore,
	{opponentScissors, responseScissors}: drawScore,
}

// Day2 returns the score of following the input strategy guide
func Day2(guide string) int {
	lines := strings.Split(guide, "\n")
	total := 0
	for _, line := range lines {
		thisRound := round{
			string(line[0]),
			string(line[2]),
		}

		total += score(thisRound)
	}
	return total
}

func score(r round) int {
	return outcomes[r] + shapeScores[r.Response]
}

// Part 2 solution

const (
	strategyLose = "X"
	strategyDraw = "Y"
	strategyWin  = "Z"
)

var strategyMap = map[round]round{
	{opponentRock, strategyLose}: {opponentRock, responseScissors},
	{opponentRock, strategyDraw}: {opponentRock, responseRock},
	{opponentRock, strategyWin}:  {opponentRock, responsePaper},

	{opponentPaper, strategyLose}: {opponentPaper, responseRock},
	{opponentPaper, strategyDraw}: {opponentPaper, responsePaper},
	{opponentPaper, strategyWin}:  {opponentPaper, responseScissors},

	{opponentScissors, strategyLose}: {opponentScissors, responsePaper},
	{opponentScissors, strategyDraw}: {opponentScissors, responseScissors},
	{opponentScissors, strategyWin}:  {opponentScissors, responseRock},
}

// Day2 returns the score of following the input strategy guide
func Day2Part2(guide string) int {
	lines := strings.Split(guide, "\n")
	total := 0
	for _, line := range lines {
		thisRoundStrategy := round{
			string(line[0]),
			string(line[2]),
		}
		thisRoundHands := strategyMap[thisRoundStrategy]

		total += score(thisRoundHands)
	}
	return total
}
