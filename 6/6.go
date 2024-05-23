package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(name string) ([]string, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Nelze načíst soubor: %v", err)
		return nil, err
	}

	groups := strings.Split(string(content), "\n\n")
	return groups, nil
}

func removeDuplicates(s string) string {
	seen := make(map[rune]bool)
	result := []rune{}

	for _, char := range s {
		if !seen[char] && char != '\n' {
			seen[char] = true
			result = append(result, char)
		}
	}
	return string(result)
}

func leaveOnlyDuplicates(s string) string {
	splitted := strings.Split(s, "\n")

	for i := 1; i < len(splitted); i++ {
		for _, char := range splitted[0] {
			if !strings.ContainsRune(splitted[i], char) {
				splitted[0] = strings.ReplaceAll(splitted[0], string(char), "")
			}
		}
	}
	return splitted[0]
}

func main() {
	groups, err := readFile("input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	sumOne := 0
	sumTwo := 0
	for _, group := range groups {
		sumOne += len(removeDuplicates(group))
		sumTwo += len(leaveOnlyDuplicates(group))
	}

	fmt.Println("Part one:", sumOne)
	fmt.Println("Part two:", sumTwo)
}
