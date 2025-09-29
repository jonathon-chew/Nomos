package rules

type Language struct {
	SingleLineComment     string
	Return                string
	VariableDeclaration   []string
	Function              string
	VariableAssigning     []string
	MultiLineCommentOpen  string
	MultiLineCommentClose string
}

func GetLanguage(language string) Language {

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
		return Go
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
		return Go
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
		return Powershell

	default:
		return Language{}
	}
}
