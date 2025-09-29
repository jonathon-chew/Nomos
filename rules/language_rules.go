package rules

type Language struct {
	SingleLineComment     string
	Return                string
	VariableDeclaration   string
	Function              string
	VariableAssigning     string
	MultiLineCommentOpen  string
	MultiLineCommentClose string
}

func GetLanguage(language string) Language {

	switch language {
	case "go":
		Go := Language{
			SingleLineComment:     "//",
			Return:                "return",
			VariableDeclaration:   "return",
			Function:              "return",
			VariableAssigning:     "return",
			MultiLineCommentOpen:  "/*",
			MultiLineCommentClose: "*/",
		}
		return Go
	case "golang":
		Go := Language{
			SingleLineComment:     "//",
			Return:                "return",
			VariableDeclaration:   "return",
			Function:              "return",
			VariableAssigning:     "return",
			MultiLineCommentOpen:  "/*",
			MultiLineCommentClose: "*/",
		}
		return Go
	case "powershell":
		Go := Language{
			SingleLineComment:     "//",
			Return:                "return",
			VariableDeclaration:   "return",
			Function:              "return",
			VariableAssigning:     "return",
			MultiLineCommentOpen:  "/*",
			MultiLineCommentClose: "*/",
		}
		return Go

	default:
		return Language{}
	}
}
