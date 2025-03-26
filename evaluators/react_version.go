package evaluators

import (
	"encoding/json"
	"fmt"
	"regexp"

	c "github.com/yonydev/frontend-audit-script/colorize"
)

var (
	evalName = ">>> Check React Version\n"
	evalDesc = "Checking for React dev dependency...\n"
)

func EvalReactVersion(content *string) (Evaluation, error) {
	var packageJSON map[string]any

  score := 0
  minScore := 50
  maxScore := 100
  weight := 3

	if err := json.Unmarshal([]byte(*content), &packageJSON); err != nil {
		return Evaluation{}, fmt.Errorf("failed to parse package.json: %v", err)
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

func evaluateReactVersion(version int) Evaluation {
	var score int
	var evalMessages []string

	if version == 17 || version == 18 {
		score = 100
		evalMessages = append(evalMessages, fmt.Sprintf(
			"Using React version %s. React version supported.",
			c.InfoFgBold(version)),
		)
	} else if version < 17 {
		score = 50
		evalMessages = append(evalMessages, fmt.Sprintf(
			"Using React version %s. Consider upgrading to version 17 or 18 for better performance and features.",
			c.InfoFgBold(version),
		))
	} else {
		score = 30
		evalMessages = append(evalMessages, fmt.Sprintf(
			"Using React version %s. This is a future version. Ensure compatibility with your other dependencies.",
			c.InfoFgBold(version)),
		)
	}

	return NewEvaluation(
		evalName,
		evalDesc,
		score,
		100,
		50,
		0,
		evalMessages,
	)
}
