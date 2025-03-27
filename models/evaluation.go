package models

type Evaluation struct {
	Name        string
	Description string
	Score       int
	MaxScore    int
	MinScore    int
	Weight      int
	Messages    []string
}
