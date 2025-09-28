package main

import (
	"fmt"
	"os"

	"github.com/jonathon-chew/Nomos/cmd"
	"github.com/jonathon-chew/Nomos/rules"

	aphrodite "github.com/jonathon-chew/Aphrodite"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please enter in some arguments to process")
		return
	}

	for _, argument := range os.Args[1:] {

		_, doesExist := os.Stat(argument)
		if doesExist != nil {
			cmd.Cmd([]string{argument})
			continue
		}

		fmt.Printf("You asked for: %s\n", argument)
		fileRules, err := rules.Parse_rules()
		if err != nil {
			aphrodite.PrintError(fmt.Sprintf("error with the rules: %v\n", err))
			return
		}

		fileContents, err := get_file_contents(argument)
		if err != nil {
			aphrodite.PrintError(fmt.Sprintf("error with the opening and getting the file contents\n error was: %v\n", err))
			return
		}

		fileError := process_file(fileContents, fileRules)
		if fileError != nil {
			aphrodite.PrintError(fmt.Sprintf("error with the processing the file: %v\n", err))
			return
		}
	}
}
