# Yamete-go
![yamete](https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExemZkOWdvbmx2NG03bWZucGJ1MTV4ZnM2MHl1bTE4bGt3a2xmcDFpOSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/l0Iy33dWjmywkCnNS/giphy.gif)


**A high-performance text censorship library** with _trie-based pattern matching_ algorithm.

## Architecture Visualization

```md
Inserted words: 
-bad
-crap
-bastard

Trie:
`<-` = end of the word
        (root)
       /      \
      c        b
     /          \
    r            a
   /  \          / \
  a <- p <-     s <- t <-
          \          \
           p          a
            \          \
             e <-       r
			 			 \
						  d <-
Search:
bastard -> true
bad -> true
crap -> true

Censored words:
bastard -> ******
bad -> ***
crap -> ****
```


## Example

```go
package main

import (
	"fmt"

	yamete "github.com/fanchann/yamete-go"
)

func main() {
	yameteCfg := yamete.YameteConfig{
		URL: "https://raw.githubusercontent.com/fanchann/toxic-word-list/refs/heads/master/id_toxic_371.txt",
		// File: "files/id_words_toxic.txt",
	}
	yameteInit, err := yamete.NewYamete(&yameteCfg)
	if err != nil {
		panic(err)
	}

	response := yameteInit.AnalyzeText("dasar lu bot!")

	fmt.Printf("response.OriginalText: %v\n", response.OriginalText) // dasar lu bot!
	fmt.Printf("response.CensoredText: %v\n", response.CensoredText) // dasar lu ***!
	fmt.Printf("response.CensoredWords: %v\n", response.CensoredWords) // [bot]
	fmt.Printf("response.CensoredCount: %v\n", response.CensoredCount) // 1
}

```

## Installation

```bash
go get github.com/fanchann/yamete-go
```