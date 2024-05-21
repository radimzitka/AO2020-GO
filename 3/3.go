package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(name string) ([]string, error) {
	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var m []string

	for fileScanner.Scan() {
		m = append(m, fileScanner.Text())
	}

	file.Close()

	return m, nil
}

func countTrees(m []string, right int, down int) int {
	pos := right
	trees := 0
	for i := down; i < len(m); i += down {
		if m[i][pos%len(m[i])] == '#' {
			trees++
		}
		pos += right
	}

	return trees
}

func main() {
	m, err := readFile("input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Task one: ", countTrees(m, 3, 1))
	fmt.Println("Task two: ", countTrees(m, 1, 1)*countTrees(m, 3, 1)*countTrees(m, 5, 1)*countTrees(m, 7, 1)*countTrees(m, 1, 2))
}
