package rules

import (
	"encoding/json"
	"errors"
	"os"
)

/*
Convert the file json into a go struct which can be used as switches to activate different parts of the code
*/
type Rules struct {
	FunctionDocStrings              bool   `json:"functions-have-doc-strings"`
	IgnoreMainFunction              bool   `json:"ignore-main-function"`
	VariableNames                   string `json:"variable-names"`
	FunctionNames                   string `json:"function-names"`
	ReadmeFile                      bool   `json:"readme-file"`
	ReadmeStats                     bool   `json:"readme-stats"`
	SideComments                    bool   `json:"side-comments"`
	PrintFNewLine                   bool   `json:"print-f-new-line"`
	IgnoreIfInComments              bool   `json:"ignore-if-in-comments"`
	OnlyShowErrors                  bool   `json:"only-show-errors"`
	ListInternalFunctions           bool   `json:"list-internal-functions"`
	ExportedIdentifiersHaveComments bool   `json:"exported-identifiers-have-comments"`
	ConstInCaps                     bool   `json:"const-in-caps"`
	NoNakedReturns                  bool   `json:"no-naked-returns"`
	ParseStatRules                  bool   `json:"show-stat-rules"`
}

/*
This converts the json file from a file to a go struct Rules
*/
func Parse_rules() (Rules, error) {
	var rulesJson string = "./nomos_rules.json"
	var fileRules Rules

	_, doesExist := os.Stat(rulesJson)
	if doesExist != nil {
		return Rules{}, errors.New("nomos_rules.json file does not exist, please create one with the options you'd like turned on")
	}

	// read the rules into memory

	fileBytes, err := os.ReadFile(rulesJson)
	if err != nil {
		return Rules{}, err
	}

	jsonErr := json.Unmarshal(fileBytes, &fileRules)
	if jsonErr != nil {
		return Rules{}, jsonErr
	}

	if fileRules.ExportedIdentifiersHaveComments {
		fileRules.FunctionDocStrings = true
	}

	return fileRules, nil
}
