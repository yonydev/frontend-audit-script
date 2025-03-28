package writers

import (
	"strconv"

	"github.com/sethvargo/go-githubactions"
	"github.com/yonydev/frontend-audit-script/models"
)

func SetEvaluationEnvVariables(evaluation models.Evaluation, envVars map[string]string) {
	evalMap := map[string]string{
		"Name":     evaluation.Name,
		"Score":    strconv.Itoa(evaluation.Score),
		"MaxScore": strconv.Itoa(evaluation.MaxScore),
		"MinScore": strconv.Itoa(evaluation.MinScore),
		"Weight":   strconv.Itoa(evaluation.Weight),
	}

	for envKey, field := range envVars {
		if value, exists := evalMap[field]; exists {
			githubactions.SetEnv(envKey, value)
			// fmt.Printf("✅ Set %s=%s\n", envKey, value) // Debug output
		} else {
			githubactions.Warningf("Couldn't set the value of field '%s' for environment variable key '%s'", field, envKey)
		}
	}
}
