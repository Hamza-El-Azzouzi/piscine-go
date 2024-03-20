package main

import (
	"fmt"
	"os"
)

// this func checks if the input parameters are valid for the Sudoku.
func CheckInput(args []string) bool {
	if len(args) != 10 {
		return false
	}
	for _, row := range args[1:] {
		if len(row) != 9 {
			return false
		}
		seen := [10]bool{} // To track numbers
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

// Initializes the grid and starts solving the Sudoku puzzle. It's called by main after the input parameters have been validated.
func ResolveSudoku(args []string) (bool, [][]rune) {
	var sudoku [][]rune
	for _, row := range args[1:] {
		var runes []rune
		for _, ch := range row {
			runes = append(runes, ch)
		}
		sudoku = append(sudoku, runes)
	}
	if SolveSudoku(sudoku) {
		return true, sudoku
	}
	return false, nil
}

// solveSudoku recursively solves the Sudoku puzzle using backtracking.
func SolveSudoku(sudoku [][]rune) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if sudoku[y][x] == '.' {
				for digit := '1'; digit <= '9'; digit++ {
					if IsValidMove(sudoku, y, x, digit) {
						sudoku[y][x] = digit
						if SolveSudoku(sudoku) {
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

// isValidMove checks if placing 'digit' at position (row, col) is valid in the Sudoku grid.
func IsValidMove(sudoku [][]rune, row, col int, digit rune) bool {
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

	// Check 3x3 box
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

// PrintSudoku prints the Sudoku grid.
func PrintSudoku(sudoku [][]rune) {
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

func main() {
	args := os.Args
	if !CheckInput(args) {
		fmt.Println("Error")
		return
	}

	solved, sudoku := ResolveSudoku(args)
	if !solved {
		fmt.Println("Error")
		return
	}

	PrintSudoku(sudoku)
}
