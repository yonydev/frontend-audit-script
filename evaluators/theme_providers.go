package evaluators

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/models"
	"github.com/yonydev/frontend-audit-script/utils"
	"github.com/yonydev/frontend-audit-script/writers"
)

func EvalThemeProviders(paths []string) (models.Evaluation, error) {
	var filesUsingThemeProvider []string
	var themeProvidersNames []string
	var messages []string

	uniqueProviders := make(map[string]struct{})
	tagPattern := regexp.MustCompile(`<(ThemeProvider|ThemeProviders|MuiThemeProvider|UIThemeProvider)\b[^>]*?>`)

	evalName := ">>> Theme Provider Check"
	evalDesc := "\nChecking for theme provider components in files...\n"

	score := 3
	minScore := -2
	maxScore := 3
	weight := 2

	type result struct {
		path string
		err  error
	}

	results := make(chan result, len(paths))
	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			file, err := os.Open(path)
			if err != nil {
				results <- result{path: path, err: err}
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			buf := make([]byte, 64*1024)   // 64KB buffer
			scanner.Buffer(buf, 1024*1024) // Increase buffer size to 1MB

			if len(buf) > 1024*1024 {
				results <- result{path: path, err: fmt.Errorf("buffer size is too large, it exceeds 1MB")}
				return
			}

			for scanner.Scan() {
				line := scanner.Text()
				matches := tagPattern.FindStringSubmatch(line)
				if len(matches) > 1 {
					filesUsingThemeProvider = append(filesUsingThemeProvider, path)
					uniqueProviders[matches[1]] = struct{}{}
					break
				}
			}

			if err := scanner.Err(); err != nil {
				results <- result{path: path, err: fmt.Errorf("error reading file %s: %v", path, err)}
				return
			}

			results <- result{path: path, err: nil}
		}(path)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.err != nil {
			return models.Evaluation{}, res.err
		}
	}

	for themeName := range uniqueProviders {
		themeProvidersNames = append(themeProvidersNames, themeName)
	}

	numProviders := len(uniqueProviders)
	numFiles := len(filesUsingThemeProvider)

	if numProviders == 0 {
		messages = append(
			messages,
			"\nNo ThemeProvider found in any of the .js(x) or .ts(x) files. Consider using a theme provider for consistent theme usage.",
		)
	} else {
		messages = append(
			messages,
			fmt.Sprintf("Using: %s", c.InfoFgBold(strings.Join(themeProvidersNames, ", "))),
			fmt.Sprintf("\nTotal of %s files found with theme provider components:", c.InfoFgBold(numFiles)),
		)
		for _, file := range filesUsingThemeProvider {
			messages = append(messages, fmt.Sprintf("file: %s", c.WarningFg(file)))
		}
	}

	switch numProviders {
	case 1:
		score = maxScore
	case 2:
		score = 1
	default:
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

	writers.SetEvaluationEnvVariables(evaluation, utils.ThemeProvidersEnvVars)

	return evaluation, nil
}
