package api

import (
	"io/ioutil"
	"log"
)

//SetHook reads the pre-commit and saves it in the .git/hooks/ folder
func SetHook(folder string) (bool, error) {
	var hook = []byte("#!/bin/bash\ngoTrack scan")
	_, err := ioutil.ReadDir(".git")
	if err != nil {
		log.Fatalln("You are not in the root of your repository, make sure to stay in the top folder where .git is")
	}
	log.Println("Installing git hook...")

	err = ioutil.WriteFile(folder+"/.git/hooks/pre-commit", hook, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Done!")
	log.Println("If you haven't done yet, go to https://github.com/settings/tokens and create one token")
	log.Println("then use it in with the env $GOTRACK")
	log.Println("zsh users: echo 'export GOTRACK=<token>' >> ~/.zshrc")
	log.Println("bash users: echo 'export GOTRACK=<token>' >> ~/.bashrc")
	return true, nil
}
