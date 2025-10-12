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

		case "--read-me", "--readme":
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

		case "--gitignore":
			_, ErrGitIgnore := os.Stat("./.gitignore")
			if ErrGitIgnore != nil {
				aphrodite.PrintError("Could not find a .gitignore")
				return
			}

			FileBytes, ErrGitIgnoreBytes := os.ReadFile("./.gitignore")
			if ErrGitIgnoreBytes != nil {
				aphrodite.PrintError("Could not read .gitignore")
				return
			}

			var rule_file_name string = "nomos_rules.json"
			var rule_file_added_already bool = false

			for file_index := range FileBytes {
				if string(FileBytes[file_index:file_index+len(rule_file_name)]) == rule_file_name {
					rule_file_added_already = true
				}
			}

			if !rule_file_added_already {
				new_contents := "\nnomos_rules.json"
				contents := append(FileBytes, new_contents...)
				os.WriteFile("./.gitignore", contents, os.ModeAppend)
			}

		case "--help", "-h":
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

				aphrodite.PrintInfo("--make-default\n This can be used to make a default file")
				aphrodite.PrintInfo("--gitignore\n This can be used to add the rules file automatically to the local .gitignore")
				aphrodite.PrintInfo("--read-me\n This is currently under developement for returning basic stats about a Readme file")
			}

		case "--version", "-v":
			versionNumber := "0.0.1"
			aphrodite.PrintInfo(versionNumber)
		default:
			aphrodite.PrintError(fmt.Sprintf("unrecognised argument: %s\n", argument))
			continue
		}
	}
}
