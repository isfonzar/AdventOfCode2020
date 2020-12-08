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
	yCords3 := getYCoords(player.y, 3, 323)
	yCords1 := getYCoords(player.y, 1, 323)
	yCords5 := getYCoords(player.y, 5, 323)
	yCords7 := getYCoords(player.y, 7, 323)
	i := 0
	i2 := 0
	count3 := 0
	count1 := 0
	count5 := 0
	count7 := 0
	count2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		pos3 := yCords3[i]
		pos1 := yCords1[i]
		pos5 := yCords5[i]
		pos7 := yCords7[i]
		pos2 := yCords1[i2]
		pos3 = pos3 % len(line)
		pos1 = pos1 % len(line)
		pos5 = pos5 % len(line)
		pos7 = pos7 % len(line)
		pos2 = pos2 % len(line)
		if string(line[pos3]) == `#` {
			count3++
		}
		if string(line[pos1]) == `#` {
			count1++
		}
		if string(line[pos5]) == `#` {
			count5++
		}
		if string(line[pos7]) == `#` {
			count7++
		}
		if i%2 == 1 && string(line[pos2]) == `#` {
			count2++
			i2++
		}

		i++
	}

	fmt.Printf("Right 1, down 1: %d\n", count1)
	fmt.Printf("Right 3, down 1: %d\n", count3)
	fmt.Printf("Right 5, down 1: %d\n", count5)
	fmt.Printf("Right 7, down 1: %d\n", count7)
	fmt.Printf("Right 1, down 2: %d\n", count2)
	fmt.Println(count3 * count1 * count5 * count7 * count2)

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
