package evaluators

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	c "github.com/yonydev/frontend-audit-script/colorize"
)

func EvalThemeProviders(paths []string) (Evaluation, error) {
	var filesUsingThemeProvider []string
	var messages []string
	tagPattern := regexp.MustCompile(`<(ThemeProvider|ThemeProviders|MuiThemeProvider|UIThemeProvider)\b[^>]*?>`)

	evalName := ">>> Theme Provider Check\n"
	evalDesc := "Checking for theme provider components in files...\n"

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return Evaluation{}, fmt.Errorf("failed to open file: %s %v", path, err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		buf := make([]byte, 64*1024)   // 64KB buffer
		scanner.Buffer(buf, 1024*1024) // Increase buffer size to 1MB

		if len(buf) > 1024*1024 {
			panic("Buffer size is too large, it exceeds 1MB")
		}

		for scanner.Scan() {
			line := scanner.Text()
			if tagPattern.MatchString(line) {
				filesUsingThemeProvider = append(filesUsingThemeProvider, path)
				break
			}
		}

		if err := scanner.Err(); err != nil {
			return Evaluation{}, fmt.Errorf("error reading file %s: %v", err, path)
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
			messages,
		),
		nil
}
