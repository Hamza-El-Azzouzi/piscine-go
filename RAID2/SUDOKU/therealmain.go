package main

import (
	"fmt"
	"os"
)

// CheckInput checks if the input parameters are valid for the Sudoku.
func CheckInput(args []string) bool {
	if len(args) != 10 { // Check if there are exactly 10 arguments (including the command name).
		return false
	}
	for _, row := range args[1:] {
		if len(row) != 9 { // Check if each row has exactly 9 characters.
			return false
		}
		seen := [10]bool{} // Initialize a boolean array to track seen digits (1-9).
		for _, ch := range row {
			if ch != '.' && (ch < '1' || ch > '9') {
				return false
			}
			if ch != '.' { // If the character is not a dot, mark the corresponding digit as seen.
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

// SolveSudoku recursively solves the Sudoku puzzle using backtracking.
func SolveSudoku(sudoku [][]rune) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if sudoku[y][x] == '.' {
				for digit := '1'; digit <= '9'; digit++ { // Try placing digits from '1' to '9' in the empty cell.
					if IsValidMove(sudoku, y, x, digit) {
						sudoku[y][x] = digit
						if SolveSudoku(sudoku) {
							return true
						}
                        // If Sudoku cannot be solved with the current digit, backtrack by removing it.
						sudoku[y][x] = '.' 
					}
				}
				return false 
			}
		}
	}
	return true 
}

// IsValidMove checks if placing 'digit' at position (row, col) is valid in the Sudoku grid.
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
    // Check box 3x3
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

// prints the Sudoku grid.
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
	if !CheckInput(args) { // Check if the input parameters are valid.
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