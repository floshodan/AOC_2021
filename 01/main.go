package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/1/input", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Cookie", "session="+os.Getenv("AOC_SESSION"))
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("error = %s \n", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	numbers := make([]int, 0)

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	p1 := 0
	p2 := 0
	for i, element := range numbers {
		if i >= 1 && element > numbers[i-1] {
			p1 += 1
		}
		if i >= 3 && element > numbers[i-3] {
			p2 += 1
		}
	}
	/*
		for i := 0; i <= len(numbers)-1; i++ {
			if i >= 1 && numbers[i] > numbers[i-1] {
				result += 1
			}
		}
	*/
	fmt.Printf("part1: %v\n", p1)
	fmt.Printf("part2: %v\n", p2)

}
