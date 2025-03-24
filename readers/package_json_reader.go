package readers

import (
	"fmt"
	"os"
)

func FileReader(file *string) string {
	content, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println("Error reading file")
	}
	return string(content)
}
