package main

import "fmt"

// the sudoku we want to solve
func initModel() [9][9]int {
	// var model = [9][9]int{
	// 	{5, 8, 6, 0, 3, 1, 0, 7, 0},
	// 	{2, 0, 7, 8, 6, 0, 5, 1, 3},
	// 	{0, 1, 0, 7, 0, 5, 2, 0, 6},
	// 	{0, 2, 8, 0, 0, 4, 3, 6, 1},
	// 	{6, 0, 4, 9, 1, 3, 7, 2, 0},
	// 	{0, 3, 1, 6, 2, 0, 0, 9, 5},
	// 	{4, 0, 5, 0, 8, 2, 0, 3, 7},
	// 	{1, 7, 0, 4, 9, 6, 8, 0, 2},
	// 	{0, 6, 2, 3, 5, 0, 1, 0, 9} }

	var model = [9][9]int{
		{0, 0, 0, 0, 0, 0, 8, 0, 0},
		{0, 7, 0, 0, 4, 5, 1, 6, 0},
		{9, 1, 4, 0, 0, 3, 7, 0, 0},
		{0, 0, 7, 8, 0, 2, 0, 0, 0},
		{0, 0, 6, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 9, 8, 4},
		{0, 0, 0, 0, 6, 0, 0, 9, 8},
		{0, 0, 0, 0, 3, 0, 5, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0} }
	return model
}

// Return true, if tuple is valid
func isTupleValid(t1, t2, t3 []int) bool {
	var check = [10]int{}
	for i := 0; i < 3; i++ {
		check[t1[i]]++
		check[t2[i]]++
		check[t3[i]]++
	}
	for i := 1; i<10; i++ {
		if check[i] > 1 {
			return false
		}
	}
	return true
}

// Little helper func: Make a column from a state
func makeColumn(state [9][9]int, col int) [9]int {
	var helper = [9]int{}
	for i := 0; i < 9; i++ {
		helper[i] = state[i][col]
	}
	return helper
}

// Return true, if configuration is valid
func isValid(state [9][9]int) bool {
	// in keiner zeile, in keiner spalte, in keinem sektor 
	// kommt eine zahl zwischen 1..9 mehrfach vor

	// line checking
	for i := 0; i < 9; i++ {
		if !isTupleValid(state[i][0:3], state[i][3:6], state[i][6:9]) {
			return false
		}
	}
	// sector checking
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !isTupleValid(state[i*3][j*3:j*3+3], 
							state[i*3+1][j*3:j*3+3], 
							state[i*3+2][j*3:j*3+3]) {
				return false
			}
		}
	}
	// column checking
	// var column = [9]int{}
	for i := 0; i < 9; i++ {
		column := makeColumn(state, i)
		if !isTupleValid(column[0:3], column[3:6], column[6:9]) {
			return false
		}
	}	

	return true 
}

// Check, if the board is final (no zeros, and valid)
func isFinal(model [9][9]int) bool {
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			if model[i][j] == 0 {
				return false
			}
		}
	}
	return isValid(model)
}

// Solve a given state and produce a valid sudoko until it is final
func solve(state [9][9]int) {
	if isFinal(state) {
		fmt.Println("Found Solution")
		fmt.Println(state)
		return
	}

    for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			if state[i][j] == 0 {
				for n:=1; n<10; n++ {
					nextState := state
					nextState[i][j] = n
					if isValid(nextState) {
						solve(nextState)
					}
				}
				return
			}
		}
	}
}

func main() {
	fmt.Println("Hello Sudoku world")
	var model = initModel()
	fmt.Println(model)

	fmt.Println("Solving ...")

	solve(model)
	fmt.Println("Solved!")
}