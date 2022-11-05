// https://leetcode.com/problems/word-search-ii/

package p212

const (
	a           = 'a'
	moveTypeCnt = 4
)

var (
	nextX = []int{-1, 0, 1, 0}
	nextY = []int{0, 1, 0, -1}
)

type TrieNode struct {
	Children [26]*TrieNode
	Word     string
}

func (t *TrieNode) insert(word string) {
	for _, ch := range word {
		idx := ch - a
		if t.Children[idx] == nil {
			t.Children[idx] = new(TrieNode)
		}
		t = t.Children[idx]
	}
	t.Word = word
}

func findWords(board [][]byte, words []string) []string {
	tn := &TrieNode{}
	for _, word := range words {
		tn.insert(word)
	}
	var result = make([]string, 0, len(words))
	for i := range board {
		for j := range board[i] {
			idx := board[i][j] - a
			if tmp := tn.Children[idx]; tmp != nil {
				dfs(board, i, j, tn, &result)
			}
		}
	}
	return result
}

func dfs(board [][]byte, i, j int, node *TrieNode, result *[]string) {
	c := board[i][j]
	node = node.Children[c-a]
	if node == nil {
		return
	} else if node.Word != "" {
		*result = append(*result, node.Word)
		node.Word = ""
	}

	board[i][j] = '.'
	for k := 0; k < moveTypeCnt; k++ {
		x, y := i+nextX[k], j+nextY[k]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] != '.' {
			dfs(board, x, y, node, result)
		}
	}
	board[i][j] = c
}
