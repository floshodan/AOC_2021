package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main_() {

	// horizontal forward
	// depth up or down

	client := &http.Client{}

	//req, err := http.NewRequest("GET", "http://195.201.10.231/input", nil)
	req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/3/input", nil)
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
	bits := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		bits = append(bits, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// gamma rate first bit of each line most common is gamma
	// episoln rate is xor of gamma

	gammabits := make([]int, len(bits[0]))
	omegabits := make([]int, len(bits[0]))

	for j := 0; j < len(bits[0]); j++ {
		zero := 0
		one := 0
		gamma := 0

		for i := 0; i <= len(bits)-1; i++ {
			if bits[i][j] == 48 {
				zero += 1
			} else if bits[i][j] == 49 {
				one += 1
			}
		}
		if zero > one {
			gamma = 0
		} else if one > zero {
			gamma = 1
		}
		//fmt.Println(zero, one, gamma)
		gammabits[j] = gamma
		omegabits[j] = 1 - gamma
	}

	fmt.Println(gammabits, omegabits)
	// 3529 * 566
	fmt.Println(binary_to_decimal(gammabits) * binary_to_decimal(omegabits))

}

// IntPow calculates n to the mth power. Since the result is an int, it is assumed that m is a positive power
func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

// converts a int slice (must contain only 0, 1) to decimal
func binary_to_decimal(slice []int) int {
	sum := 0
	j := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		sum += slice[i] * IntPow(2, j)
		j--
	}
	return sum
}

func problem2() {

}
