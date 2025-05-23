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

func EvalStylingLibs(content *string) (models.Evaluation, error) {
	var packageJSON map[string]any
	var foundStylingLibs []string
	var allowedStylingLibs []string
	var disallowedStylingLibs []string
	var messages []string

	evalName := ">> Styling Libraries"
	evalDesc := "\nChecking for common styling libraries...\n"

	score := 0
	maxScore := 2
	minScore := -3
	weight := 3

	if err := json.Unmarshal([]byte(*content), &packageJSON); err != nil {
		return models.Evaluation{}, fmt.Errorf("failed to parse package.json: %v", err)
	}

	dependencies, foundDeps := packageJSON["dependencies"].(map[string]any)
	devDependenciesMerge, foundDevDeps := packageJSON["devDependencies"].(map[string]any)
	// Merge dependencies and devDependencies -> devDependenciesMerge
	maps.Copy(devDependenciesMerge, dependencies)

	if !foundDeps && !foundDevDeps {
		messages = append(
			messages,
			c.WarningFg("No dependencies nor devDependencies found in package.json"),
		)
	}

	for lib, allowed := range utils.StylingLibs {
		if _, exists := devDependenciesMerge[lib]; exists {
			foundStylingLibs = append(foundStylingLibs, lib)
			if !allowed {
				disallowedStylingLibs = append(disallowedStylingLibs, lib)
			} else {
				allowedStylingLibs = append(allowedStylingLibs, lib)
			}
		}
	}

	if len(foundStylingLibs) > 0 {
		messages = append(
			messages,
			fmt.Sprintf(
				"Found %s in package.json.\n",
				c.InfoFgBold(strings.Join(foundStylingLibs, ", ")),
			),
		)
	}

	if len(allowedStylingLibs) > 0 {
		messages = append(
			messages,
			c.SuccessFg(fmt.Sprintf("✅ %s are allowed and needed for '@clipmx/cods-ui'.\n", strings.Join(allowedStylingLibs, ", "))),
		)
	}

	if len(disallowedStylingLibs) > 0 {
		messages = append(
			messages,
			c.WarningFg(fmt.Sprintf("❗ %s are recommended to be removed from project.\n", strings.Join(disallowedStylingLibs, ", "))),
		)
	}

	if len(allowedStylingLibs) == 2 {
		score = maxScore
	} else if len(allowedStylingLibs) == 1 {
		score = 1
	} else {
		score = minScore
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
	writers.SetEvaluationEnvVariables(evaluation, utils.StylingLibsEnvVars)

	return evaluation, nil
}
