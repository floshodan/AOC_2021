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
	req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/4/input", nil)
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
	var bingoboard [5][5]string
	//line := strings.Fields(input[i]) // line[0] = 22; line[1] = 13; line[2] = 17; line[3] = 11; line[4] = 0
	row := 0
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			row = 0
			boards = append(boards, bingoboard)
			continue
		}
		line = strings.Fields(input[i])
		for col := 0; col < 5; col++ {
			bingoboard[row][col] = line[col]
			// bingogrid[0][0] = 22
			// bingogrid[0][1] =13
		}
		row += 1
		if i == len(input)-1 {
			boards = append(boards, bingoboard)
		}
	}
	fmt.Printf("total boards = %d \n", len(boards))
	fmt.Printf("boards array %s \n", boards)
	fmt.Printf("---------------- \n")

	winning_boards := make([]int, len(boards))

	for _, v := range bingonumbers {
		//fmt.Printf("number is %s \n", v)
		//loop through boards
		for i, _ := range boards {
			if intInSlice(i, winning_boards) {
				continue
			}
			for j := 0; j < 5; j++ {
				for y := 0; y < 5; y++ {
					//fmt.Println(boards[i][j][y])
					if v == boards[i][j][y] {
						//fmt.Printf("FOUND in Board %d at row=%d col=%d  \n", i+1, j+1, y+1)
						boards[i][j][y] = "-1"
						//check if won
						won := check(boards[i], j, y)
						if won {
							//fmt.Printf("board %d with endnumber %s has won! \n", i+1, v)
							winning_boards = append(winning_boards, i)
							sum := 0
							for row := 0; row < 5; row++ {
								for col := 0; col < 5; col++ {
									if boards[i][row][col] == "-1" {
										continue
									}
									no, _ := strconv.Atoi(boards[i][row][col])
									sum += no
								}
							}
							endnumber, _ := strconv.Atoi(v)
							fmt.Printf("board %d with endnumber %s has won! Boardsum=%d \n", i+1, v, sum*endnumber)
						}
					}
				}
			}
		}
	}
}

func check(board [5][5]string, row int, col int) bool {
	won := true
	for i := 0; i < 5; i++ {
		if board[row][i] != "-1" {

			for j := 0; j < 5; j++ {
				if board[j][col] != "-1" {
					return false
				} else {
					continue
				}
			}

		} else {
			continue
		}
	}

	return won
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
