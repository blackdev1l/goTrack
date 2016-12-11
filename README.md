# GoTrack

## Installation
`go get -u github.com/blackdev1l/goTrack/cmd/goTrack`

## Usage
`goTrack i[nit]` inside the top level of a git repo with the remote hosted on github
this will install a pre-commit hook which will scan the diff before a commit and create an issue when it will find a "TODO" in a comment.

**note:** you will need a github authorization token, you can create one [here](https://github.com/settings/tokens)

## Token
Authorization token will be read by the environment variable `GOTRACK`
* for bash users: `echo 'export GOTRACK=<token> >> ~/.bashrc'
* for zsh users: `echo 'export GOTRACK=<token> >> ~/.zshrc'

### TODO TAG 
You can use any kind of comment tag, the only thing it matters is that you will use the word "TODO" in caps lock, otherwise it will be ignored.
Todo tags will be parsed this way:
`TODO <title> - <label> - <assignee>` other stuff after that will be ignored.

## Feature
- [x] Api for github 
- [ ] Api for Google calendar
- [ ] Time Tracker
- [x] Git hooks
- [ ] settings

## Process

1. `goTrack init` <-- this initialize the repo installing git hooks
	1.  git hook -> pre-commit: Parses all the project and create issues with all specific tags and such
2. `goTrack list` <-- show all the issues in this repo with their ID



## TAG Structure

`TODO <title> <labels> <assignee>`

Example: `// TODO finish this readme - enhancement - me`
this will create an issue with title "finish this readme" , labeled "enhancement" and assigned to my user
