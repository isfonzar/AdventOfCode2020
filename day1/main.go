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

	fmt.Println("Finding two numbers in input that add up to 2020")
	num1, num2 := findTwoSum(input, 2020)

	fmt.Printf("num1 %d\n", num1)
	fmt.Printf("num2 %d\n", num2)
	fmt.Printf("num1 + num2 = %d\n", num1+num2)
	fmt.Printf("num1 * num2 = %d\n", num1*num2)
	fmt.Println("---")
	fmt.Println("Finding three numbers in input that add up to 2020")
	num1, num2, num3 := findThreeSum(input, 2020)

	fmt.Printf("num1 %d\n", num1)
	fmt.Printf("num2 %d\n", num2)
	fmt.Printf("num3 %d\n", num3)
	fmt.Printf("num1 * num2 * num3 = %d\n", num1*num2*num3)
}

func findTwoSum(in []int, sum int) (int, int) {
	m := make(map[int]int)

	for key, val := range in {
		lt := sum - val

		_, ok := m[lt]
		if !ok {
			m[val] = key
		} else {
			return val, lt
		}
	}

	return -1, -1
}

func findThreeSum(in []int, sum int) (int, int, int) {
	for k, v := range in {
		ls := sum - v

		num1, num2 := findTwoSum(in[k:], ls)
		if num1 != -1 && num2 != -1 {
			return num1, num2, v
		}
	}

	return -1, -1, -1
}
