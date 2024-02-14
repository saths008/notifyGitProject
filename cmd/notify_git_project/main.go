package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
)

func main() {
	err := beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello, World!")
}
