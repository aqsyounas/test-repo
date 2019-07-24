package main

import (
	"fmt"

	"github.com/aqsyounas/test-repo/database"
)

func main() {

	fmt.Println("This is going to be a test branch")

	database.Database()
	fmt.Println("Done")
	fmt.Println("added into fork")
	fmt.Println("this commit was added after the pull request was submitted")
}
