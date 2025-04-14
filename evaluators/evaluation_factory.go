package evaluators

import "github.com/yonydev/frontend-audit-script/models"

func NewEvaluation(
	name string,
	desc string,
	score int,
	maxScore int,
	minScore int,
	weight int,
	msgs []string,
) models.Evaluation {
	return models.Evaluation{
		Name:        name,
		Description: desc,
		Score:       score,
		MaxScore:    maxScore,
		MinScore:    minScore,
		Weight:      weight,
		Messages:    msgs,
	}
}
