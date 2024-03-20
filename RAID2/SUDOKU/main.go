package main

import (
	"fmt"
	"os"
)

// checkParams checks if the input parameters are valid for the Sudoku.
func checkParams(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for _, row := range args {
		if len(row) != 9 {
			return false
		}
		seen := make([]bool, 10) // To track numbers 1 to 9
		for _, ch := range row {
			if ch != '.' && (ch < '1' || ch > '9') {
				return false
			}
			if ch != '.' {
				digit := ch - '0' // Convert rune to integer
				if seen[digit] {
					return false // check repeated number in the same row
				}
				seen[digit] = true
			}
		}
	}
	return true
}

// prints the Sudoku.
func printSudoku(sudoku [][]rune) {
	for _, row := range sudoku {
		for i, ch := range row {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}

// isValid checks if placing 'digit' at position (row, col) is valid in the Sudoku grid.
func isValid(sudoku [][]rune, row, col int, digit rune) bool {
	// Check row
	for i := 0; i < 9; i++ {
		if sudoku[row][i] == digit {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if sudoku[i][col] == digit {
			return false
		}
	}

	// Check 3x3 subgrid
	startRow, startCol := (row/3)*3, (col/3)*3
	for y := startRow; y < startRow+3; y++ {
		for x := startCol; x < startCol+3; x++ {
			if sudoku[y][x] == digit {
				return false
			}
		}
	}

	return true
}

// solveSudoku recursively solves the Sudoku puzzle using backtracking.
func solveSudoku(sudoku [][]rune) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if sudoku[y][x] == '.' {
				for digit := '1'; digit <= '9'; digit++ {
					if isValid(sudoku, y, x, digit) {
						sudoku[y][x] = digit
						if solveSudoku(sudoku) {
							return true
						}
						sudoku[y][x] = '.' // Backtrack
					}
				}
				return false // No valid digit found for this cell
			}
		}
	}
	return true // Sudoku solved successfully
}

// resolveSudoku initializes the grid and starts solving the Sudoku puzzle.
func resolveSudoku(args []string) (bool, [][]rune) {
	var sudoku [][]rune
	for _, row := range args {
		sudoku = append(sudoku, []rune(row))
	}
	if solveSudoku(sudoku) {
		return true, sudoku
	}
	return false, nil
}

func main() {
	args := os.Args
	if len(args) != 10 || !checkParams(args[1:]) {
		fmt.Println("Error")
		return
	}

	solved, sudoku := resolveSudoku(args[1:])
	if !solved {
		fmt.Println("Error")
		return
	}

	printSudoku(sudoku)
}
