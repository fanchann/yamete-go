package yametego

import (
	"strings"
	"sync"
)

const alphabetSize = 26 // Jumlah huruf dalam alfabet (a-z)

type iYameteTrie interface {
	insert(word string)
	searchText(word string) bool
	censorText(phrase string) (string, int, []string)
	textCensorshipResult(text string) *TextCensorshipResult
	getAllTextTtl() int
}

type yameteTrieNode struct {
	children [alphabetSize]*yameteTrieNode 
	isEnd    bool
}

type yameteTrie struct {
	root *yameteTrieNode
	pool sync.Pool
	mu   sync.RWMutex
}

func newYameteTrie() *yameteTrie {
	return &yameteTrie{
		root: &yameteTrieNode{},
		pool: sync.Pool{
			New: func() interface{} {
				return &yameteTrieNode{}
			},
		},
	}
}

func charToIndex(c rune) int {
	return int(c - 'a')
}

func (y *yameteTrie) insert(word string) {
	node := y.root
	word = strings.ToLower(word)
	for _, char := range word {
		if char < 'a' || char > 'z' {
			continue
		}
		index := charToIndex(char)
		if node.children[index] == nil {
			node.children[index] = y.pool.Get().(*yameteTrieNode)
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (y *yameteTrie) searchText(word string) bool {
	node := y.root
	word = strings.ToLower(word)
	for _, char := range word {
		if char < 'a' || char > 'z' {
			continue
		}
		index := charToIndex(char)
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd
}

func (y *yameteTrie) getAllTextTtl() int {
	return countWords(y.root)
}

func countWords(node *yameteTrieNode) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.isEnd {
		count++
	}
	for _, child := range node.children {
		count += countWords(child)
	}
	return count
}

// Optimized `censorText` method
func (y *yameteTrie) censorText(phrase string) (string, int, []string) {
	y.mu.RLock()
	defer y.mu.RUnlock()

	phrase = strings.ToLower(phrase)
	chars := []rune(phrase)
	n := len(chars)
	censored := make([]rune, n)
	copy(censored, chars)

	wordCensoredCount := 0
	var words []string

	for i := 0; i < n; i++ {
		node := y.root
		for j := i; j < n; j++ {
			char := chars[j]
			if char < 'a' || char > 'z' {
				break // Abaikan karakter non-alfabet
			}
			index := charToIndex(char)
			if node.children[index] == nil {
				break
			}
			node = node.children[index]
			if node.isEnd {
				toxicWord := string(chars[i : j+1])
				words = append(words, toxicWord)

				for k := i; k <= j; k++ {
					censored[k] = '*'
				}
				wordCensoredCount++
				break
			}
		}
	}

	return string(censored), wordCensoredCount, words
}

func (y *yameteTrie) textCensorshipResult(text string) *TextCensorshipResult {
	censoredText, censoredWordCount, wordsBeforeCensored := y.censorText(text)

	return &TextCensorshipResult{
		OriginalText:  text,
		CensoredText:  censoredText,
		CensoredCount: censoredWordCount,
		CensoredWords: wordsBeforeCensored,
	}
}
