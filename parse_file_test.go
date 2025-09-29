package main

import (
	"bytes"
	"testing"

	"github.com/jonathon-chew/Nomos/rules"
)

// --- Mock setup ---

var captured []string

func resetCapture() { captured = []string{} }

func containsMessage(substr string) bool {
	for _, msg := range captured {
		if bytes.Contains([]byte(msg), []byte(substr)) {
			return true
		}
	}
	return false
}

var calledDocStrings []string

func resetDocStrings() { calledDocStrings = []string{} }

// --- Tests ---

func TestConstNotInCaps(t *testing.T) {
	rules := rules.Rules{ConstInCaps: true}
	src := []byte(`const Pi = 3.14`)

	resetCapture()
	err := process_file(src, rules)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !containsMessage("Const variable Pi isn't in caps lock") {
		t.Errorf("expected warning, got %v", captured)
	}
}

func TestInternalFunctionListed(t *testing.T) {
	rules := rules.Rules{ListInternalFunctions: true}
	src := []byte(`func internalFunc() {}`)

	resetCapture()
	_ = process_file(src, rules)

	if !containsMessage("internal function only") {
		t.Errorf("expected internal function info, got %v", captured)
	}
}

func TestNakedReturn(t *testing.T) {
	rules := rules.Rules{NoNakedReturns: true}
	src := []byte(`func f() {
		return
	}`)

	resetCapture()
	_ = process_file(src, rules)

	if !containsMessage("naked return") {
		t.Errorf("expected naked return warning, got %v", captured)
	}
}

func TestPrintfWithoutNewline(t *testing.T) {
	rules := rules.Rules{PrintFNewLine: true}
	src := []byte(`package main
import "fmt"
func main() {
	fmt.Printf("hello")
}`)

	resetCapture()
	_ = process_file(src, rules)

	if !containsMessage("does not end with a new line character") {
		t.Errorf("expected warning, got %v", captured)
	}
}

func TestPrintfWithNewline(t *testing.T) {
	rules := rules.Rules{PrintFNewLine: true}
	src := []byte(`package main
import "fmt"
func main() {
	fmt.Printf("hello\n")
}`)

	resetCapture()
	_ = process_file(src, rules)

	if !containsMessage("ends with a new line character") {
		t.Errorf("expected info, got %v", captured)
	}
}

func TestExportedFunctionRequiresDocstring(t *testing.T) {
	rules := rules.Rules{ExportedIdentifiersHaveComments: true, FunctionDocStrings: true}
	src := []byte(`package main

func PublicFunc() {}`)

	resetDocStrings()
	_ = process_file(src, rules)

	if len(calledDocStrings) == 0 {
		t.Fatalf("expected checkForDocStrings call, got none")
	}
	if calledDocStrings[0] != "Function:PublicFunc@3" {
		t.Errorf("unexpected docstring check: %v", calledDocStrings)
	}
}

func TestExportedVariableRequiresDocstring(t *testing.T) {
	rules := rules.Rules{ExportedIdentifiersHaveComments: true}
	src := []byte(`package main

var PublicVar = 10`)

	resetDocStrings()
	_ = process_file(src, rules)

	if len(calledDocStrings) == 0 {
		t.Fatalf("expected checkForDocStrings call for variable, got none")
	}
	if calledDocStrings[0] != "Variable:PublicVar@3" {
		t.Errorf("unexpected docstring check: %v", calledDocStrings)
	}
}

func ExportedTyepRequiresDocstring(t *testing.T) {
	rules := rules.Rules{ExportedIdentifiersHaveComments: true}
	src := []byte(`package main

type PublicType struct {}`)

	resetDocStrings()
	_ = process_file(src, rules)

	if len(calledDocStrings) == 0 {
		t.Fatalf("expected checkForDocStrings call for type, got none")
	}
	if calledDocStrings[0] != "Type:PublicType@3" {
		t.Errorf("unexpected docstring check: %v", calledDocStrings)
	}
}

func TestIgnoreIfInComments(t *testing.T) {
	rules := rules.Rules{IgnoreIfInComments: true, ConstInCaps: true}
	src := []byte(`
/* 
const Pi = 3.14
*/

`)

	resetCapture()
	_ = process_file(src, rules)

	if containsMessage("Const variable Pi isn't in caps lock") {
		t.Errorf("should have ignored because inside comment, got %v", captured)
	}
}
