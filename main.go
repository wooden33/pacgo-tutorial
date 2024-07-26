package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var maze []string

func loadMase(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

func printScreen() {
	for _, line := range maze {
		fmt.Println(line)
	}
}

func main() {
	// initialize game

	// load resources
	err := loadMase("maze01.txt")
	if err != nil {
		log.Println("failed to load maze", err)
		return
	}

	// game loop
	for {
		// update screen
		printScreen()
		// process input

		// process movement

		// process collisions

		// Temp: break infinite loop
		break

		// repeat
	}
}
