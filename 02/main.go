package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	// horizontal forward
	// depth up or down

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/2/input", nil)
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
	instructions := make([]string, 0)

	for scanner.Scan() {
		instruction := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, instruction)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	hor := 0
	ver := 0
	aim := 0

	for i := 0; i <= len(instructions)-1; i++ {
		//instruction, mov := instructions[i].split()
		s := strings.Split(instructions[i], " ")
		instruction := s[0]
		mov, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		if instruction == "forward" {
			hor += mov
			ver += aim * mov
		} else if instruction == "up" {
			//	ver -= mov
			aim -= mov
		} else if instruction == "down" {
			//	ver += mov
			aim += mov
		}
	}

	fmt.Println(ver * hor)
}
