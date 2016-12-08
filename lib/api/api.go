package api

import (
	"io/ioutil"
	"log"
)

func SetHook(folder string) (bool, error) {
	file, err := ioutil.ReadFile("../hooks/pre-commit")
	if err != nil {
		log.Fatalln(err)
	}

	ioutil.WriteFile(folder+"/.git/pre-commit", file, 0777)
	return true, nil
}
