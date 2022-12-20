package main

import (
	"fmt"
	"os"
	"strings"
)

// Map ascii font
func mapFonts(font string) (map[int][]string, error) {
	readFile, err := os.ReadFile(font)
	if err != nil {
		return nil, fmt.Errorf("could not read the content in the file: %v", err)
	}
	slice := strings.Split(string(readFile), "\n")

	ascii := make(map[int][]string)
	i := 31

	for _, ch := range slice {
		if ch == "" {
			i++
		} else {
			ascii[i] = append(ascii[i], ch)
		}
	}

	return ascii, nil
}

// Handle user input and map it to ascii
func mapUserInput(input string, ascii map[int][]string) (string, error) {
	lines := strings.Split(input, "\r\n")
	output := ""

	for _, line := range lines {
		characters := []rune(line)

		if line != "" {
			for i := 0; i < 8; i++ {
				for _, ch := range characters {
					if ch >= 32 && ch <= 126 {
						output = output + ascii[int(ch)][i]
					} else {
						return "", fmt.Errorf("input includes non ascii character(s), please use ascii character(s)")
					}
				}
				output = output + "\n"
			}
		} else {
			output = output + "\n"
		}
	}
	return output, nil
}

// User input validation
func inputValidation(input string) bool {
	for _, ch := range input {
		if ch < 32 || ch > 126 {
			if ch == 10 || ch == 13 {
				continue
			}
			return false
		}
	}
	return true
}
