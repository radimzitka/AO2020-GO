package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BagContain struct {
	value int
	color string
}

func readFile(name string) ([]string, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Nelze načíst soubor: %v", err)
		return nil, err
	}

	rules := strings.Split(string(content), "\n")

	var full_strings []string
	for _, rule := range rules {
		if strings.TrimSpace(rule) != "" {
			full_strings = append(full_strings, rule)
		}
	}
	return full_strings, nil
}

func howManyBagsBagContain(bagColor string, m map[string][]BagContain) int {
	sum := 0
	for _, bag := range m[bagColor] {
		sum += bag.value * (1 + howManyBagsBagContain(bag.color, m))
	}
	return sum
}

func containsBag(bagColor string, color string, m map[string][]BagContain) bool {
	for _, bag := range m[bagColor] {
		if bag.color == color {
			return true
		}
		if containsBag(bag.color, color, m) == true {
			return true
		}
	}
	return false
}

func main() {
	rules, err := readFile("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	bags := map[string][]BagContain{}

	for _, rule := range rules {
		splitted := strings.Split(rule, " ")
		name := splitted[0] + splitted[1]
		if splitted[4] != "no" {
			// Bag contains 1 other bag at least
			by_commas := strings.Split(rule, ",")
			for _, bag := range by_commas {
				splitted = strings.Split(bag, " ")
				str_len := len(splitted)
				num, _ := strconv.Atoi(splitted[str_len-4])
				b := BagContain{num, splitted[str_len-3] + splitted[str_len-2]}
				bags[name] = append(bags[name], b)
			}
		}
	}

	containShinyGold := 0
	for color, _ := range bags {
		if containsBag(color, "shinygold", bags) {
			containShinyGold++
		}
	}
	fmt.Println("Task one:", containShinyGold)
	fmt.Println("Task two:", howManyBagsBagContain("shinygold", bags))
}
