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

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(regex)
		matches := r.FindStringSubmatch(line)

		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		isValid := process(min, max, matches[3], matches[4])
		if isValid {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Valid passwords: %d\n", count)
}

func process(min int, max int, char string, pw string) bool {
	c := strings.Count(pw, char)

	if c < min {
		return false
	}
	if c > max {
		return false
	}

	return true
}
