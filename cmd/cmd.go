package cmd

import (
	"fmt"
	"os"
	"reflect"

	aphrodite "github.com/jonathon-chew/Aphrodite"
	"github.com/jonathon-chew/Nomos/readme"
	"github.com/jonathon-chew/Nomos/rules"
)

// Parse and action arguments passed in that AREN'T files
func Command_parse(commandArguments []string) {

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

			_, errWrite := filePointer.Write(fileContents)
			if errWrite != nil {
				aphrodite.PrintError("unable to write the rules file: \n")
				return
			}

		case "--read-me":
			fileRules, err := rules.Parse_rules()
			if err != nil {
				aphrodite.PrintError(fmt.Sprintf("error with the rules: %v\n", err))
				return
			}
			if fileRules.ReadmeFile || fileRules.ReadmeStats {
				if readme.Check_for_README() {
					if fileRules.ReadmeStats {
						fileContents, err := os.ReadFile("./README.md")
						if err != nil {
							aphrodite.PrintError("error opening and reading the file!")
							return
						}
						readme.Stats(string(fileContents))
					}

					if fileRules.ReadmeFile {
						fmt.Println("Properly contains a README")
					}
				} else {
					aphrodite.PrintError("Could not find a README")
				}
			}

		case "--help":
			r := reflect.TypeOf(rules.Rules{})

			fmt.Println("Available options:")
			for i := 0; i < r.NumField(); i++ {
				field := r.Field(i)
				name := field.Name
				jsonName := field.Tag.Get("json")
				kind := field.Type.Kind().String()

				// fallback if json tag is empty
				if jsonName == "" {
					jsonName = name
				}

				fmt.Printf("  %s  %s\n", aphrodite.ReturnInfo(kind), jsonName)
			}

		default:
			aphrodite.PrintError(fmt.Sprintf("unrecognised argument: %s\n", argument))
			continue
		}
	}
}
