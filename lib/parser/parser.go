package parser

import (
	"github.com/google/go-github/github"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

/*
 Scans diff finding for lines with "TODO" in them,
 WHen find one, it parses the info and initialize an Issue struct,
 In the end, it returns an array of Issue
*/
func Scan() []github.IssueRequest {
	var issues []github.IssueRequest
	diff, err := exec.Command("git", "diff").Output()
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(diff), "\n")
	for _, line := range lines {
		if strings.Contains(line, "TODO") {
			issue := ParseLine(line)
			issues = append(issues, issue)
		}
	}
	return issues
}

func ParseLine(line string) github.IssueRequest {
	var issue github.IssueRequest
	expression := "[\\s\\S]+/(/|\\*)(|\\s)(todo|TODO)"
	regex, err := regexp.Compile(expression)
	if err != nil {
		log.Fatalln(err)
	}
	parsedLine := string(regex.ReplaceAll([]byte(line), []byte("")))
	splitted := strings.Split(parsedLine, "-")
	if len(splitted) == 1 {
		issue.Title = &parsedLine
		return issue
	}

	for j, v := range splitted {
		switch j {
		case 0:
			issue.Title = &v
		case 1:
		//	issue.Labels = v
		case 2:
			//	i.Assignee = strings.TrimSpace(v)
		}
	}

	return issue
}
