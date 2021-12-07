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

	//req, err := http.NewRequest("GET", "http://"+os.Getenv("SERVER_IP")+"/input7", nil)
	req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/7/input", nil)
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

	/*for _, v := range input_int {
		fmt.Println(v)
	}
	*/

	sorted := sorted(input_int)
	//	median2 := sorted[len(sorted)/2]
	median, _ := median(sorted)
	median--

	fmt.Println(median)

	fuel := 0
	for _, v := range input_int {
		add := v - median
		if add < 0 {
			add = add * -1
		}
		fuel += add
	}
	fmt.Println(fuel)
}

// bubble sort i think
func sorted(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-2; j++ {
			if slice[j] > slice[j+1] {
				first := slice[j]
				second := slice[j+1]
				slice[j] = second
				slice[j+1] = first
			}
			continue
		}
	}
	return slice
}

// input slice must be sorted!
func median(input []int) (median int, err error) {

	l := len(input)
	if l == 0 {
		return
	} else if l%2 == 0 {
		median, _ = mean(input[l/2-1 : l/2+1])
	} else {
		median = input[l/2]
	}

	return median, nil
}

func mean(input []int) (int, error) {

	sum := 0
	for _, n := range input {
		sum += n
	}
	return sum / len(input), nil
}
