package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("sudoku.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(file)
	gridStr, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	grid := make([][]int, 9)
	for i := 0; i < 9; i++ {
		grid[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			if len(gridStr[i][j]) > 0 {
				num, err := strconv.ParseInt(gridStr[i][j], 10, 64)
				if err != nil {
					log.Fatal("Error parsing csv file.", err)
				}
				grid[i][j] = int(num)
			} else {
				grid[i][j] = 0 // this isn't really necessary since the 0 value of an int is 0
			}
		}
	}
	PrintGrid(grid)
}

func PrintGrid(grid [][]int) {
	for i, hLine := range grid {
		if i%3 == 0 {
			fmt.Println("  -------------------------- ")
		}
		for j, num := range hLine {
			if j%3 == 0 {
				fmt.Print(" | ")
			}
			fmt.Printf(" %d", num)
			if j == 8 {
				fmt.Print(" |\n")
			}
		}
	}
	fmt.Println("  -------------------------- ")
}
