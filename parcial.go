//Pregunta 2.a
package main

import (
	"fmt"
	"strings"
)

func numberQueens(n, row int, column, a, b []bool, board []string, ans *[][]string) {
	if row == n {
		*ans = append(*ans, append([]string{}, board...))
		return
	}
	for col := 0; col < n; col++ {
		if column[col] || a[row+col] || b[row-col+n-1] {
			continue
		}
		column[col] = true
		a[row+col] = true
		b[row-col+n-1] = true
		board[row] = board[row][:col] + "Q" + board[row][col+1:]
		numberQueens(n, row+1, column, a, b, board, ans)

		column[col] = false
		a[row+col] = false
		b[row-col+n-1] = false
		board[row] = board[row][:col] + "." + board[row][col+1:]
	}
}

func solution(n int) [][]string {
	ans := [][]string{}
	column := make([]bool, n)
	a := make([]bool, 2*n)
	b := make([]bool, 2*n)
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	numberQueens(n, 0, column, a, b, board, &ans)
	return ans
}

func main() {
	var n int
	fmt.Printf("Introduzca un número: ")
	fmt.Scanf("%d", &n)
	fmt.Print(solution(n))
}
