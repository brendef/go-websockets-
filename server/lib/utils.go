package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func WatchFile(filePath string, lastModifiedTime time.Time) string {
	for {
		// Get the file information
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return fmt.Sprintf("Error: %s", err)
		}

		// Get the modification time of the file
		modifiedTime := fileInfo.ModTime()

		// Check if the file has been modified
		if !modifiedTime.Equal(lastModifiedTime) {
			// File has been modified, read the content
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return fmt.Sprintf("Error reading file: %s", err)
			}

			// Update the last modified time
			lastModifiedTime = modifiedTime

			// Return the file content as a string
			return string(content)
		}

		// Sleep for a duration before checking again (e.g., every 1 second)
		time.Sleep(1 * time.Second)
	}
}

func TextToMap(content string) (map[string]int, error) {
	lines := strings.Split(string(content), "\n")
	resultMap := make(map[string]int)

	for _, line := range lines {
		// Split each line by ":"
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue // Skip lines without proper key-value format
		}

		// Trim spaces and convert value to an integer
		key := strings.TrimSpace(parts[0])
		value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("error converting value to int: %s", err)
		}

		// Add key-value pair to the map
		resultMap[key] = value
	}

	return resultMap, nil
}

func MonitorFileChanges(fileName string) <-chan string {
	changes := make(chan string)

	go func() {
		defer close(changes)

		// Read the initial content of the file
		initialContent, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		for {
			// Read the content of the file
			currentContent, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			// Compare current content with the initial content
			if string(currentContent) != string(initialContent) {
				// Content has changed, send the changes through the channel
				changes <- string(currentContent)
				initialContent = currentContent // Update initial content
			}

			// Sleep for a while before checking again (e.g., every 2 seconds)
			time.Sleep(2 * time.Second)
		}
	}()

	return changes
}
