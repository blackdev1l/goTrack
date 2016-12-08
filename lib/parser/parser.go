package parser

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// TODO  TEST
func GetIssues(filename string) bool {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "TODO") {
			log.Println("found TODO")
			log.Println(scanner.Text())
			return true
		}
	}
	return false
}
