package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Could not convert to integer val: %s", line)
		}

		input = append(input, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	num1, num2 := process(input, 2020)

	fmt.Printf("num1 %d\n", num1)
	fmt.Printf("num2 %d\n", num2)
	fmt.Printf("num1 + num2 = %d\n", num1+num2)
	fmt.Printf("num1 * num2 = %d\n", num1*num2)

}

func process(in []int, target int) (int, int) {
	m := make(map[int]int)

	for key, val := range in {
		lt := target - val

		_, ok := m[lt]
		if !ok {
			m[val] = key
		} else {
			return val, lt
		}
	}

	return -1, -1
}
