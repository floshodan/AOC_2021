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

	req, err := http.NewRequest("GET", "http://"+os.Getenv("SERVER_IP")+"/input6", nil)
	//req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/6/input", nil)
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

	fmt.Printf("initial state = %s \n", initial)
	m := make(map[int][]int, 0)

	for i := 0; i < 9; i++ {
		for j := 0; j < 19; j++ {
			m[i] = append(m[i], 0)
		}
	}
	fmt.Println(m)
	/*	for key, element := range m {
			fmt.Println("Key:", key, "=>", "Element:", element)
		}
	*/

	//fishes = append(fishes, len(initial))

	for i := 0; i < 19; i++ {
		for _, n := range input_int {
			m[n][i] += 1
		}
	}

	for key, element := range m {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

}

// returns count of fishes the next day
func nextDay(table []int) int {
	backup_table := make([]int, 9)
	backup_table[8] += table[0]
	backup_table[6] += table[0]
	for i := 1; i < 9; i++ {
		backup_table[i-1] += table[i]

	}
	copy(table, backup_table)

	//count fish
	sum := 0
	for _, v := range table {
		sum += v
	}

	return sum

}
