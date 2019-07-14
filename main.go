package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type sudokuGrid [][]int

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
	var grid sudokuGrid = make([][]int, 9)
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

	PrintGrid(&grid)
	solveSudoku(&grid)
	PrintGrid(&grid)
}

func solveSudoku(grid *sudokuGrid) bool {
	x, y, found := FindUnassignedCell(grid)
	if !found {
		return true
	}
	n := len(*grid)
	for num := 1; num <= n; num++ {
		isValidRow := checkRow(grid, x, y, num)
		isValidColumn := checkColumn(grid, x, y, num)
		isValidBox := checkBox(grid, x, y, num)
		if isValidBox && isValidColumn && isValidRow {
			(*grid)[x][y] = num
			solved := solveSudoku(grid)
			if solved {
				return true
			}
			(*grid)[x][y] = 0
		}
	}
	return false
}

func checkRow(grid *sudokuGrid, row, col, num int) bool {
	n := len(*grid)
	for i := 0; i < n; i++ {
		if (*grid)[row][i] == num {
			return false
		}
	}
	return true
}

func checkColumn(grid *sudokuGrid, row, col, num int) bool {
	n := len(*grid)
	for i := 0; i < n; i++ {
		if (*grid)[i][col] == num {
			return false
		}
	}
	return true
}

func checkBox(grid *sudokuGrid, row, col, num int) bool {
	var sqrt = int(math.Sqrt(float64(len(*grid))))
	var boxRowStart = row - row%sqrt
	var boxColStart = col - col%sqrt

	for r := boxRowStart; r < boxRowStart+sqrt; r++ {
		for d := boxColStart; d < boxColStart+sqrt; d++ {
			if (*grid)[r][d] == num {
				return false
			}
		}
	}
	return true
}

// FindUnassignedCell blah blah blah
func FindUnassignedCell(grid *sudokuGrid) (x, y int, found bool) {
	n := len(*grid)
	for x = 0; x < n; x++ {
		for y = 0; y < n; y++ {
			if (*grid)[x][y] == 0 {
				found = true
				return
			}
		}
	}
	return -1, -1, false
}

// PrintGrid prints the complete Sudoku grid. It only works for 9x9 grids.
func PrintGrid(grid *sudokuGrid) {
	for i, hLine := range *grid {
		if i == 0 {
			fmt.Println("  -------------------------- ")
		} else if i%3 == 0 {
			fmt.Println("   ------   ------   ------  ")
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
