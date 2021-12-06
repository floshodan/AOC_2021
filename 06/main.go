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

	client := &http.Client{}

	//req, err := http.NewRequest("GET", "http://"+os.Getenv("SERVER_IP")+"/input4", nil)
	req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/6/input", nil)
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
	input := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	input_int := make([]int, 0)
	initial := strings.Split(input[0], ",")

	for _, v := range initial {
		no, _ := strconv.Atoi(v)
		input_int = append(input_int, no)
	}

	// 3,4,3,1,2
	// after one day 2 ... 0, reset to 6 + create new with 8
	// each day if fish becomes 0 it becomes a 6 + adds new on with 8
	// how many lanternwith after 80 day

	fmt.Printf("initial state = %s \n", initial)
	for i := 0; i < 80; i++ {

		for i, v := range input_int {
			if v != 0 {
				input_int[i] = v - 1
				continue
			}
			input_int[i] = 6
			input_int = append(input_int, 8)
		}

	}
	fmt.Println(len(input_int))
}
