package readers

import (
	"fmt"
	"os"
)

func PackageJSONReader(file *string) string {
	content, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println("Error reading package.json file")
	}
	return string(content)
}
