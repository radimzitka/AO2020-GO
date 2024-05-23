package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Password struct {
	min      int
	max      int
	char     byte
	password string
}

func (p Password) isValidTaskOne() int {
	count := strings.Count(p.password, string(p.char))
	if count >= p.min && count <= p.max {
		return 1
	}
	return 0
}

func (p Password) isValidTaskTwo() int {
	// Task two
	if p.password[p.min-1] == p.char && p.password[p.max-1] != p.char {
		return 1
	}
	if p.password[p.max-1] == p.char && p.password[p.min-1] != p.char {
		return 1
	}
	return 0
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

	// row format is like '4-8 g: ggtxgtgbg'
	for fileScanner.Scan() {
		row := fileScanner.Text()

		// split by space
		by_space := strings.Split(row, " ")
		min, _ := strconv.Atoi(strings.Split(by_space[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(by_space[0], "-")[1])
		char := by_space[1][0]
		password := by_space[2]

		/* BAD SOLUTION  */
		/* match1 := regexp.MustCompile(`(.*?)\-`).FindStringSubmatch(row)
		min, _ := strconv.Atoi(match1[1])

		match2 := regexp.MustCompile(`\-(.*?)\ `).FindStringSubmatch(row)
		max, _ := strconv.Atoi(match2[1])

		char := regexp.MustCompile(`\ (.*?)\:`).FindString(row)[0]

		password := regexp.MustCompile(`\:\ (.*?)$`).FindStringSubmatch(row) */

		passwords = append(passwords, Password{min, max, byte(char), password})
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
		countTaskOne += p.isValidTaskOne()
		countTaskTwo += p.isValidTaskTwo()
	}

	fmt.Println("Task one:", countTaskOne)
	fmt.Println("Task two:", countTaskTwo)
}
