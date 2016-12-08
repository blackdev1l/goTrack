package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func printIssues(client *github.Client, author string, name string) {
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

	println("Printing issues...")
	printIssues(client, repoAuthor, repoName)
}