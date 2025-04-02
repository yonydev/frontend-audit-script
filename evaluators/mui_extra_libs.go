package evaluators

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/models"
	"github.com/yonydev/frontend-audit-script/utils"
	"github.com/yonydev/frontend-audit-script/writers"
)

func EvalMuiExtraLibs(content *string) (models.Evaluation, error) {
	var packageJSON map[string]any
	var messages []string
	var foundLibs []string

	evalName := ">>> MUI Extra Libraries"
	evalDesc := "\nChecking for MUI extra libraries...\n"

	score := 0
	weight := 1
	minScore := 40
	maxScore := 100

	if err := json.Unmarshal([]byte(*content), &packageJSON); err != nil {
		return models.Evaluation{}, fmt.Errorf("failed to parse package.json: %v", err)
	}

	dependencies, foundDeps := packageJSON["dependencies"].(map[string]any)
	devDependenciesMerge, foundDevDeps := packageJSON["devDependencies"].(map[string]any)

	maps.Copy(devDependenciesMerge, dependencies)

	if !foundDeps && !foundDevDeps {
		messages = append(
			messages,
			c.WarningFg("No dependencies nor devDependencies found in package.json"),
		)
	}

	for _, lib := range utils.MuiExtraLibs {
		if _, exists := devDependenciesMerge[lib]; exists {
			foundLibs = append(foundLibs, lib)
		}
	}

	switch len(foundLibs) {
	case 0:
		score = maxScore
		messages = append(
			messages,
			c.SuccessFg("No MUI extra libraries found. Nice, keep it up! 🦾, keep it clean! 🧹"),
		)
	case 1:
		score = 70
		messages = append(
			messages,
			fmt.Sprintf(
				"Found %s in package.json.\n",
				c.InfoFgBold(strings.Join(foundLibs, ", ")),
			),
		)
	default:
		score = minScore
		messages = append(
			messages,
			fmt.Sprintf(
				"Found %s in package.json.\n",
				c.InfoFgBold(strings.Join(foundLibs, ", ")),
			),
		)
	}

	evaluation := NewEvaluation(
		evalName,
		evalDesc,
		score,
		maxScore,
		minScore,
		weight,
		messages,
	)
	writers.SetEvaluationEnvVariables(evaluation, utils.MuiExtraLibsEnvVars)

	return evaluation, nil
}
