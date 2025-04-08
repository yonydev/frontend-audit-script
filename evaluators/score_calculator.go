package evaluators

import (
	"fmt"
	"math"
	"os"

	"github.com/yonydev/frontend-audit-script/models"
)

//	type ScaleValueParams struct {
//		Value       float64
//		Min         float64
//		Max         float64
//		NewScaleMin float64
//		NewScaleMax float64
//	}
//
//	func calculateNewScaleValue(params ScaleValueParams) float64 {
//		if params.NewScaleMin == 0 {
//			params.NewScaleMin = 0
//		}
//		if params.NewScaleMax == 0 {
//			params.NewScaleMax = 10
//		}
//		return params.NewScaleMin + ((params.Value-params.Min)*(params.NewScaleMax-params.NewScaleMin))/(params.Max-params.Min)
//	}
func calculateNewScaleValue(value, min, max float64) float64 {
	return 0 + ((value-min)*(10-0))/(max-min)
}

func CalculateScore(evaluations []models.Evaluation) map[string]float64 {
	generalScore := 0.0
	maxScore := 0.0
	minScore := 0.0

	for _, evaluation := range evaluations {
		truncatedScore := math.Max(float64(evaluation.MinScore), math.Min(float64(evaluation.MaxScore), float64(evaluation.Score))) * float64(evaluation.Weight)
		generalScore += truncatedScore
		maxScore += float64(evaluation.MaxScore) * float64(evaluation.Weight)
		minScore += float64(evaluation.MinScore) * float64(evaluation.Weight)
	}

	normalizedScore := calculateNewScaleValue(generalScore, minScore, maxScore)
	// normalizedScore := calculateNewScaleValue()

	githubEnvPath := os.Getenv("GITHUB_ENV")
	if githubEnvPath != "" {
		file, err := os.OpenFile(githubEnvPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err == nil {
			defer file.Close()
			_, _ = fmt.Fprintf(file, "EVALUATION_TOTAL_SCORE=%.2f\n", generalScore)
		} else {
			fmt.Printf("⚠️ Failed to write to GITHUB_ENV: %v\n", err)
		}
	} else {
		fmt.Println("⚠️ GITHUB_ENV is not set, skipping environment export.")
	}

	output := map[string]float64{
		"general":         generalScore,
		"max":             maxScore,
		"min":             minScore,
		"normalizedScore": normalizedScore,
	}

	fmt.Println("This is the output of the evaluation: ", output)

	return map[string]float64{
		"general":         generalScore,
		"max":             maxScore,
		"min":             minScore,
		"normalizedScore": normalizedScore,
	}
}

// func normalizeWeights(evaluations []models.Evaluation) []float64 {
// 	var totalWeight int
// 	for _, eval := range evaluations {
// 		totalWeight += eval.Weight
// 	}
//
// 	normalizedWeights := make([]float64, len(evaluations))
// 	for i, eval := range evaluations {
// 		normalizedWeights[i] = float64(eval.Weight) / float64(totalWeight)
// 	}
//
// 	return normalizedWeights
// }
//
// func normalizeScore(score, minScore, maxScore int) float64 {
// 	if maxScore == minScore {
// 		return 0 // Prevent division by zero
// 	}
// 	return float64(score-minScore) / float64(maxScore-minScore) * 100
// }
//
// func CalculateScore(evaluations []models.Evaluation) float64 {
// 	var totalScore float64
// 	normalizedWeights := normalizeWeights(evaluations)
//
// 	for i, eval := range evaluations {
// 		if eval.MaxScore == 0 || eval.MaxScore == eval.MinScore {
// 			fmt.Printf("🚨 Skipping %s due to invalid score range.\n", eval.Name)
// 			continue
// 		}
//
// 		normalizedScore := normalizeScore(eval.Score, eval.MinScore, eval.MaxScore)
// 		totalScore += normalizedScore * normalizedWeights[i]
//
// 		if eval.Score < eval.MinScore {
// 			fmt.Printf("❌ Critical issue in %s: Score (%d) is below the minimum allowed (%d).\n", eval.Name, eval.Score, eval.MinScore)
// 		}
// 	}
//
// 	formattedTotalScore := strconv.FormatFloat(totalScore, 'f', 2, 64)
// 	githubEnvPath := os.Getenv("GITHUB_ENV")
//
// 	if githubEnvPath != "" {
// 		file, err := os.OpenFile(githubEnvPath, os.O_APPEND|os.O_WRONLY, 0644)
// 		if err == nil {
// 			defer file.Close()
// 			_, _ = fmt.Fprintf(file, "EVALUATION_TOTAL_SCORE=%.2f\n", totalScore)
// 		} else {
// 			fmt.Printf("⚠️ Failed to write to GITHUB_ENV: %v\n", err)
// 		}
// 	} else {
// 		fmt.Println("⚠️ GITHUB_ENV is not set, skipping environment export.")
// 	}
//
// 	fmt.Println(
// 		"Total Score: ",
// 		c.InfoFgBold(formattedTotalScore),
// 	)
//
// 	return totalScore
// }
