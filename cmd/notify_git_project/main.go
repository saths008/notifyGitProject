package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatal("USAGE: notify_git_project PATH")
	}
	path := args[1]
	fmt.Println("path", path)
}
