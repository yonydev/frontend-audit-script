package evaluators

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/utils"
)

func EvalStylingLibs(content *string) (Evaluation, error) {
	var packageJSON map[string]any
	var foundStylingLibs []string
	var allowedStylingLibs []string
	var disallowedStylingLibs []string
	var messages []string

	evalName := "\n>>> Styling Libraries\n"
	evalDesc := "Checking for common styling libraries...\n"

	initialScore := 100
	penaltyPoints := 0
	minScore := 30
	maxScore := 100
	weight := 2

	if err := json.Unmarshal([]byte(*content), &packageJSON); err != nil {
		return Evaluation{}, fmt.Errorf("failed to parse package.json: %v", err)
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
				penaltyPoints += 10
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

	if len(disallowedStylingLibs) > 0 {
		messages = append(
			messages,
			c.WarningFg(fmt.Sprintf("❗ %s are recommended to be removed from project.\n", strings.Join(disallowedStylingLibs, ", "))),
		)
	}

	if len(allowedStylingLibs) == 2 {
		initialScore = 100
	} else if len(allowedStylingLibs) == 1 {
		initialScore = 70
	} else {
		initialScore = minScore
	}

	// Subtract penalty points from initial score
	score := max(initialScore-penaltyPoints, minScore)

	if len(allowedStylingLibs) > 0 {
		messages = append(
			messages,
			c.SuccessFg(fmt.Sprintf("✅ %s are allowed and needed for '@clipmx/cods-ui'.\n", strings.Join(allowedStylingLibs, ", "))),
		)
	}

	return NewEvaluation(
			evalName,
			evalDesc,
			score,
			maxScore,
			minScore,
			weight,
			messages,
		),
		nil
}
