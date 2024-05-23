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
	var seats []string

	for fileScanner.Scan() {
		seats = append(seats, fileScanner.Text())
	}

	return seats, nil
}

func findMissingSeat(s []string) int {
	seatsOccupied := map[int]bool{}
	for _, seat := range s {
		seatsOccupied[countSeatID(seat)] = true
	}

	for i := 15; i < len(seatsOccupied); i++ {
		if !seatsOccupied[i] {
			return i
		}
	}

	return 0
}

func countSeatID(s string) int {
	seatID := 0
	values := [7]int{64, 32, 16, 8, 4, 2, 1}

	for i := 0; i < 8; i++ {
		if s[i] == 'B' {
			seatID += values[i]
		}
	}

	seatID *= 8

	for i := 4; i < 7; i++ {
		if s[i+3] == 'R' {
			seatID += values[i]
		}
	}
	return seatID
}

func main() {
	seats, err := readFile("input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	maxSeatID := 0
	for _, seat := range seats {
		if countSeatID(seat) > maxSeatID {
			maxSeatID = countSeatID(seat)
		}
	}

	fmt.Println("Part one:", maxSeatID)
	fmt.Println("Part one:", findMissingSeat(seats))
}
