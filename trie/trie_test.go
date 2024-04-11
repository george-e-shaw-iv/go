package trie_test

import (
	"testing"

	"github.com/george-e-shaw-iv/go/trie"
	"github.com/stretchr/testify/assert"
)

func TestTrie_Search(t *testing.T) {
	tr := trie.NewTrie("foo", "bar", "baz")

	assert.True(t, tr.Search("bar"))
	assert.False(t, tr.Search("quack"))
	assert.False(t, tr.Search("foop"))
}

func TestTrie_StartsWith(t *testing.T) {
	tr := trie.NewTrie("foo", "bar", "baz")

	assert.True(t, tr.StartsWith("ba"))
	assert.False(t, tr.StartsWith("fe"))
	assert.False(t, tr.StartsWith("baz"))
}

func TestTrie_GetAllWords(t *testing.T) {
	tr := trie.NewTrie("foo", "bar", "baz")

	words := tr.GetAllWords()
	assert.Equal(t, 3, len(words))
	assert.Contains(t, words, "foo")
	assert.Contains(t, words, "bar")
	assert.Contains(t, words, "baz")
}

func TestTrie_GetAllWordsWithPrefix(t *testing.T) {
	tr := trie.NewTrie("foo", "bar", "baz")

	words := tr.GetAllWordsWithPrefix("ba")
	assert.Equal(t, 2, len(words))
	assert.Contains(t, words, "bar")
	assert.Contains(t, words, "baz")
}
