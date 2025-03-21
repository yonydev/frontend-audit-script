package evaluators

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/utils"
)

func EvalMuiExtraLibs(content *string) (Evaluation, error) {
	var packageJSON map[string]any
	var messages []string
	var foundLibs []string

	evalName := "\n>>> MUI Extra Libraries\n"
	evalDesc := "Checking for MUI extra libraries...\n"

	if err := json.Unmarshal([]byte(*content), &packageJSON); err != nil {
		return Evaluation{}, fmt.Errorf("failed to parse package.json: %v", err)
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

	if len(foundLibs) == 0 {
		messages = append(
			messages,
			c.SuccessFg("No MUI extra libraries found. Nice, keep it up! 🦾, keep it clean! 🧹"),
		)
	} else {
		messages = append(
			messages,
			fmt.Sprintf(
				"Found %s in package.json.\n",
				c.InfoFgBold(strings.Join(foundLibs, ", ")),
			),
		)
	}

	return NewEvaluation(
			evalName,
			evalDesc,
			0,
			0,
			0,
			messages,
		),
		nil
}
