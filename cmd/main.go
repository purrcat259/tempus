package main

import (
	"log"
	"os"
	"tempus/models"
	"time"
	// "tempus/models"
)

// go run cmd/main.go -- start

func main() {
	log.Println("Starting Tempus")
	args := os.Args[1:]
	if args[0] == "--" {
		args = os.Args[2:]
	}
	log.Println(args)
	if len(args) != 1 {
		panic("Incorrect number of args")
	}
	command := args[0]
	log.Println("Running: " + command + "!")
	if command == "start" {
		message := "Message should be optional arg which goes here"
		entry := models.Entry{ID: "abc123", Start: time.Now(), Message: &message}
		log.Println(entry)
	} else if command == "stop" {
		// TODO:
		// 1. Find existing entry
		// 1a. If does not exist, panic
		// 2. Add end time
		// 3. OPTIONAL? Add "end" message?
		// 4. Save
	} else {
		panic("Unknown command")
	}
}
