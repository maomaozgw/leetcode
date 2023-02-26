package p208

type Node struct {
	m map[rune]Node
}

type Trie struct {
	trie Node
}

func Constructor() Trie {
	return Trie{
		trie: Node{
			m: map[rune]Node{},
		},
	}
}

func (this *Trie) Insert(word string) {
	trie := this.trie
	for _, l := range word {
		if _, ok := trie.m[l]; !ok {
			trie.m[l] = Node{
				m: map[rune]Node{},
			}
		}
		trie = trie.m[l]
	}
	trie.m['0'] = Node{
		m: map[rune]Node{},
	}
}

func (this *Trie) Search(word string) bool {
	trie := this.trie
	for _, l := range word {
		if _, ok := trie.m[l]; !ok {
			return false
		}
		trie = trie.m[l]
	}
	_, ok := trie.m['0']
	return ok
}

func (this *Trie) StartsWith(prefix string) bool {
	trie := this.trie
	for _, l := range prefix {
		if _, ok := trie.m[l]; !ok {
			return false
		}
		trie = trie.m[l]
	}
	return true
}
