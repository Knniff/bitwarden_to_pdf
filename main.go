package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("No File given, exiting!")
	}
	arg := os.Args[1]

	makePDF(arg)
}
