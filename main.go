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
	var grid [9][9]int
	for i := 0; i < 9; i++ {
		grid[i] = [9]int{}
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
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			posibleSolutions := FindPosibleSolutions(x, y, grid)
			if len(posibleSolutions) == 1 {
				fmt.Printf("Definitive solution for (%d, %d): %v\n", x, y, posibleSolutions)
			}
		}
	}
}

// PrintGrid prints the complete Sudoku grid.
func PrintGrid(grid [9][9]int) {
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

// GetBox takes a position and a grid and returns the box that position belongs to.
// Each position having their respective value.
func GetBox(x, y int, grid [9][9]int) [3][3]int {
	var box [3][3]int
	x -= x % 3
	y -= y % 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			box[i][j] = grid[x+i][y+j]
		}
	}
	return box
}

// FindPosibleSolutions calculates the posible solutions for a single position in the grid.
//
// If the position already has a value, then this function returns an empty slice.
func FindPosibleSolutions(x, y int, grid [9][9]int) []int {
	// if the position is already filled, there's no need to find any other solution
	if grid[x][y] != 0 {
		return []int{}
	}
	// start with all posible solutions
	posibleSolutions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// find numbers in the same line (either vertical or horizontal)
	for i := 0; i < 9; i++ {

		// horizontal line check
		if grid[x][i] != 0 {
			// remove from posibleSolutions slice
			for index, element := range posibleSolutions {
				if element == grid[x][i] {
					posibleSolutions = append(posibleSolutions[:index], posibleSolutions[index+1:]...)
					break
				}
			}
		}

		// vertical line check
		if grid[i][y] != 0 {
			// remove from posibleSolutions slice
			for index, element := range posibleSolutions {
				if element == grid[i][y] {
					posibleSolutions = append(posibleSolutions[:index], posibleSolutions[index+1:]...)
					break
				}
			}
		}
	}
	// find numbers in the same box
	boxX := x - (x % 3)
	boxY := y - (y % 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			elem := grid[boxX+i][boxY+j]
			for i := range posibleSolutions {
				if posibleSolutions[i] == elem {
					posibleSolutions = append(posibleSolutions[:i], posibleSolutions[i+1:]...)
					break
				}
			}
		}
	}
	return posibleSolutions
}
