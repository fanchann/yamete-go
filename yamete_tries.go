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

func newYameteTrie() iYameteTrie {
	return &yameteTrie{
		root: &yameteTrieNode{},
		pool: sync.Pool{
			New: func() interface{} {
				node := &yameteTrieNode{}
				node.isEnd = false
				for i := range node.children {
					node.children[i] = nil
				}
				return node
			},
		},
	}
}

func charToIndex(c rune) int {
	return int(c - 'a')
}

func (y *yameteTrie) insert(word string) {
	y.mu.Lock()
	defer y.mu.Unlock()

	word = strings.ToLower(word)
	node := y.root

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
	word = strings.ToLower(word)
	node := y.root

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

func countWords(root *yameteTrieNode) int {
	stack := []*yameteTrieNode{root}
	count := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.isEnd {
			count++
		}

		for i := alphabetSize - 1; i >= 0; i-- {
			if node.children[i] != nil {
				stack = append(stack, node.children[i])
			}
		}
	}
	return count
}

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
	i := 0

	for i < n {
		maxEnd := -1
		currentNode := y.root

		for j := i; j < n; j++ {
			char := chars[j]
			if char < 'a' || char > 'z' {
				break
			}

			index := charToIndex(char)
			if currentNode.children[index] == nil {
				break
			}

			currentNode = currentNode.children[index]
			if currentNode.isEnd {
				maxEnd = j
			}
		}

		if maxEnd != -1 {
			toxicWord := string(chars[i : maxEnd+1])
			words = append(words, toxicWord)
			wordCensoredCount++

			for k := i; k <= maxEnd; k++ {
				censored[k] = '*'
			}
			i = maxEnd + 1
		} else {
			i++
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
