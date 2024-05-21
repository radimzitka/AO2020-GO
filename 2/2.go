package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	min      int
	max      int
	char     byte
	password string
}

func (p Password) isValidTaskOne() bool {
	count := strings.Count(p.password, string(p.char))
	if count >= p.min && count <= p.max {
		return true
	}

	return false
}

func (p Password) isValidTaskTwo() bool {
	if p.password[p.min-1] == p.char && p.password[p.max-1] != p.char {
		return true
	}

	if p.password[p.max-1] == p.char && p.password[p.min-1] != p.char {
		return true
	}

	return false
}

func readFile(name string) ([]Password, error) {
	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var passwords []Password

	// split je lepsi!
	for fileScanner.Scan() {
		row := fileScanner.Text()

		match1 := regexp.MustCompile(`(.*?)\-`).FindStringSubmatch(row)
		min, _ := strconv.Atoi(match1[1])

		match2 := regexp.MustCompile(`\-(.*?)\ `).FindStringSubmatch(row)
		max, _ := strconv.Atoi(match2[1])

		char := regexp.MustCompile(`\ (.*?)\:`).FindString(row)[0]

		password := regexp.MustCompile(`\:\ (.*?)$`).FindStringSubmatch(row)

		p := Password{min, max, byte(char), password[1]}

		passwords = append(passwords, p)
	}

	file.Close()

	return passwords, nil
}

func main() {
	passwords, err := readFile("input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	countTaskOne, countTaskTwo := 0, 0

	for _, p := range passwords {
		if p.isValidTaskOne() {
			countTaskOne++
		}

		if p.isValidTaskTwo() {
			countTaskTwo++
		}
	}

	fmt.Println("Task one:", countTaskOne)
	fmt.Println("Task two:", countTaskTwo)
}
