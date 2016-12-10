package parser

import (
	"io/ioutil"
	"log"
	"os"
	//"strings"
	"testing"
)

func TestScan(t *testing.T) {
	text := "// TODO TESTING"
	folder, err := os.Getwd()
	if err != nil {
		t.Error("cannot find current folder")
	}

	file, err := ioutil.TempFile(folder, "scan")
	if err != nil {
		t.Error("cannot create temporary file")
	}

	file.WriteString(text)
	issues := Scan()
	if len(issues) == 0 {
		t.Error("issues not found")
	}
	os.Remove(file.Name())
}

func TestParseLine(t *testing.T) {
	testLines := [...]string{
		"//TODO fix - bug - me",
		"// TODO fix - bug - me",
		"/*TODO fix - bug - me",
		"/* TODO fix - bug - me",
	}

	for _, line := range testLines {
		log.Printf("testing line %v\n", line)
		issue := ParseLine(line)
		log.Printf("parsed: %v", *issue.Title)
		/*
			if strings.TrimSpace(*issue.Title) != "fix" {
				t.Error("title assign is broken")
			}

			if strings.TrimSpace(issue.Label) != "bug" {
				t.Error("label assign is broken")
			}

			if issue.Assignee != "me" {
				t.Error("assignee is broken")
			}
		*/
	}
}
