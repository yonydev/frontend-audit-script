package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
	"github.com/yonydev/frontend-audit-script/evaluators"
	"github.com/yonydev/frontend-audit-script/models"
	"github.com/yonydev/frontend-audit-script/readers"
	"github.com/yonydev/frontend-audit-script/utils"
)

var (
	packageJSONContent string
	frontendFiles      []string
	assetsFiles        []string
	stylesFiles        []string
	evaluations        []models.Evaluation
)

func main() {
	color.NoColor = false

	dir, _ := os.Getwd()
	walking_directory_err := filepath.WalkDir(dir, walkDirFunc)

	if walking_directory_err != nil {
		panic(walking_directory_err)
	}

	if len(frontendFiles) > 0 {
		themeProvidersEvaluation, _ := evaluators.EvalThemeProviders(frontendFiles)
		webFontsEvaluation, _ := evaluators.EvalWebFonts(stylesFiles)

		evaluations = append(evaluations, themeProvidersEvaluation, webFontsEvaluation)
		// Theme Providers Evaluation
		fmt.Printf(
			"%s%s%v\n",
			themeProvidersEvaluation.Name,
			themeProvidersEvaluation.Description,
			utils.MapMessagePrinter(themeProvidersEvaluation.Messages),
			// themeProvidersEvaluation.Score,
		)
		// Web Fonts Evaluation
		fmt.Printf(
			"%s%s%v\n",
			webFontsEvaluation.Name,
			webFontsEvaluation.Description,
			utils.MapMessagePrinter(webFontsEvaluation.Messages),
			// webFontsEvaluation.Score,
		)
	} else {
		fmt.Println("No .js, .jsx, .ts, .tsx files found")
	}

	if len(assetsFiles) > 0 {
		assetsEvaluation, _ := evaluators.EvalAssets(assetsFiles)

		evaluations = append(evaluations, assetsEvaluation)
		// Assets Evaluation
		fmt.Printf(
			"%s%s%v\n",
			assetsEvaluation.Name,
			assetsEvaluation.Description,
			utils.MapMessagePrinter(assetsEvaluation.Messages),
			// assetsEvaluation.Score,
		)
	} else {
		fmt.Println("No image assets found")
	}

	evaluators.CalculateScore(evaluations)
}

func walkDirFunc(path string, d fs.DirEntry, err error) error {
	fileName := d.Name()
	isDir := d.IsDir()

	if err != nil {
		fmt.Printf("Error encountered: %v\n", err)
		return err
	}

	if isDir && utils.IgnoredDirsAndFiles[fileName] {
		return fs.SkipDir
	}

	if fileName == "package.json" {
		packageJSONContent = readers.FileReader(&path)

		reactEvaluation, _ := evaluators.EvalReactVersion(&packageJSONContent)
		iconsEvaluation, _ := evaluators.EvalIconLibs(&packageJSONContent)
		muiExtraLibsEvaluation, _ := evaluators.EvalMuiExtraLibs(&packageJSONContent)
		stylingLibsEvaluation, _ := evaluators.EvalStylingLibs(&packageJSONContent)

		evaluations = append(evaluations, reactEvaluation, iconsEvaluation, muiExtraLibsEvaluation, stylingLibsEvaluation)

		// React Version Evaluation
		fmt.Printf(
			"%s%s%v\n",
			reactEvaluation.Name,
			reactEvaluation.Description,
			utils.MapMessagePrinter(reactEvaluation.Messages),
			// reactEvaluation.Score,
		)

		// Icons Libs Evaluation
		fmt.Printf(
			"%s%s%v\n",
			iconsEvaluation.Name,
			iconsEvaluation.Description,
			utils.MapMessagePrinter(iconsEvaluation.Messages),
			// iconsEvaluation.Score,
		)

		// Mui Extra Libs Evaluation
		fmt.Printf(
			"%s%s%v\n",
			muiExtraLibsEvaluation.Name,
			muiExtraLibsEvaluation.Description,
			utils.MapMessagePrinter(muiExtraLibsEvaluation.Messages),
			// muiExtraLibsEvaluation.Score,
		)

		// Styling Libs Evaluation
		fmt.Printf(
			"%s%s%v\n",
			stylingLibsEvaluation.Name,
			stylingLibsEvaluation.Description,
			utils.MapMessagePrinter(stylingLibsEvaluation.Messages),
			// stylingLibsEvaluation.Score,
		)

	}

	if matchedFrontendFiles, _ := regexp.MatchString(`\.(js|jsx|ts|tsx)$`, path); matchedFrontendFiles {
		if !regexp.MustCompile(`\.(spec|specs|test|tests|story)\.(js|jsx|ts|tsx)$`).MatchString(path) {
			frontendFiles = append(frontendFiles, path)
		}
	}

	if matchedAssetsFiles, _ := regexp.MatchString(`\.(jpg|jpeg|png|gif|svg|webp)$`, path); matchedAssetsFiles {
		assetsFiles = append(assetsFiles, path)
	}

	if matchedStylesFiles, _ := regexp.MatchString(`\.(css|scss|sass|less|html)$`, path); matchedStylesFiles {
		stylesFiles = append(stylesFiles, path)
	}

	return nil
}
