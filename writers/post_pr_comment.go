package writers

import (
	"fmt"
	"strings"

	"github.com/sethvargo/go-githubactions"
	"github.com/yonydev/frontend-audit-script/models"
)

func PostPrComment(evaluations []models.Evaluation) {
	action := githubactions.New()
	prNumber := action.GetInput("pr_number")

	if prNumber == "" {
		action.Warningf("No PR number found, skipping comment")
		return
	}

	// Create Markdown Table Header
	var tableBuilder strings.Builder
	tableBuilder.WriteString("### 🛠 Frontend Evaluations Report\n\n")
	tableBuilder.WriteString("| **Metric**       | **Score**  | **Messages** |\n")
	tableBuilder.WriteString("|-------------------|:------------:|:------------|\n")

	// Populate the table with evaluation results
	totalScore := 0
	for _, eval := range evaluations {
		tableBuilder.WriteString(fmt.Sprintf(
			"| %s | %d / 100 | %s |\n",
			eval.Name, eval.Score, strings.Join(eval.Messages, ", "),
		))
		totalScore += eval.Score
	}

	tableBuilder.WriteString(fmt.Sprintf(
		"| ⭐ **Final Score** | **%v** / 100 | |\n",
		23,
	))

	// Print for debugging
	// fmt.Println(tableBuilder.String())
  //
  action.SetEnv("EVALUATION_PR_COMMENT", tableBuilder.String())

}
