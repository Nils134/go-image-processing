package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting program")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter path for input images: ")
	path, err := reader.ReadString('\n')

	if err != nil {
		panic("Parsing input failed")
	}

	path = strings.TrimSpace(path)

	fmt.Println("Path used is: " + path)
	entries, err := os.ReadDir(path)

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot find directory")
	}

	_, err = fmt.Printf("Found %d entries \n", len(entries))

	if err != nil {
		fmt.Println("Cannot parse number of entries")
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			// fmt.Println(entry.Name())
		}
	}

	fmt.Println("Program finished")
}
