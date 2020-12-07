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

const (
	regex = `^(\d+)\-(\d+) ([a-zA-Z]): ([a-zA-Z]+)$`
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	countValidOccurrences := 0
	countValidPositions := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(regex)
		matches := r.FindStringSubmatch(line)

		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		if isValidByOccurrences(min, max, matches[3], matches[4]) {
			countValidOccurrences++
		}
		if isValidByPosition(min, max, matches[3], matches[4]) {
			countValidPositions++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Valid passwords by occurrences: %d\n", countValidOccurrences)
	fmt.Printf("Valid passwords by position: %d\n", countValidPositions)
}

func isValidByOccurrences(min int, max int, char string, pw string) bool {
	c := strings.Count(pw, char)

	if c < min {
		return false
	}
	if c > max {
		return false
	}

	return true
}

func isValidByPosition(min int, max int, char string, pw string) bool {
	leftPos := min - 1
	rightPos := max - 1

	if string(pw[leftPos]) == char && string(pw[rightPos]) == char {
		return false
	}
	if string(pw[leftPos]) == char || string(pw[rightPos]) == char {
		return true
	}

	return false
}
