# GoTrack

## project structure 

```
goTrack
├── cmd <-- here goes all the CLI packages
├── lib <-- modules 
│   ├── api
│   ├── github
│   ├── issue
│   ├── parser
│   │   ├── parser.go
│   │   └── parser_test.go
│   └── tracker
└── README.md
```
## What we need

- [ ] Api for github 
- [ ] Api for Google calendar
- [ ] Time Tracker
- [ ] Git hooks
- [ ] settings

## Process

1. `goTrack init` <-- this initialize the repo installing git hooks
	1.  git hook -> on commit: Parses all the project and create issues with all specific tags and such
2. `goTrack list` <-- show all the issues in this repo



## TAG Structure

`TODO <title> <labels> <assignee>`

Example: `// TODO finish this readme - enhancement - me`
this will create an issue with title "finish this readme" , labeled "enhancement" and assigned to my user



