package yametego

type YameteConfig struct {
	File string // load from txt file
	URL  string // load from url, note: must raw content!
}

type TextCensorshipResult struct {
	OriginalText  string
	CensoredText  string
	CensoredCount int
	CensoredWords []string
}
