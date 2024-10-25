package common

import (
	"math/rand"
	"time"
)
const N = 9
type Sudoku [N][N]int

func (s *Sudoku) AddSudoku()(Sudoku) {
// 	// 数独
	// var sudoku Sudoku
// 	// ランダム
	if s.assignNum() {
		return *s
	}
	return Sudoku{} 
}

func (board *Sudoku) assignNum() (bool) {
	// ランダムなシードを設定
	rand.Seed(time.Now().UnixNano())
	return board.solve(0, 0)
}
func (board *Sudoku) solve(row, col int) bool {
	if row == N {
		return true
	}

	if col == N {
		return board.solve(row+1, 0)
	}

	if board[row][col] != 0 {
		return board.solve(row, col+1)
	}

	// 1から9の候補数字を格納するスライス
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	for _, num := range numbers {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if board.solve(row, col+1) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

// 数字が行、列、ブロック内で有効かどうかをチェックする関数
func isValid(board *Sudoku, row, col, num int) bool {
	for i := 0; i < N; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) RemoveDigits(k int) {
	count := k
	for count != 0 {
			cellId := rand.Intn(N * N)
			row := cellId / N
			col := cellId % N
			if s[row][col] != 0 {
					count--
					s[row][col] = 0
			}
	}
}

// [N][N]int を [][]int に変換する関数
func (sudoku Sudoku) ToSlice() [][]int {
	// [][]intの初期化
	slice := make([][]int, N)
	for i := range slice {
		slice[i] = make([]int, N)
	}

	// [N][N]int の内容を [][]int にコピー
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			slice[i][j] = sudoku[i][j]
		}
	}
	return slice
}