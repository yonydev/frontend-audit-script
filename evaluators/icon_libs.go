package evaluators

import (
	"encoding/json"
	"fmt"
	"maps"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/models"
	"github.com/yonydev/frontend-audit-script/utils"
	"github.com/yonydev/frontend-audit-script/writers"
)

func EvalIconLibs(content *string) (models.Evaluation, error) {
	var packageJSON map[string]any
	var foundIconsLibs []string
	var evalMessages []string

	evalName := ">>> Icon Libraries"
	evalDesc := "\nChecking for common icon libraries...\n"
	foundLibsCount := 0

	score := 0
	weight := 3
	maxScore := 0
	minScore := -3

	if err := json.Unmarshal([]byte(*content), &packageJSON); err != nil {
		return models.Evaluation{}, fmt.Errorf("failed to parse package.json: %v", err)
	}

	dependencies, foundDeps := packageJSON["dependencies"].(map[string]any)
	devDependenciesMerge, foundDevDeps := packageJSON["devDependencies"].(map[string]any)
	// Merge dependencies and devDependencies -> devDependenciesMerge
	maps.Copy(devDependenciesMerge, dependencies)

	if !foundDeps && !foundDevDeps {
		return NewEvaluation(
				evalName,
				evalDesc,
				score,
				maxScore,
				minScore,
				weight,
				[]string{"No dependencies nor devDependencies found in package.json"}),
			nil
	}

	for _, lib := range utils.CommonIconLibs {
		if _, found := devDependenciesMerge[lib]; found {
			foundLibsCount++
			foundIconsLibs = append(foundIconsLibs, lib)
		}
	}

	switch foundLibsCount {
	case 0:
		score = 0
		evalMessages = append(
			evalMessages,
			c.WarningFg("No icon library found. Consider adding one for consistent icon usage.\n"),
		)
	case 1:
		score = 0
		evalMessages = append(evalMessages, fmt.Sprintf(
			"Using a single icon library: %s, which is ideal.\n",
			c.InfoFgBold(foundIconsLibs[0])),
		)
	case 2:
		score = -2
		evalMessages = append(evalMessages, fmt.Sprintf(
			"Using: %s, please just use 1 for better performance.\n",
			c.InfoFgBold(foundIconsLibs)),
		)
	default:
		score = minScore
		evalMessages = append(
			evalMessages,
			fmt.Sprintf(
				"%s icon libraries found in package.json. Consider using a single icon library for consistent icon usage.\n",
				c.InfoFgBold(foundLibsCount),
			),
		)
		for _, lib := range foundIconsLibs {
			evalMessages = append(evalMessages, fmt.Sprintf("- %s", c.WarningFgBold(lib)))
		}
	}

	evaluation := NewEvaluation(
		evalName,
		evalDesc,
		score,
		maxScore,
		minScore,
		weight,
		evalMessages,
	)
	writers.SetEvaluationEnvVariables(evaluation, utils.IconLibsEnvVars)

	return evaluation, nil
}
