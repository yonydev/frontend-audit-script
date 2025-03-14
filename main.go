package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
	"github.com/yonydev/frontend-audit-script/evaluators"
	"github.com/yonydev/frontend-audit-script/readers"
	"github.com/yonydev/frontend-audit-script/utils"
)

var (
	packageJSONContent string
	frontendFiles      []string
	assetsFiles        []string
)

func main() {
	color.NoColor = false

	dir, _ := os.Getwd()
	walking_directory_err := filepath.WalkDir(dir, walkDirFunc)

	if walking_directory_err != nil {
		panic(walking_directory_err)
	}

	themeProvidersEvaluation, _ := evaluators.EvalThemeProviders(frontendFiles)
	fmt.Printf(
		"%s%s%v\n",
		themeProvidersEvaluation.Name,
		themeProvidersEvaluation.Description,
		utils.MapMessagePrinter(themeProvidersEvaluation.Messages),
		// themeProvidersEvaluation.Score,
	)

	assetsEvaluation, _ := evaluators.EvalAssets(assetsFiles)
	fmt.Printf(
		"%s%s%v\n",
		assetsEvaluation.Name,
		assetsEvaluation.Description,
		utils.MapMessagePrinter(assetsEvaluation.Messages),
		// assetsEvaluation.Score,
	)

	evaluators.EvalAssets(assetsFiles)
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
		packageJSONContent = readers.PackageJSONReader(&path)

		reactEvaluation, _ := evaluators.EvalReactVersion(&packageJSONContent)
		fmt.Printf(
			"%s%s%v\n",
			reactEvaluation.Name,
			reactEvaluation.Description,
			utils.MapMessagePrinter(reactEvaluation.Messages),
			// reactEvaluation.Score,
		)

		iconsEvaluation, _ := evaluators.EvalIconLibs(&packageJSONContent)
		fmt.Printf(
			"%s%s%v\n",
			iconsEvaluation.Name,
			iconsEvaluation.Description,
			utils.MapMessagePrinter(iconsEvaluation.Messages),
			// iconsEvaluation.Score,
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

	return nil
}
