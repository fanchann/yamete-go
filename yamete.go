package yametego

import (
	"errors"
	"fmt"
)

type Yamete interface {
	AnalyzeText(text string) *TextCensorshipResult
	GetTotalOfDictionaries() int
}

type yameteImpl struct {
	trie   iYameteTrie
	loader iDictionaryLoader
}

func NewYamete(cfg *YameteConfig) (Yamete, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	trie := newYameteTrie()
	loader := newTrieDictionaryLoader(trie)

	if cfg.URL != "" {
		if err := loadFromURL(loader, cfg.URL); err != nil {
			return nil, fmt.Errorf("failed to load from URL: %w", err)
		}
	}

	if cfg.File != "" {
		if err := loadFromFile(loader, cfg.File); err != nil {
			return nil, fmt.Errorf("failed to load from file: %w", err)
		}
	}

	return &yameteImpl{
		trie:   trie,
		loader: loader,
	}, nil
}

func validateConfig(cfg *YameteConfig) error {
	if cfg.URL == "" && cfg.File == "" {
		return errors.New("at least one dictionary source (URL or file path) is required")
	}
	return nil
}

func loadFromURL(loader iDictionaryLoader, url string) error {
	return loader.LoadFromURL(url)
}

func loadFromFile(loader iDictionaryLoader, path string) error {
	return loader.LoadFromFile(path)
}

func (a *yameteImpl) AnalyzeText(text string) *TextCensorshipResult {
	return a.trie.textCensorshipResult(text)
}

func (a *yameteImpl) GetTotalOfDictionaries() int {
	return a.trie.getAllTextTtl()
}
