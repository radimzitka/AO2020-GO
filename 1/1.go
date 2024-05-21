package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func readFile(name string) ([]int, error) {
	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []int

	for fileScanner.Scan() {
		num, err := strconv.Atoi(fileScanner.Text())

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		fileLines = append(fileLines, num)
	}

	file.Close()

	return fileLines, nil
}

func main() {
	input, err := readFile("input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, numA := range input {
		if slices.Contains(input, 2020-numA) {
			fmt.Println("Task one: ", numA*(2020-numA))
			break
		}
	}

	for _, numA := range input {
		for _, numB := range input {
			if slices.Contains(input, 2020-numA-numB) {
				fmt.Println("Task two: ", numA*numB*(2020-numA-numB))
				return
			}
		}
	}
}
