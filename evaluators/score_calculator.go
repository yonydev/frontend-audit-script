package evaluators

import (
	"fmt"
	"strconv"

	"github.com/sethvargo/go-githubactions"
	"github.com/yonydev/frontend-audit-script/models"
)

func normalizeWeights(evaluations []models.Evaluation) []float64 {
	var totalWeight int
	for _, eval := range evaluations {
		totalWeight += eval.Weight
	}

	normalizedWeights := make([]float64, len(evaluations))
	for i, eval := range evaluations {
		normalizedWeights[i] = float64(eval.Weight) / float64(totalWeight)
	}

	return normalizedWeights
}

func normalizeScore(score, minScore, maxScore int) float64 {
	if maxScore == minScore {
		return 0 // Prevent division by zero
	}
	return float64(score-minScore) / float64(maxScore-minScore) * 100
}

func CalculateScore(evaluations []models.Evaluation) float64 {
	var totalScore float64
	normalizedWeights := normalizeWeights(evaluations)

	for i, eval := range evaluations {
		if eval.MaxScore == 0 || eval.MaxScore == eval.MinScore {
			fmt.Printf("🚨 Skipping %s due to invalid score range.\n", eval.Name)
			continue
		}

		normalizedScore := normalizeScore(eval.Score, eval.MinScore, eval.MaxScore)
		totalScore += normalizedScore * normalizedWeights[i]

		if eval.Score < eval.MinScore {
			fmt.Printf("❌ Critical issue in %s: Score (%d) is below the minimum allowed (%d).\n", eval.Name, eval.Score, eval.MinScore)
		}
	}

	githubactions.SetEnv("EVALUATION_TOTAL_SCORE", strconv.FormatFloat(totalScore, 'f', 2, 64))

	return totalScore
}
