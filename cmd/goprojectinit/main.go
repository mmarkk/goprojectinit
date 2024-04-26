package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Create a reader to read input from standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of your application: ")
	appName, _ := reader.ReadString('\n')
	appName = appName[:len(appName)-1]  // Remove the newline character

	// Define the folder structure
	directories := []string{
		"cmd",
		"cmd/" + appName,
		"configs",
		"internal",
		"pkg",
		"scripts",
		"vendor",
		"api",
		"ui",
		"deployments",
		"build",
		"tests",
		"docs",
	}

	// Create base directory for the application
	basePath := filepath.Join(".", appName)
	if err := os.Mkdir(basePath, 0755); err != nil {
		fmt.Printf("Failed to create base directory: %s\n", err)
		return
	}

	// Create subdirectories
	for _, dir := range directories {
		dirPath := filepath.Join(basePath, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			fmt.Printf("Failed to create %s directory: %s\n", dir, err)
			return
		}
	}

	fmt.Println("Project folders created successfully.")
}

