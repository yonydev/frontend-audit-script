package evaluators

type Evaluation struct {
	Name        string
	Description string
	Score       int
	MaxScore    int
	MinScore    int
	Weight      float64
	Messages    []string
}
