package cmd

import (
	"fmt"
	"os"

	aphrodite "github.com/jonathon-chew/Aphrodite"
)

func Cmd(commandArguments []string) {

	for _, argument := range commandArguments {
		switch argument {
		case "--make-default":
			fileContents := []byte(`{
				"functions-have-doc-strings": true,
				"variable-names": "camel_case",
				"function-names": "snake_case",
				"print-f-new-line": true,
				"only-show-errors": true,
				"ignore-if-in-comments": true,
				"exported-identifiers-have-comments": true,
				"const-in-caps": true,
				"no-naked-returns": true
			}`)

			filePointer, err := os.Create("./nomos_rules.json")
			if err != nil {
				aphrodite.PrintError("unable to create the rules file: \n")
				return
			}

			_, ErrWrite := filePointer.Write(fileContents)
			if ErrWrite != nil {
				aphrodite.PrintError("unable to write the rules file: \n")
				return
			}

		default:
			aphrodite.PrintError(fmt.Sprintf("unrecognised argument: %s\n", argument))
			continue
		}
	}
}
