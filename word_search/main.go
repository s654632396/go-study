package main

import (
	"fmt"
	"log"
)

func main() {
	//board := [][]byte{
	//	[]byte{'A', 'B', 'C', 'E'},
	//	[]byte{'S', 'F', 'C', 'S'},
	//	[]byte{'A', 'D', 'E', 'E'},
	//}
	//board2 := [][]byte{
	//	[]byte{'a', 'a', 'a', 'a'},
	//	[]byte{'a', 'a', 'a', 'a'},
	//	[]byte{'a', 'a', 'a', 'a'},
	//}
	board3 := [][]byte{
		[]byte{'C', 'A', 'A'},
		[]byte{'A', 'A', 'A'},
		[]byte{'B', 'C', 'D'},
	}
	fmt.Println(
		//exist(board, "ABCCED"),
		//exist(board, "SEE"),
		//exist(board, "ABCB"),
		//exist(board2, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		exist(board3, "AAB"),
	)

}

func exist(board [][]byte, word string) bool {
	// 寻找root
	var root byte = word[0]
	var rootIndex = make([][]int, 0)
	for i, r := range board {
		for j, b := range r {
			if b == root {
				rootIndex = append(rootIndex, []int{i, j})
			}
		}
	}
	var res bool
	for _, ri := range rootIndex {
		if search(board, word, 1, ri[0], ri[1], "root") {
			res = true
			break
		}
		//debugBoard(board)
	}

	return res
}

func search(board [][]byte, w string, wi, i, j int, direct string) bool {
	if len(w)-1 < wi {
		return true
	}
	var swap byte
	if direct != "up" && i+1 < len(board) && board[i+1][j] == w[wi] {
		log.Println("down")
		board[i][j],swap = '@', board[i][j] //

		// 下
		if wi == len(w)-1 {
			// 匹配成功
			return true
		} else {
			if search(board, w, wi+1, i+1, j, "down") {
				return true
			} else {
				board[i][j] = swap
			}
		}
	}
	if direct != "down" && i-1 >= 0 && board[i-1][j] == w[wi] {
		log.Println("up")
		board[i][j],swap = '@', board[i][j] //

		// 上
		if wi == len(w)-1 {
			// 匹配成功
			return true
		} else {
			if search(board, w, wi+1, i-1, j, "up") {
				return true
			}else {
				board[i][j] = swap
			}
		}
	}
	if direct != "left" && j+1 < len(board[0]) && board[i][j+1] == w[wi] {
		log.Println("right")
		board[i][j],swap = '@', board[i][j] //
		// 右
		if wi == len(w)-1 {
			// 匹配成功
			return true
		} else {
			if search(board, w, wi+1, i, j+1, "right") {
				return true
			}else {
				board[i][j] = swap
			}
		}
	}
	if direct != "right" && j-1 >= 0 && board[i][j-1] == w[wi] {
		log.Println("left")
		board[i][j],swap = '@', board[i][j] //
		// 左
		if wi == len(w)-1 {
			// 匹配成功
			return true
		} else {
			if search(board, w, wi+1, i, j-1, "left") {
				return true
			}else {
				board[i][j] = swap
			}
		}
	}

	return false
}

func debugBoard(board [][]byte) {
	for _, a := range board {
		for _, b := range a {
			fmt.Printf("%c\t", b)
		}
		fmt.Println()
	}
	fmt.Println()
}
