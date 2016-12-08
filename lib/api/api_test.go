package api

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestSetHooks(t *testing.T) {
	os.MkdirAll("tmp/.git", 0777)
	folder, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	SetHook(folder + "/tmp")
	_, err = ioutil.ReadFile("tmp/.git/pre-commit")
	if err != nil {
		t.Error(err)
	}

	os.RemoveAll("tmp/")

}
