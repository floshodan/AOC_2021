package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://"+os.Getenv("SERVER_IP")+"/input", nil)
	//req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/3/input", nil)
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
	// 00100
	// 11110
	// 10110
	// 10111
	// 10101
	// 01111
	// 00111
	// 11100
	// 10000
	// 11001
	// 00010
	// 01010

	// Step1 keep only with 1 bits in front (because 1bit has more 1ons then zeors) ->
	// 11110
	// 10110
	// 10111
	// 10101
	// 11100
	// 10000
	// 11001

	// Step2: consider second bit
	// more 0 bits (4) than 1 bits (3) so keep only zeros bits as second position
	// 10110
	// 10111
	// 10101
	// 10000

	// step3: third bit has more ones, keep them
	// 10110
	// 10111
	// 10101

	// step4: keep with one
	// 10110
	// 10111

	// if theres eqaul bits in the fith the oxygen is with a 1 so 10111 is oxygen

	gammabits := make([]int, len(bits[0]))

	//fmt.Println(bits[0][0])
	//fmt.Println(len(bits[0])) / 5
	//fmt.Println(bits)

	for j := 0; j < len(bits[0]); j++ {
		zero := 0
		one := 0
		gamma := 0

		// line by line first element
		for i := 0; i <= len(bits)-1; i++ {
			if bits[i][j] == 48 {
				zero += 1
			} else if bits[i][j] == 49 {
				one += 1
			}

			//fmt.Printf("i = %d, zero = %d, one = %d \n", i, zero, one)
		}

		//fmt.Println(zero, one, gamma)

		if zero > one {
			gamma = 48
		} else if one > zero {
			gamma = 49
		}
		//fmt.Println(zero, one, gamma)
		gammabits[j] = gamma
	}

	//fmt.Println(bits[11])
	//fmt.Println(gammabits)
	fmt.Println(gammabits)
	fmt.Printf("length gammabits %d \n", len(gammabits))
	fmt.Printf("length bits %d \n", len(bits))
	lenall := len(bits)
	removed := 0
	for i := 0; i < len(gammabits)-1; i++ {
		removed = 0
		for j, v := range bits {
			//fmt.Println(v[i])
			if lenall-removed <= 2 {
				fmt.Println("break")
				break
			}

			if int(v[i]) != gammabits[i] {
				//fmt.Printf("nope index =%d \n", j) // expect 0, 5, 6, 10, 11
				//bits[j] = "xxxxxxxxxxxx"
				bits[j] = "xxxxx"
				removed++
				continue
			}

			fmt.Printf("%s index = %d \n", bits[j], i)
			//fmt.Println(int(v[i]), g)
			//fmt.Println(left)
		}

		fmt.Printf("new line ---- i is %d removed = %d left = %d \n", i, removed, lenall-removed)
		//	fmt.Println()
		//  fmt.Println()
	}

	fmt.Println(gammabits)

	/*	for _, v := range bits {
			fmt.Println(v)
		}
	*/

	//fmt.Println(binary_to_decimal(gammabits) * binary_to_decimal(omegabits))

}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
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
