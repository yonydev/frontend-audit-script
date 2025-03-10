package evaluators

import (
	"fmt"
	"os"

	c "github.com/yonydev/frontend-audit-script/colorize"
	"github.com/yonydev/frontend-audit-script/utils"
)

func EvalAssets(paths []string) (Evaluation, error) {
	var moderateAssetsToOptimize []string
	var criticalAssetsToOptimize []string
	var messages []string

	evalName := ">>> Assets Optimization Check\n"
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

	for _, path := range paths {
		file, err := os.Stat(path)
		if err != nil {
			panic(err)
		}

		fileSizeMB, _ := utils.ConvertSize(file.Size(), "B", "MB")
		if file.Size() > maxFileSizeThreshold {
			criticalAssetsToOptimize = append(criticalAssetsToOptimize, path)
			messages = append(messages, fmt.Sprintf(
				"\nTotal of %s critical assets to optimize found with size greater than 1MB, consider optimizing them for better score:",
				c.ErrorFgBold(len(criticalAssetsToOptimize)),
			))
			for _, asset := range criticalAssetsToOptimize {
				messages = append(messages, fmt.Sprintf(
					"asset: %s %.2f MB",
					asset, fileSizeMB,
				))
			}
		} else if file.Size() >= idealFileSizeThreshold && file.Size() <= maxFileSizeThreshold {
			moderateAssetsToOptimize = append(moderateAssetsToOptimize, path)
			messages = append(messages, fmt.Sprintf(
				"\nTotal of %s moderate assets to optimize found with size between 200KB to 1MB",
				c.WarningFgBold(len(moderateAssetsToOptimize)),
			))
			for _, asset := range moderateAssetsToOptimize {
				messages = append(messages, fmt.Sprintf("asset: %s %.2f MB", asset, fileSizeMB))
			}
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
			messages,
		),
		nil
}
