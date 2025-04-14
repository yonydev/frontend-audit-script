package evaluators

import (
	"fmt"
	"math"
	"os"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/models"
)

// calculateNewScaleValue scales a given value from one range to another specified range.
// It takes a ScaleEvaluationValues struct as input, which includes:
// - Value: The value to be scaled.
// - Min: The minimum value of the original range.
// - Max: The maximum value of the original range.
// - NewScaleMin: The minimum value of the new range (default is 0 if not set).
// - NewScaleMax: The maximum value of the new range (default is 10 if not set).
// The function returns the scaled value in the new range.
func calculateNewScaleValue(params models.ScaleEvaluationValues) float64 {
	if params.NewScaleMin == 0 {
		params.NewScaleMin = 0
	}
	if params.NewScaleMax == 0 {
		params.NewScaleMax = 10
	}
	return params.NewScaleMin + ((params.Value-params.Min)*(params.NewScaleMax-params.NewScaleMin))/(params.Max-params.Min)
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

	normalizedScore := calculateNewScaleValue(models.ScaleEvaluationValues{Value: generalScore, Min: minScore, Max: maxScore, NewScaleMin: 0, NewScaleMax: 10})
	truncateNormalizedScore := math.Round(math.Max(float64(minScore), math.Min(float64(maxScore), float64(normalizedScore)))*100) / 100

	githubEnvPath := os.Getenv("GITHUB_ENV")
	if githubEnvPath != "" {
		file, err := os.OpenFile(githubEnvPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err == nil {
			defer file.Close()
			_, _ = fmt.Fprintf(file, "EVALUATION_TOTAL_SCORE=%.2f\n", truncateNormalizedScore)
		} else {
			fmt.Printf("⚠️ Failed to write to GITHUB_ENV: %v\n", err)
		}
	} else {
		fmt.Println("⚠️ GITHUB_ENV is not set, skipping environment export.")
	}

	finalResult := map[string]float64{
		"general":         generalScore,
		"max":             maxScore,
		"min":             minScore,
		"normalizedScore": truncateNormalizedScore,
	}

	fmt.Print("\n")
	fmt.Println("This is the general score:", c.InfoFgBold(generalScore))
	fmt.Println("This is the max score:", c.InfoFgBold(maxScore))
	fmt.Println("This is the min score:", c.InfoFgBold(minScore))
	fmt.Println("This is the normalized score:", c.InfoFgBold(truncateNormalizedScore))

	return finalResult
}
