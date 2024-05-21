package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Passport map[string]string

func fillPassport(s string) Passport {
	codes := strings.Split(s, " ")
	p := Passport{}

	for _, code := range codes {
		if len(strings.TrimSpace(code)) == 0 {
			continue
		}
		split := strings.Split(code, ":")
		p[split[0]] = split[1]
	}
	return p
}

func readFile(name string) ([]Passport, error) {
	passports := []Passport{}
	content, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pData := strings.Split(string(content), "\n\n")
	for _, onePass := range pData {
		p_str := strings.Replace(onePass, "\n", " ", -1)
		passports = append(passports, fillPassport(p_str))
	}
	return passports, nil
}

func (p Passport) isValid() bool {
	if len(p) == 8 {
		return true
	}

	_, cid_exists := p["cid"]
	if len(p) == 7 && !cid_exists {
		return true
	}

	return false
}

func main() {
	passports, err := readFile("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	valid := 0
	for _, p := range passports {
		if p.isValid() {
			valid++
		}
	}

	fmt.Println("Task one: ", valid)
	fmt.Println("Task two: ")
}
