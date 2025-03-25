package evaluators

func NewEvaluation(
	name string,
	desc string,
	score int,
	maxScore int,
	minScore int,
	weight float64,
	msgs []string,
) Evaluation {
	return Evaluation{
		Name:        name,
		Description: desc,
		Score:       score,
		MaxScore:    maxScore,
		MinScore:    minScore,
		Weight:      weight,
		Messages:    msgs,
	}
}
