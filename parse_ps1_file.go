package main

import (
	"fmt"
	"slices"
	"strings"

	aphrodite "github.com/jonathon-chew/Aphrodite"
	"github.com/jonathon-chew/Nomos/rules"
)

/*
This is the entry point of the function
Currently this processes the file in a variety of ways regardless of the user input
This SHOULD in future ignore function checking IF there is no function check option even in the rules.json
However, I don't want to have to check if the rules exist at each byte, I also don't want to have to process the file multiplep times if I don't have to
*/
func process_ps1_file(fileBytes []byte, fileRules rules.Rules) error {

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
		if index+1 < len(fileBytes) && fileByte == '<' && fileBytes[index+1] == '#' {
			// Local slice for *this* comment
			var thisCommentLines []int
			thisCommentLines = append(thisCommentLines, lineNumber) // start line
			commentLine := lineNumber

			index += 2 // skip "/*"

			for index+1 < len(fileBytes) && !(fileBytes[index] == '#' && fileBytes[index+1] == '>') { // until the closing bytes
				b := fileBytes[index]
				if b == '\n' {
					commentLine++
					thisCommentLines = append(thisCommentLines, commentLine)
				}
				index++
			}

			// Skip closing "*/"
			if index+1 < len(fileBytes) && fileBytes[index] == '#' && fileBytes[index+1] == '>' {
				index += 2
			}

			// Now we have the full set of lines for this comment
			// log.Printf("Comment on lines: %v\n", thisCommentLines)
			commentLines = append(commentLines, thisCommentLines...)
		}

		/*
			This is for dealing with comment lines
		*/
		if index+1 < len(fileBytes) && fileByte == '#' && is_white_space(fileBytes[index+1]) {
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
		if string(combineBytes) == "function" && fileBytes[index+1] == ' ' {

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
			if fileRules.FunctionNames != "" && fileRules.FunctionNames != "ignore" {
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
		if index+1 < len(fileBytes) && string(combineBytes) == "=" && (strings.Contains(previousWord, "$") || slices.Contains(combineBytes, '$')) {

			switch {
			case previousWord[0] == '$':
				variable_name = previousWord
			case strings.Contains(previousWord, "$") && previousWord[0] != '$':
				for index, char := range previousWord {
					if char == '$' {
						variable_name = previousWord[index:]
						break
					}
				}
			default:
				continue
			}

			// Check for the case naming convention
			if fileRules.VariableNames != "ignore" && fileRules.VariableNames != "" && len(variable_name) > 2 {
				nameing_convention(variable_name[1:], fileRules.VariableNames, "Variable", fileRules)
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
			Dealing with key words
		*/
		if string(combineBytes) == "return" && fileBytes[index+1] == '\n' {
			if fileRules.NoNakedReturns {
				aphrodite.PrintWarning(fmt.Sprintf("There is a naked return on line %d\n", lineNumber))
			}
		}
	}
	return nil
}
