package rules

import "errors"

type Language struct {
	SingleLineComment     string
	Return                string
	VariableDeclaration   []string
	Function              string
	VariableAssigning     []string
	MultiLineCommentOpen  string
	MultiLineCommentClose string
}

func GetLanguage(language string) (Language, error) {

	switch language {
	case "go":
		Go := Language{
			SingleLineComment:     "//",
			Return:                "return",
			VariableDeclaration:   []string{"var", "const"},
			Function:              "func",
			VariableAssigning:     []string{":=", "="},
			MultiLineCommentOpen:  "/*",
			MultiLineCommentClose: "*/",
		}
		return Go, nil
	case "golang":
		Go := Language{
			SingleLineComment:     "//",
			Return:                "return",
			VariableDeclaration:   []string{"var", "const"},
			Function:              "func",
			VariableAssigning:     []string{":=", "="},
			MultiLineCommentOpen:  "/*",
			MultiLineCommentClose: "*/",
		}
		return Go, nil
	case "powershell":
		Powershell := Language{
			SingleLineComment:     "#",
			Return:                "return",
			VariableDeclaration:   []string{"$"},
			Function:              "function",
			VariableAssigning:     []string{"="},
			MultiLineCommentOpen:  "<#",
			MultiLineCommentClose: "#>",
		}
		return Powershell, nil
	case "python":
		Python := Language{
			SingleLineComment:     "#",
			Return:                "return",
			VariableDeclaration:   []string{},
			Function:              "def",
			VariableAssigning:     []string{"="},
			MultiLineCommentOpen:  "'''",
			MultiLineCommentClose: "'''",
		}
		return Python, nil

	default:
		return Language{}, errors.New("language not recognised")
	}
}
