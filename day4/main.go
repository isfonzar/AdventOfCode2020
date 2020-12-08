package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	Passport     map[string]string
	PassportList []Passport
)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	entry := ""
	var ppList PassportList

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			pp := processPassport(entry)
			ppList = append(ppList, pp)
			entry = ""
		} else {
			if entry != "" {
				entry += " " + line
			} else {
				entry += line
			}
		}
	}

	fmt.Println(checkValidPassports(ppList))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processPassport(passport string) Passport {
	pp := Passport{}

	entries := strings.Split(passport, " ")
	for _, entry := range entries {
		attr := strings.Split(entry, ":")
		pp[attr[0]] = attr[1]
	}

	return pp
}

func checkValidPassports(list PassportList) int {
	valids := 0

	for _, p := range list {
		if hasRequiredFields(p) {
			valids++
		}
	}

	return valids
}

func hasRequiredFields(p Passport) bool {
	for _, field := range requiredFields {
		_, ok := p[field]
		if !ok {
			return false
		}
	}

	return true
}
