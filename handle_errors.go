package main

import (
	"fmt"
	"log"
	"os"
)

func HandleProgramError(msg string, err error) {
	fmt.Printf(msg, err)
	log.Fatalln(err)
	os.Exit(1)
}