package stats

import (
	"fmt"
	"strconv"

	aphrodite "github.com/jonathon-chew/Aphrodite"
)

type Stats struct {
	Number int
	Errors int
}

type IssueTracking struct {
	Variables Stats
	Functions Stats
	KeyWords  Stats
	Comments  Stats
}

func PrintStats(stats IssueTracking) {
	aphrodite.PrintInfo(fmt.Sprintf("Number of Variables %s\n", strconv.Itoa(stats.Variables.Number)))
	aphrodite.PrintError(fmt.Sprintf("Number of Variables with issues %s\n", strconv.Itoa(stats.Variables.Errors)))
	aphrodite.PrintInfo(fmt.Sprintf("Number of Functions %s\n", strconv.Itoa(stats.Functions.Number)))
	aphrodite.PrintError(fmt.Sprintf("Number of Functions with issues %s\n", strconv.Itoa(stats.Functions.Errors)))
	aphrodite.PrintInfo(fmt.Sprintf("Number of Keywords %s\n", strconv.Itoa(stats.KeyWords.Number)))
	aphrodite.PrintError(fmt.Sprintf("Number of Keywords with issues %s\n", strconv.Itoa(stats.KeyWords.Errors)))
	aphrodite.PrintInfo(fmt.Sprintf("Number of Comments %s\n", strconv.Itoa(stats.Comments.Number)))
}
