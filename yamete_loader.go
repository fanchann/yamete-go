package yametego

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type iDictionaryLoader interface {
	LoadFromURL(url string) error
	LoadFromFile(filePath string) error
}

type trieDictionaryLoader struct {
	trie iYameteTrie
}

func newTrieDictionaryLoader(trie iYameteTrie) iDictionaryLoader {
	return &trieDictionaryLoader{trie: trie}
}

func (l *trieDictionaryLoader) LoadFromURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	return l.processInputSource(bufio.NewScanner(resp.Body))
}

func (l *trieDictionaryLoader) LoadFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return l.processInputSource(bufio.NewScanner(file))
}

func (l *trieDictionaryLoader) processInputSource(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			l.trie.insert(word)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	return nil
}
