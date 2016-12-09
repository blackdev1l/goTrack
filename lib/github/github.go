package github

import (
	"fmt"
	"os"
	"github.com/google/go-github/github"
/*  Imports used in examples
	"golang.org/x/oauth2"
	"os/exec"
	"regexp"
	"strings"
*/
)

func PrintIssues(client *github.Client, author string, name string) {
	issues, _, err := client.Issues.ListByRepo(author, name, nil)

	if err != nil {
		print("An error occurred.")
		os.Exit(1)
	}
	print(len(issues))
	println(" issue(s) open.")

	for i := 0; i < len(issues); i++ {
		issue := issues[i]
		fmt.Printf("#%d - %s\n", *issue.Number, *issue.Title)
	}
}

func CreateIssue(client *github.Client, author string, name string, title string, body string) *github.Issue {
	issueRequest := github.IssueRequest{Title: &title, Body: &body}
	issue, _, err := client.Issues.Create(author, name, &issueRequest)
	if err != nil {
		println("Couldn't create issue.")
		return nil
	}
	return issue
}

/* Example usage
func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GOTRACK")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	gitArgs := []string{"remote", "show", "origin"}
	gitOutput, err := exec.Command("git", gitArgs...).Output()
	if err != nil {
		print("Couldn't fetch repository URL.")
		os.Exit(1)
	}
	repoRegex := regexp.MustCompile("github\\.com/[A-Za-z0-9\\-]+/[A-Za-z0-9\\-]+")
	repoUrl := repoRegex.Find(gitOutput)
	repoParts := strings.Split(string(repoUrl), "/")
	if len(repoParts) != 3 {
		print("Couldn't parse repository URL.")
		os.Exit(1)
	}
	repoAuthor := repoParts[1]
	repoName := repoParts[2]

	CreateIssue(client, repoAuthor, repoName, "Fix this issue too", "This issue must also be fixed, but the interesting thing is that it was generated automatically.")
	println("Printing issues...")
	PrintIssues(client, repoAuthor, repoName)
}
*/
