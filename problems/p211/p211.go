package p211

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
	}
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: NewTrieNode(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for i := range word {
		c := word[i]
		if curr.children[c] == nil {
			curr.children[c] = NewTrieNode()
		}

		curr = curr.children[c]
	}
	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(word, this.root)
}

func dfs(word string, root *TrieNode) bool {
	curr := root

	for i := range word {
		c := word[i]
		if c == '.' {
			for _, child := range curr.children {
				if dfs(word[i+1:], child) {
					return true
				}
			}
			return false
		} else {
			if _, ok := curr.children[c]; !ok {
				return false
			}

			curr = curr.children[c]
		}
	}
	return curr.isEnd
}
