package main

import (
	"fmt"
	"os"
)

// CheckInput checks if the input parameters are valid for the Sudoku puzzle.
func CheckInput(args []string) bool {
	// Check if there are exactly 10 arguments (including the command name).
	if len(args) != 10 {
		return false
	}
	// Iterate over each row of the Sudoku puzzle.
	for _, row := range args[1:] {
		// Check if each row has exactly 9 characters.
		if len(row) != 9 {
			return false
		}
		// Initialize a boolean array to track seen digits (1-9).
		seen := [10]bool{}
		// Iterate over each character in the row.
		for _, ch := range row {
			// Check if the character is not a dot and not a digit between 1 and 9.
			if ch != '.' && (ch < '1' || ch > '9') {
				return false
			}
			// If the character is not a dot, mark the corresponding digit as seen.
			if ch != '.' {
				digit := ch - '0'
				if seen[digit] {
					return false // Check for repeated digits in the same row.
				}
				seen[digit] = true
			}
		}
	}
	return true
}

// ResolveSudoku initializes the grid and starts solving the Sudoku puzzle.
func ResolveSudoku(args []string) (bool, [][]rune) {
	var sudoku [][]rune
	// Iterate over each row of the Sudoku puzzle.
	for _, row := range args[1:] {
		var runes []rune
		// Convert each character to rune and append to the runes slice.
		for _, ch := range row {
			runes = append(runes, ch)
		}
		// Append the row (as runes slice) to the sudoku grid.
		sudoku = append(sudoku, runes)
	}
	// Attempt to solve the Sudoku puzzle.
	if SolveSudoku(sudoku) {
		return true, sudoku // If solved, return true and the solved Sudoku grid.
	}
	return false, nil // If not solved, return false and nil grid.
}

// SolveSudoku recursively solves the Sudoku puzzle using backtracking.
func SolveSudoku(sudoku [][]rune) bool {
	// Iterate over each cell in the Sudoku grid.
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			// If the cell is empty (contains a dot).
			if sudoku[y][x] == '.' {
				// Try placing digits from '1' to '9' in the empty cell.
				for digit := '1'; digit <= '9'; digit++ {
					// If the placement of the current digit is valid.
					if IsValidMove(sudoku, y, x, digit) {
						// Place the digit in the cell and recursively solve the Sudoku.
						sudoku[y][x] = digit
						if SolveSudoku(sudoku) {
							return true // If Sudoku is solved, return true.
						}
						// If Sudoku cannot be solved with the current digit, backtrack by removing it.
						sudoku[y][x] = '.'
					}
				}
				return false // If no valid digit can be placed in the cell, return false.
			}
		}
	}
	return true // If all cells are filled, Sudoku is solved successfully.
}

// IsValidMove checks if placing 'digit' at position (row, col) is valid in the Sudoku grid.
func IsValidMove(sudoku [][]rune, row, col int, digit rune) bool {
	// Check row, column, and 3x3 subgrid for the presence of the digit.
	for i := 0; i < 9; i++ {
		if sudoku[row][i] == digit || sudoku[i][col] == digit {
			return false // If the digit is found in the row or column, it's invalid.
		}
	}
	// Determine the starting row and column of the 3x3 subgrid.
	startRow, startCol := (row/3)*3, (col/3)*3
	// Iterate over the cells in the 3x3 subgrid.
	for y := startRow; y < startRow+3; y++ {
		for x := startCol; x < startCol+3; x++ {
			if sudoku[y][x] == digit {
				return false // If the digit is found in the subgrid, it's invalid.
			}
		}
	}
	return true // If the digit is not found in the row, column, or subgrid, it's valid.
}

// PrintSudoku prints the Sudoku grid.
func PrintSudoku(sudoku [][]rune) {
	for _, row := range sudoku {
		for i, ch := range row {
			if i > 0 {
				fmt.Print(" ") // Print a space between each digit in the row.
			}
			fmt.Print(string(ch)) // Print the character representing the digit.
		}
		fmt.Println() // Move to the next line after printing each row.
	}
}

func main() {
	args := os.Args
	// Check if the input parameters are valid.
	if !CheckInput(args) {
		fmt.Println("Error")
		return
	}

	// Attempt to resolve the Sudoku puzzle.
	solved, sudoku := ResolveSudoku(args)
	if !solved {
		fmt.Println("Error")
		return
	}

	// Print the solved Sudoku grid.
	PrintSudoku(sudoku)
}
