package main

import (
    "fmt"
    "os"
)

func main() {
    // Step 1: Parsing Input
    args := os.Args[1:]
    if len(args) != 9 {
        fmt.Println("Error: Invalid number of rows. Expected 9 rows.")
        return
    }

    // Step 2: Validation
    for _, row := range args {
        if len(row) != 9 {
            fmt.Println("Error: Each row should contain exactly 9 characters.")
            return
        }
        for _, ch := range row {
            if ch != '.' && (ch < '1' || ch > '9') {
                fmt.Println("Error: Invalid character detected in the input.")
                return
            }
        }
    }

    // Step 3: Convert Input to Data Structure
    sudoku := make([][]rune, 9)
    for i, row := range args {
        sudoku[i] = []rune(row)
    }

    // Step 4: Solve the Sudoku
    if !solveSudoku(sudoku) {
        fmt.Println("Error: No solution found for the Sudoku puzzle.")
        return
    }

    // Step 5: Output
    printSudoku(sudoku)
}

// Helper function to print the solved Sudoku puzzle
func printSudoku(sudoku [][]rune) {
    for _, row := range sudoku {
        for _, ch := range row {
            fmt.Printf("%c ", ch)
        }
        fmt.Println()
    }
}

// Helper function to check if a digit is valid to be placed in a given position
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
    for i := startRow; i < startRow+3; i++ {
        for j := startCol; j < startCol+3; j++ {
            if sudoku[i][j] == digit {
                return false
            }
        }
    }

    return true
}

// Main backtracking function to solve the Sudoku puzzle
func solveSudoku(sudoku [][]rune) bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if sudoku[i][j] == '.' {
                for digit := '1'; digit <= '9'; digit++ {
                    if isValid(sudoku, i, j, digit) {
                        sudoku[i][j] = digit
                        if solveSudoku(sudoku) {
                            return true
                        }
                        sudoku[i][j] = '.' // Backtrack
                    }
                }
                return false // No valid digit found for this cell
            }
        }
    }
    return true // Sudoku solved successfully
}
