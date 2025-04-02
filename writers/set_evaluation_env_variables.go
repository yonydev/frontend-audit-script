package writers

import (
	"strconv"

	"github.com/sethvargo/go-githubactions"
	"github.com/yonydev/frontend-audit-script/models"
)

func SetEvaluationEnvVariables(evaluation models.Evaluation, envVars map[string]string) {
	action := githubactions.New()
	evalMap := map[string]string{
		"Name":     evaluation.Name,
		"Score":    strconv.Itoa(evaluation.Score),
		"MaxScore": strconv.Itoa(evaluation.MaxScore),
		"MinScore": strconv.Itoa(evaluation.MinScore),
		"Weight":   strconv.Itoa(evaluation.Weight),
	}

	for envKey, field := range envVars {
		if value, exists := evalMap[field]; exists {
			action.SetEnv(envKey, value)
			// fmt.Printf("✅ Set %s=%s\n", envKey, value) // Debug output
		} else {
			action.Warningf("Couldn't set the value of field '%s' for environment variable key '%s'", field, envKey)
		}
	}
}
