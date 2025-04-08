package evaluators

import (
	"encoding/json"
	"fmt"
	"regexp"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/models"
	"github.com/yonydev/frontend-audit-script/utils"
	"github.com/yonydev/frontend-audit-script/writers"
)

var (
	evalName = ">>> Check React Version"
	evalDesc = "\nChecking for React dev dependency...\n"
)

func EvalReactVersion(content *string) (models.Evaluation, error) {
	var packageJSON map[string]any

	score := 0
	minScore := -2
	maxScore := 2
	weight := 1

	if err := json.Unmarshal([]byte(*content), &packageJSON); err != nil {
		return models.Evaluation{}, fmt.Errorf("failed to parse package.json: %v", err)
	}

	dependencies, found := packageJSON["dependencies"].(map[string]any)
	if !found {
		score = minScore
		return NewEvaluation(
				evalName,
				evalDesc,
				score,
				maxScore,
				minScore,
				weight,
				[]string{c.WarningFg("React dependency not found in package.json")},
			),
			nil
	}

	reactVersion, found := dependencies["react"].(string)
	if !found {
		score = minScore
		return NewEvaluation(
				evalName,
				evalDesc,
				score,
				maxScore,
				minScore,
				weight,
				[]string{c.WarningFg("React dependency not found in package.json")},
			),
			nil
	}

	majorVersion := extractMajorVersion(reactVersion)

	return evaluateReactVersion(majorVersion), nil
}

func extractMajorVersion(version string) int {
	// Remove special characters (^, ~, *) and keep only numbers before the first dot
	re := regexp.MustCompile(`^[^0-9]*([0-9]+)\..*`)
	matches := re.FindStringSubmatch(version)
	if len(matches) > 1 {
		var major int
		fmt.Sscanf(matches[1], "%d", &major)
		return major
	}
	return 0
}

func evaluateReactVersion(version int) models.Evaluation {
	var score int
	var evalMessages []string

	maxScore := 2
	minScore := -2
	weight := 3

	if version == 17 || version == 18 {
		score = 2
		evalMessages = append(evalMessages, fmt.Sprintf(
			"Using React version %s. React version supported.",
			c.InfoFgBold(version)),
		)
		// } else if version < 17 {
		// 	score = 50
		// 	evalMessages = append(evalMessages, fmt.Sprintf(
		// 		"Using React version %s. Consider upgrading to version 17 or 18 for better performance and features.",
		// 		c.InfoFgBold(version),
		// 	))
	} else {
		score = -2
		evalMessages = append(evalMessages, fmt.Sprintf(
			"Using React version %s. This is a future version. Ensure compatibility with your other dependencies.",
			c.InfoFgBold(version)),
		)
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
	writers.SetEvaluationEnvVariables(evaluation, utils.ReactVersionEnvVars)

	return evaluation
}
