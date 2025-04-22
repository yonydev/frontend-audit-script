package writers

import (
	"fmt"
	"os"
	"strconv"

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

	githubOutputPath := os.Getenv("GITHUB_OUTPUT")
	var file *os.File
	var err error

	if githubOutputPath != "" {
		file, err = os.OpenFile(githubOutputPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("❌ Error opening GITHUB_ENV file: %v\n", err)
			file = nil
		}
	}

	for envKey, field := range envVars {
		if value, exists := evalMap[field]; exists {
			err := os.Setenv(envKey, value)
			if err != nil {
				fmt.Printf("⚠️ Error setting %s: %v\n", envKey, err)
			}
			if file != nil {
				_, err := file.WriteString(fmt.Sprintf("%s=%s\n", envKey, value))
				if err != nil {
					fmt.Printf("⚠️ Error writing %s to GITHUB_OUTPUT: %v\n", envKey, err)
				}
			}
		} else {
			fmt.Printf("⚠️ Couldn't set field '%s' for environment variable key '%s'\n", field, envKey)
		}
	}
	// Close file if it was opened
	if file != nil {
		file.Close()
	}
}
