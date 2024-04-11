package trie

type node struct {
	data [26]*node
	end  bool
}

type Trie struct {
	root *node

	// size is the number of words stored in the Trie
	size int
}

func NewTrie(words ...string) *Trie {
	t := Trie{
		root: &node{},
	}

	for i := range words {
		t.Insert(words[i])
	}
	return &t
}

// normalizeCharIdx ensures that the index of a given character is between 0 and 25.
func (t *Trie) normalizeCharIdx(ch byte) byte {
	return ch - 'a'
}

func (t *Trie) reverseCharIdx(idx byte) byte {
	return idx + 'a'
}

func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) Clear() {
	t.root = &node{}
	t.size = 0
}

func (t *Trie) Insert(word string) {
	n := t.root
	for i := range word {
		idx := t.normalizeCharIdx(word[i])
		if n.data[idx] == nil {
			n.data[idx] = &node{}
		}
		n = n.data[idx]
	}
	n.end = true
	t.size++
}

func (t *Trie) Search(word string) bool {
	n := t.root
	for i := range word {
		idx := t.normalizeCharIdx(word[i])
		if next := n.data[idx]; next != nil {
			n = next
			continue
		}
		return false
	}
	return n.end
}

func (t *Trie) StartsWith(prefix string) bool {
	n := t.root
	for i := range prefix {
		idx := t.normalizeCharIdx(prefix[i])
		if next := n.data[idx]; next != nil {
			n = next
			continue
		}
		return false
	}
	return !n.end
}

func (t *Trie) Delete(word string) {
	// TODO: Implement this. I don't actually know how to do this.
}

func (t *Trie) GetAllWords() []string {
	var res []string
	t.getAllWords(t.root, "", &res)
	return res
}

func (t *Trie) getAllWords(node *node, prefix string, result *[]string) {
	if node == nil {
		return
	}

	if node.end {
		*result = append(*result, prefix)
	}

	for i := range node.data {
		t.getAllWords(node.data[i], prefix+string(t.reverseCharIdx(byte(i))), result)
	}
}

func (t *Trie) GetAllWordsWithPrefix(prefix string) []string {
	n := t.root
	for i := range prefix {
		idx := t.normalizeCharIdx(prefix[i])
		if next := n.data[idx]; next != nil {
			n = next
			continue
		}
		return nil
	}

	var res []string
	t.getAllWords(n, prefix, &res)
	return res
}
