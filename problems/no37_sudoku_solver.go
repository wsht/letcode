package main

import "fmt"

func check(board [][]byte, value byte, row_no int, column_no int) bool {
	//row check
	for i := 0; i < 9; i++ {
		if board[row_no][i] == value {
			return false
		}
	}
	//column check
	for i := 0; i < 9; i++ {
		if board[i][column_no] == value {
			return false
		}
	}

	//3x3 board check
	row_start := (row_no / 3) * 3
	column_start := (column_no / 3) * 3
	for i := row_start; i < row_start+3; i++ {
		for j := column_start; j < column_start+3; j++ {
			if board[i][j] == value {
				return false
			}
		}
	}

	return true
}

func solveSudoku(board [][]byte) {
	var stack [][3]int
	start_row := 0
	start_column := 0
	start_num := 1
	while := true
	for while {
		break_mark := false
		for row := start_row; row < 9; row++ {
			for column := start_column; column < 9; column++ {
				if board[row][column] != byte('.') {
					if column == 8 && row == 8 && break_mark == false {
						while = false
					}
					continue
				}
				fill := byte('.')
				for i := start_num; i <= 9; i++ {
					//转成AscII
					if check(board, byte(i+48), row, column) {
						fill = byte(i)
						break
					}
				}
				// fmt.Println(row, column, int(fill), fill == '.', start_num)
				//很有可能是题目出错，现在先不做检测
				if fill == byte('.') {
					var tmp [3]int
					tmp, stack = stack[len(stack)-1], stack[:len(stack)-1]
					start_row = tmp[0]
					start_column = tmp[1]
					start_num = int(tmp[2]) + 1

					board[start_row][start_column] = byte('.')
					break_mark = true
					break
				} else {
					stack = append(stack, [3]int{row, column, int(fill)})
					//转成AscII
					board[row][column] = fill + 48
					start_num = 1
					//这里多做了一些循环，可以想办法优化一下
					start_column = 0
					// start_row = row
				}
			}
			if break_mark {
				break
			}
		}
	}

	for _, row := range board {
		for _, val := range row {
			fmt.Print(string(val), ",")
		}
		fmt.Print("\n")
	}
}

// func main() {
// 	board := make([][]byte, 9)
// 	// s := byte('5')

// 	board[0] = []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
// 	board[1] = []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
// 	board[2] = []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
// 	board[3] = []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
// 	board[4] = []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
// 	board[5] = []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
// 	board[6] = []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
// 	board[7] = []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
// 	board[8] = []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
// 	solveSudoku(board)
// 	// solveSudoku()
// }
