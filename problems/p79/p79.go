// https://leetcode.com/problems/word-search/

package p79

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	var m = len(board)
	var n = len(board[0])

	var checkFunc func(i, j int, word string) bool

	checkFunc = func(i, j int, word string) bool {
		if len(word) == 0 {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[0] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = '.'
		w := word[1:]
		res := checkFunc(i+1, j, w) || checkFunc(i-1, j, w) || checkFunc(i, j+1, w) || checkFunc(i, j-1, w)
		board[i][j] = tmp
		return res
	}

	for i := range board {
		for j := range board[i] {
			if checkFunc(i, j, word) {
				return true
			}
		}
	}

	return false
}
