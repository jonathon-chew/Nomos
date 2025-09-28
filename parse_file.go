package main

import (
	"fmt"
	"os"
	"strings"

	aphrodite "github.com/jonathon-chew/Aphrodite"
	"github.com/jonathon-chew/Nomos/rules"
)

// This is a simple way to check for the white spaces that I care about
func is_white_space(currentByte byte) bool {
	if currentByte == ' ' || currentByte == '\n' || currentByte == '\t' {
		return true
	}
	return false
}

/*
Implimenting a checker for naming conventions which can be shared between variables AND functions
*/
func nameing_convention(check_name, name_rule, fileType string, fileRules rules.Rules) bool {
	firstLetter := string(check_name[0])

	//(#1) TODO: Check length - anything bigger than 8 letters SURELY should fit, but something under might not work
	//(#2) TODO: Impliment simple fix? Offer the ability to do simple find and replace and match the new type -> Snake to Kebab "_" to "-"

	switch name_rule {
	case "camel_case":
		if strings.ToLower(firstLetter) == firstLetter && !strings.Contains(check_name, "_") && !strings.Contains(check_name, "-") {
			if !fileRules.OnlyShowErrors {
				aphrodite.PrintInfo(fmt.Sprintf("%s %s is camelCase and should be\n", fileType, check_name))
			}
			return true
		} else {
			aphrodite.PrintWarning(fmt.Sprintf("%s %s is not camelCase and should be\n", fileType, check_name))
			return false
		}

	case "snake_case":
		if strings.Contains(check_name, "_") && !strings.Contains(check_name, "-") {
			if !fileRules.OnlyShowErrors {
				aphrodite.PrintInfo(fmt.Sprintf("%s %s is snake_case and should be\n", fileType, check_name))
			}
			return true
		} else {
			aphrodite.PrintWarning(fmt.Sprintf("%s %s is not snake_case and should be\n", fileType, check_name))
			return false
		}

	case "kebab_case":
		if strings.Contains(check_name, "-") && !strings.Contains(check_name, "_") {
			if !fileRules.OnlyShowErrors {
				aphrodite.PrintInfo(fmt.Sprintf("%s %s is kebab-case and should be\n", fileType, check_name))
			}
			return true
		} else {
			aphrodite.PrintWarning(fmt.Sprintf("%s %s is not kebab-case and should be\n", fileType, check_name))
			return false
		}
	case "pascal_case":
		if firstLetter == strings.ToUpper(firstLetter) {
			if !fileRules.OnlyShowErrors {
				aphrodite.PrintInfo(fmt.Sprintf("%s %s is Pascal_case and should be\n", fileType, check_name))
			}
			return true
		} else {
			aphrodite.PrintWarning(fmt.Sprintf("%s %s is not Pascal_case and should be\n", fileType, check_name))
			return false
		}
	case "ignore":
		return false
	default:
		return false
	}
}

/*
A universal function to check if exported identifiers have a comment above to explain what they are a do in order to support LSP suport
*/
func check_for_doc_strings(commentLines []int, lineNumber int, identifierType, identifierName string, fileRules rules.Rules) {
	if len(commentLines) > 0 {
		if commentLines[len(commentLines)-1] == lineNumber-1 && !fileRules.OnlyShowErrors {
			aphrodite.PrintInfo(fmt.Sprintf("%s %s has a comment to explain it\n", identifierType, identifierName))
		} else {
			if !fileRules.OnlyShowErrors {
				aphrodite.PrintWarning(fmt.Sprintf("%s %s does not have a comment to explain it\n", identifierType, identifierName))
			}
		}
	} else {
		aphrodite.PrintWarning(fmt.Sprintf("%s %s does not have a comment to explain it\n", identifierType, identifierName))
	}
	// log.Println(commentLines, "\n")
}

// Check for is this inside a comment
func is_in_comment(line int, commentLines []int) bool {
	for _, l := range commentLines {
		if l == line {
			return true
		}
	}
	return false
}

func get_file_contents(fileName string) ([]byte, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		return []byte{}, err
	}

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return []byte{}, err
	}

	return fileBytes, nil
}

