package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://"+os.Getenv("SERVER_IP")+"/input4", nil)
	//req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/4/input", nil)
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

	boards := make([][5][5]string, 0)

	bingonumbers := strings.Split(input[0], ",")
	line := strings.Fields(input[2])
	fmt.Printf("first element of first line = %s \n", line[0])
	fmt.Printf("length of line %d \n", len(line))
	fmt.Printf("bingo order %s \n", bingonumbers)
	fmt.Printf("---------------- \n")
	//fmt.Println(line)
	// start at one because 0 is bingo numbers
	var bingogrid [5][5]string
	counter := 0
	//line := strings.Fields(input[i]) // line[0] = 22; line[1] = 13; line[2] = 17; line[3] = 11; line[4] = 0
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			fmt.Println()
			continue
		}
		//line := strings.Fields(input[i]) // line[0] = 22; line[1] = 13; line[2] = 17; line[3] = 11; line[4] = 0
		// rows
		for row := 0; row < 5; row++ {

			//fill column
			for col := 0; col < 5; col++ {
				fmt.Printf("%s col=%d row=%d i = %d \n", line[col], col, row, i)
				bingogrid[row][col] = line[col]

				// bingogrid[0][0] = line[x] // 22 x = 0
				// bingogrid[0][1] = line[1] // 13 x = 1
				// bingogrid[0][4] = line[4] // 0  x = 4
			}
			if input[i+counter] == "" {
				fmt.Println("HI THERE!")
				counter = 0
			} else {
				fmt.Println("HERE!")
				line = strings.Fields(input[i+counter])
				counter += 1

				boards = append(boards, bingogrid)
				fmt.Println(bingogrid)
			}

		}
		// add bingogrid to boards[]

		/*		for y := 0; y < 5; y++ {
					for z := 0; z < 5; z++ {
						fmt.Printf("a[%d][%d] = %s\n", y, z, bingogrid[y][z])
					}
				}
		*/

		//fmt.Println(line)
	}

	//fmt.Println(boards[0])
	//fmt.Println(boards[0][1][0])

	//for _, v := range bingo {
	//fmt.Println(v)
	//}

}
