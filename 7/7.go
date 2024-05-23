package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BagInBag struct {
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

func howManyBagsBagInBag(bagColor string, m map[string][]BagInBag) int {
	sum := 0
	for _, bag := range m[bagColor] {
		sum += bag.value * (1 + howManyBagsBagInBag(bag.color, m))
	}
	return sum
}

func containsBag(bagColor string, color string, m map[string][]BagInBag) bool {
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

	bags := map[string][]BagInBag{}

	for _, rule := range rules {
		splitted := strings.Split(rule, " ")
		name := splitted[0] + splitted[1]

		if splitted[4] != "no" {
			// Bag contains 1 other bag at least
			for _, bag := range strings.Split(rule, ",") {
				splitted = strings.Split(bag, " ")
				num, _ := strconv.Atoi(splitted[len(splitted)-4])
				bagName := splitted[len(splitted)-3] + splitted[len(splitted)-2]
				bags[name] = append(bags[name], BagInBag{num, bagName})
			}
		}
	}

	containShinyGold := 0
	for color := range bags {
		if containsBag(color, "shinygold", bags) {
			containShinyGold++
		}
	}
	fmt.Println("Task one:", containShinyGold)
	fmt.Println("Task two:", howManyBagsBagInBag("shinygold", bags))
}
