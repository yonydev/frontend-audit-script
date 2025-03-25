package evaluators

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/utils"
)

func EvalWebFonts(paths []string) (Evaluation, error) {
	var messages []string
	var mu sync.Mutex
	var webFonts []string
	var filesWithWebFonts []string

	filesWithWebFontsSet := make(map[string]struct{}) // Use a map to avoid duplicates
	webFontsSet := make(map[string]struct{})          // Use a map to avoid duplicates

	fontFamilyRegex := regexp.MustCompile(`font-family:\s*['"]?([^'",\s;]+)`)
	googleFontsImportsRegex := regexp.MustCompile(`https://fonts.googleapis.com/css\?family=([^&"']+)`)
	googleFontsLinksRegex := regexp.MustCompile(`<link[^>]+href=["']https://fonts.googleapis.com/css\?family=([^:"'&,]+)`)

	evalName := "\n>>> Web Fonts Check\n"
	evalDesc := "Checking for web fonts in .css, .scss, .sass files...\n"

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
				if matches := fontFamilyRegex.FindStringSubmatch(line); len(matches) > 1 {
					if !utils.IsGenericFontFamily(matches[1]) {
						mu.Lock()
						webFontsSet[matches[1]] = struct{}{}
						filesWithWebFontsSet[path] = struct{}{}
						mu.Unlock()
					}
				}
				if matches := googleFontsImportsRegex.FindStringSubmatch(line); len(matches) > 1 {
					mu.Lock()
					fontName := strings.Split(strings.ReplaceAll(matches[1], "+", " "), ":")[0]
					webFontsSet[fontName] = struct{}{}
					filesWithWebFontsSet[path] = struct{}{}
					mu.Unlock()
				}
				if matches := googleFontsLinksRegex.FindStringSubmatch(line); len(matches) > 1 {
					mu.Lock()
					fontName := strings.Split(strings.ReplaceAll(matches[1], "+", " "), ":")[0]
					webFontsSet[fontName] = struct{}{}
					filesWithWebFontsSet[path] = struct{}{}
					mu.Unlock()
				}
			}

			if err := scanner.Err(); err != nil {
				results <- result{path: path, err: fmt.Errorf("error reading file %s: %v", path, err)}
				return
			}
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

	// Append webFontsSet to normal slice
	for font := range webFontsSet {
		webFonts = append(webFonts, font)
	}
	// Append filesWithWebFontsSet to normal slice
	for file := range filesWithWebFontsSet {
		filesWithWebFonts = append(filesWithWebFonts, file)
	}

	if len(webFontsSet) == 0 {
		messages = append(
			messages,
			c.WarningFg("No web fonts found in the project. Consider using system fonts or generic font families."),
		)
	} else {
		messages = append(
			messages,
			fmt.Sprintf("Total of %s fonts used in %s files", c.InfoFgBold(len(webFontsSet)), c.InfoFgBold(len(filesWithWebFontsSet))),
			fmt.Sprintf("Fonts used: %s", c.InfoFgBold(strings.Join(webFonts, ", "))),
		)
		for _, file := range filesWithWebFonts {
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
