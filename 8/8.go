package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name     string
	value    int
	executed bool
}

func readFile(name string) ([]Instruction, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Nelze načíst soubor: %v", err)
		return nil, err
	}

	codes := strings.Split(string(content), "\n")

	instructions := []Instruction{}
	for _, code := range codes {
		value, _ := strconv.Atoi(code[4:])
		instructions = append(instructions, Instruction{code[:3], value, false})
	}

	return instructions, nil
}

func execute(instructions []Instruction) (bool, int) {
	acc := 0
	for i := 0; i < len(instructions); i++ {
		if instructions[i].executed {
			return false, acc
		}
		instructions[i].executed = true
		if instructions[i].name == "acc" {
			acc += instructions[i].value
		}
		if instructions[i].name == "jmp" {
			i += instructions[i].value - 1
		}
	}
	return true, acc
}

func main() {
	instructions, err := readFile("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	executed, acc := execute(instructions)
	fmt.Println("Task one:", acc)

	executed, acc = false, 0
	for i := 0; i < len(instructions); i++ {
		if instructions[i].name == "jmp" {
			instructions[i].name = "nop"
			executed, acc = execute(instructions)
			instructions[i].name = "jmp"
		}
		if instructions[i].name == "nop" {
			instructions[i].name = "jmp"
			executed, acc = execute(instructions)
			instructions[i].name = "nop"
		}
		if executed {
			fmt.Println("Task two", acc)
			break
		}
		// All instructions are not executed
		for i := 0; i < len(instructions); i++ {
			instructions[i].executed = false
		}
	}
}
