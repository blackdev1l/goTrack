package parser

import (
	"testing"
)

func TestGetIssues(t *testing.T) {
	res := GetIssues("parser.go")
	if !res {
		t.Error("Expected true in return from GetIssues")
	}
}
