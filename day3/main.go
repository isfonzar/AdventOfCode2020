package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type (
	Position struct {
		x int
		y int
	}
)

func NewPosition(x, y int) *Position {
	return &Position{
		x: x,
		y: y,
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	player := NewPosition(0, 0)
	yCords := getYCoords(player.y, 3, 323)
	i := 0
	count := 0

	fmt.Println(yCords)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		pos := yCords[i]
		pos = pos % len(line)
		if string(line[pos]) == `#` {
			count++
		}

		i++
	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getYCoords(startPos, stepY, steps int) []int {
	var cords []int

	cords = append(cords, startPos)
	for i := 0; i < steps; i++ {
		startPos += stepY
		cords = append(cords, startPos)
	}

	return cords
}