/*
This is the entry point of the function
Currently this processes the file in a variety of ways regardless of the user input
This SHOULD in future ignore function checking IF there is no function check option even in the rules.json
However, I don't want to have to check if the rules exist at each byte, I also don't want to have to process the file multiplep times if I don't have to
*/
func process_file(fileBytes []byte, fileRules rules.Rules) error {

	var combineBytes []byte
	var previousWord, variable_name, commentString string
	var lineNumber int = 1
	var commentLines []int

	for index, fileByte := range fileBytes {

		/*
			Dealing with spaces
		*/
		if is_white_space(fileByte) {
			previousWord = string(combineBytes)
			if fileByte == '\n' {
				lineNumber++
			}
			// log.Println("new word", previousWord)
			combineBytes = []byte{}
			continue
		}

		combineBytes = append(combineBytes, fileByte)

		/*
			This is for dealing with comment blocks
		*/
		if index+1 < len(fileBytes) && fileByte == '/' && fileBytes[index+1] == '*' {
			// Local slice for *this* comment
			var thisCommentLines []int
			thisCommentLines = append(thisCommentLines, lineNumber) // start line
			commentLine := lineNumber

			index += 2 // skip "/*"

			for index+1 < len(fileBytes) && !(fileBytes[index] == '*' && fileBytes[index+1] == '/') {
				b := fileBytes[index]
				if b == '\n' {
					commentLine++
					thisCommentLines = append(thisCommentLines, commentLine)
				}
				index++
			}

			// Skip closing "*/"
			if index+1 < len(fileBytes) && fileBytes[index] == '*' && fileBytes[index+1] == '/' {
				index += 2
			}

			// Now we have the full set of lines for this comment
			// log.Printf("Comment on lines: %v\n", thisCommentLines)
			commentLines = append(commentLines, thisCommentLines...)
		}

		/*
			This is for dealing with comment lines
		*/
		if fileByte == '/' && fileBytes[index+1] == '/' {
			for fileBytes[index] != '\n' && index+1 < len(fileBytes) {
				commentString += string(fileByte)
				index++
			}
			commentLines = append(commentLines, lineNumber)
		}

		/*
			Don't look any further if in comments rule is applied
		*/
		if fileRules.IgnoreIfInComments && is_in_comment(lineNumber, commentLines) {
			continue
		}

		/*
			Dealing with function declarations
		*/
		if string(combineBytes) == "func" && fileBytes[index+1] == ' ' {

			// Get function name
			var breaker bool = false
			var functionName string
			index += 2
			for !breaker {
				if is_white_space(fileBytes[index]) || fileBytes[index] == '(' {
					breaker = true
				}
				functionName += string(fileBytes[index])
				index++
			}
			functionName = functionName[:len(functionName)-1]
			// log.Printf("Found a func, on line, %d %s\n", lineNumber, functionName)

			// Check the form of the function
			if fileRules.FunctionDocStrings {
				nameing_convention(functionName, fileRules.FunctionNames, "Function", fileRules)
			}

			// Show only the internal functions
			if fileRules.ListInternalFunctions {
				firstLetter := string(functionName[0])
				if firstLetter == strings.ToLower(firstLetter) {
					aphrodite.PrintInfo(fmt.Sprintf("The function %s is an internal function only\n", functionName))
				}
			}

			//Check that the function has doc strings
			if fileRules.FunctionDocStrings {
				check_for_doc_strings(commentLines, lineNumber, "Function", functionName, fileRules) // if the first character isn't lower case, check for a doc string
			}
		}

		// This is for dealing with variable declarations
		if (string(combineBytes) == "var" && is_white_space(fileBytes[index+1])) || string(combineBytes) == "const" || string(combineBytes) == ":=" {

			// Get the variable name if it's the next thing declared
			if string(combineBytes) == "var" || string(combineBytes) == "const" {
				index += 1
				for fileBytes[index+1] != ' ' {
					variable_name += string(fileBytes[index+1])
					index++
				}
				// log.Printf("new declared variable %s %s\n", previousWord, variable_name)
			}

			var isConst bool = false
			if string(combineBytes) == "const" && fileRules.ConstInCaps {
				isConst = true
			}

			// Get the variable name if it's been declared right before / auto assigned type
			if string(combineBytes) == ":=" {
				variable_name = previousWord
				// log.Printf("new auto assign variable %s %s\n", variable_name, string(combineBytes))
			}

			// Check for const variables to be in CAPS
			if fileRules.ConstInCaps && isConst {
				if variable_name != strings.ToUpper(variable_name) {
					aphrodite.PrintWarning(fmt.Sprintf("Const variable %s isn't in caps lock\n", variable_name))
				} else {
					if !fileRules.OnlyShowErrors {
						aphrodite.PrintInfo(fmt.Sprintf("Const variable %s is in caps lock\n", variable_name))
					}
				}
			}

			// Check for the case naming convention
			if fileRules.VariableNames != "ignore" && fileRules.VariableNames != "" && !isConst {
				nameing_convention(variable_name, fileRules.VariableNames, "Variable", fileRules)
			}

			if fileRules.ExportedIdentifiersHaveComments {
				firstVariableLetter := string(variable_name[0])
				if firstVariableLetter != strings.ToLower(firstVariableLetter) {
					check_for_doc_strings(commentLines, lineNumber, "Variable", variable_name, fileRules) // if the first character isn't lower case, check for a doc string
				}
			}
			variable_name = ""
		}

		/*
			Dealing with types
		*/

		if string(combineBytes) == "type" {
			// Must be followed by space or tab
			if index+1 < len(fileBytes) && is_white_space(fileBytes[index+1]) {
				// Also make sure we’re not in a comment

				index++ // skip the space

				var typeName string
				for index+1 < len(fileBytes) && !is_white_space(fileBytes[index+1]) && fileBytes[index+1] != '{' {
					typeName += string(fileBytes[index+1])
					index++
				}

				if fileRules.ExportedIdentifiersHaveComments {
					first := string(typeName[0])
					if first != strings.ToLower(first) {
						check_for_doc_strings(commentLines, lineNumber, "Type", typeName, fileRules)
					}
				}

			}
		}

		/*
			Dealing with key words
		*/
		if string(combineBytes) == "return" && fileBytes[index+1] == '\n' {
			if fileRules.NoNakedReturns {
				aphrodite.PrintWarning(fmt.Sprintf("There is a naked return on line %d\n", lineNumber))
			}
		}

		/*
			Dealing with standard library conventions
		*/
		if fileRules.PrintFNewLine {
			if string(combineBytes) == "fmt.Printf(\"" {
				for fileBytes[index+1] != '"' {
					index++
				}
				if fileBytes[index] == 'n' && fileBytes[index-1] == '\\' {
					if !fileRules.OnlyShowErrors {
						aphrodite.PrintInfo(fmt.Sprintf("Printf statement ends with a new line character, %d\n", lineNumber))
					}
				} else {
					aphrodite.PrintWarning(fmt.Sprintf("Printf statement does not end with a new line character, %d\n", lineNumber))
				}
			}
		}
	}
	return nil
}
