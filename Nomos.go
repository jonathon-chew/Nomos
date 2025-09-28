package main

import (
	"fmt"
	"os"

	"github.com/jonathon-chew/Nomos/readme"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please enter in some arguments to process")
		return
	}

	for _, argument := range os.Args[1:] {
		fmt.Printf("You asked for: %s\n", argument)
		fileRules, err := parse_rules()
		if err != nil {
			fmt.Printf("error with the rules: %v\n", err)
			return
		}

		fileError := process_file(argument, fileRules)
		if fileError != nil {
			fmt.Printf("error with the processing the file: %v\n", err)
			return
		}

		if fileRules.ReadmeFile {
			if readme.Check_for_README() {
				fmt.Println("Properly contains a README")
				if fileRules.ReadmeStats {
					fileContents, err := os.ReadFile("./README.md")
					if err != nil {
						return
					}
					readme.Stats(string(fileContents))
				}
			} else {
				fmt.Println("Could not find a README")
			}
		}
	}
}
