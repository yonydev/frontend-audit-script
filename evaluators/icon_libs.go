package evaluators

import (
	"encoding/json"
	"fmt"
	"maps"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/utils"
)

func EvalIconLibs(content *string) (Evaluation, error) {
	var packageJSON map[string]any
	var foundIconsLibs []string
	var evalMessages []string

	evalName := "\n>>> Icon Libraries\n"
	evalDesc := "Checking for common icon libraries...\n"
	foundLibsCount := 0
	score := 100

	if err := json.Unmarshal([]byte(*content), &packageJSON); err != nil {
		return Evaluation{}, fmt.Errorf("failed to parse package.json: %v", err)
	}

	dependencies, foundDeps := packageJSON["dependencies"].(map[string]any)
	devDependenciesMerge, foundDevDeps := packageJSON["devDependencies"].(map[string]any)
	// Merge dependencies and devDependencies -> devDependenciesMerge
	maps.Copy(devDependenciesMerge, dependencies)

	if !foundDeps && !foundDevDeps {
		return NewEvaluation(
				evalName,
				evalDesc,
				0,
				100,
				0,
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
		evalMessages = append(evalMessages, "No icon library found. Consider adding one for consistent icon usage.\n")
	case 1:
		score = 100
		evalMessages = append(evalMessages, fmt.Sprintf(
			"Using a single icon library: %s, which is ideal.\n",
			c.InfoFgBold(foundIconsLibs[0])),
		)
	default:
		score = 50
		evalMessages = append(evalMessages, fmt.Sprintf(
			"%d icon libraries found in package.json. Consider using a single icon library for consistent icon usage.\n",
			foundLibsCount),
		)
		for _, lib := range foundIconsLibs {
			evalMessages = append(evalMessages, fmt.Sprintf("- %s\n", lib))
		}
	}

	return NewEvaluation(evalName, evalDesc, score, 0, 0, evalMessages), nil
}
