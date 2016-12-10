package github

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
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

func CreateIssue(client *github.Client, issueReq *github.IssueRequest) *github.Issue {
	author, repo := getRepoInfo()
	log.Println(*issueReq.Title)
	issue, _, err := client.Issues.Create(author, repo, issueReq)
	if err != nil {
		println("Couldn't create issue.")
		log.Fatalln(err)
		return nil
	}
	return issue
}

func GetClient() *github.Client {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GOTRACK")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	return client
}

/* Example usage
func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GOTRACK")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	CreateIssue(client, repoAuthor, repoName, "Fix this issue too", "This issue must also be fixed, but the interesting thing is that it was generated automatically.")
	println("Printing issues...")
	PrintIssues(client, repoAuthor, repoName)
}
*/

func getRepoInfo() (author, repo string) {
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
	author = repoParts[1]
	repository := repoParts[2]
	return author, repository
}
