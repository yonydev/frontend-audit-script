package evaluators

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"

	c "github.com/yonydev/frontend-audit-script/colorize"
)

func EvalThemeProviders(paths []string) (Evaluation, error) {
	var filesUsingThemeProvider []string
	var messages []string
	tagPattern := regexp.MustCompile(`<(ThemeProvider|ThemeProviders|MuiThemeProvider|UIThemeProvider)\b[^>]*?>`)

	evalName := ">>> Theme Provider Check\n"
	evalDesc := "Checking for theme provider components in files...\n"

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
				if tagPattern.MatchString(line) {
					filesUsingThemeProvider = append(filesUsingThemeProvider, path)
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
			return Evaluation{}, res.err
		}
	}

	if len(filesUsingThemeProvider) == 0 {
		messages = append(
			messages,
			"\nNo ThemeProvider found in any of the .js(x) or .ts(x) files. Consider using a theme provider for consistent theme usage.",
		)
	} else {
		messages = append(
			messages,
			fmt.Sprintf(
				"\nTotal of %s files found theme provider components in the following files:",
				c.InfoFgBold(len(filesUsingThemeProvider)),
			),
		)
		for _, file := range filesUsingThemeProvider {
			messages = append(messages, fmt.Sprintf("file: %s", c.WarningFg(file)))
		}
	}

	return NewEvaluation(
			evalName,
			evalDesc,
			0,
			0,
			0,
			0,
			messages,
		),
		nil
}
