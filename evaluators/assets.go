package evaluators

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"

	_ "golang.org/x/image/webp"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/utils"
)

func EvalAssets(paths []string) (Evaluation, error) {
	var moderateAssetsToOptimize []string
	var criticalAssetsToOptimize []string
	var messages []string

	evalName := "\n>>> Assets Optimization Check\n"
	evalDesc := "Looking for Checking for .jpg, .jpeg, .png, .gif, .svg, .webp files..."

	if len(paths) == 0 {
		messages = append(messages, "\nNo assets found in project.")
	} else {
		messages = append(messages, fmt.Sprintf(
			"\nTotal of %s assets found in project.",
			c.InfoFgBold(len(paths)),
		))
	}

	const maxFileSizeThreshold = 1024 * 1024  // 1MB
	const idealFileSizeThreshold = 200 * 1024 // 200KB

	type result struct {
		path   string
		sizeMB float64
		width  int
		height int
		err    error
	}

	results := make(chan result, len(paths))
	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			file, err := os.Stat(path)
			if err != nil {
				results <- result{path: path, err: err}
				return
			}

			fileSizeMB, _ := utils.ConvertSize(file.Size(), "B", "MB")
			width, height := 0, 0

			if ext := strings.ToLower(filepath.Ext(path)); ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".webp" {
				width, height, err = utils.GetImageDimensions(path)
				if err != nil {
					results <- result{path: path, err: err}
					return
				}
			}

			results <- result{path: path, sizeMB: fileSizeMB, width: width, height: height}
		}(path)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.err != nil {
			panic(res.err)
		}

		file, _ := os.Stat(res.path)
		if file.Size() > maxFileSizeThreshold {
			criticalAssetsToOptimize = append(criticalAssetsToOptimize, res.path)
		} else if file.Size() >= idealFileSizeThreshold && file.Size() <= maxFileSizeThreshold {
			moderateAssetsToOptimize = append(moderateAssetsToOptimize, res.path)
		}
	}

	if len(criticalAssetsToOptimize) > 0 {
		messages = append(messages, fmt.Sprintf(
			"\nTotal of %s critical assets to optimize found with size greater than 1MB, consider optimizing them for better score:",
			c.ErrorFgBold(len(criticalAssetsToOptimize)),
		))
		for _, asset := range criticalAssetsToOptimize {
			file, _ := os.Stat(asset)
			sizeMB, _ := utils.ConvertSize(file.Size(), "B", "MB")
			width, height, _ := utils.GetImageDimensions(asset)
			messages = append(messages, fmt.Sprintf(
				"asset: %s %.2f MB, dimensions: %sx%s",
				c.WarningFg(asset), sizeMB, c.InfoFg(width), c.InfoFg(height),
			))
		}
	}

	if len(moderateAssetsToOptimize) > 0 {
		messages = append(messages, fmt.Sprintf(
			"\nTotal of %s moderate assets to optimize found with size between 200KB to 1MB",
			c.WarningFgBold(len(moderateAssetsToOptimize)),
		))
		for _, asset := range moderateAssetsToOptimize {
			file, _ := os.Stat(asset)
			sizeMB, _ := utils.ConvertSize(file.Size(), "B", "MB")
			width, height, _ := utils.GetImageDimensions(asset)
			messages = append(messages, fmt.Sprintf(
				"asset: %s %.2f MB, dimemsions: %sx%s",
				c.WarningFg(asset), sizeMB, c.InfoFg(width), c.InfoFg(height),
			))
		}
	}

	if len(criticalAssetsToOptimize) == 0 && len(moderateAssetsToOptimize) == 0 {
		messages = append(messages, c.SuccessFg("No assets found to optimize. Keep up the good work!"))
	}

	return NewEvaluation(
			evalName,
			evalDesc,
			0,
			0,
			0,
			3,
			messages,
		),
		nil
}
