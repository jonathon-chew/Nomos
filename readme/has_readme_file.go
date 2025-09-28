package readme

import (
	"os"
	"strings"
)

func Check_for_README() bool {
	files, err := os.ReadDir(".")
	if err != nil {
		return false
	}

	for _, file := range files {
		if strings.ToLower(file.Name()) == "readme.md" {
			return true
		}
	}

	return false
}
