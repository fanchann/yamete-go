# Yamete-go
<p align="center">
  <img src="https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExemZkOWdvbmx2NG03bWZucGJ1MTV4ZnM2MHl1bTE4bGt3a2xmcDFpOSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/l0Iy33dWjmywkCnNS/giphy.gif" alt="yamete"/>
</p>

**Yamete-Go** comes from the Japanese word **"Yamete" (やめて)**, which means **"stop"**.  
In this context, **Yamete-Go is a high-performance text censorship library** that utilizes a **Trie-based pattern matching** algorithm to detect and censor unwanted words in a text.


## Architecture Visualization

```md
Input:
- badword
- banana

Trie Visualization:
Root
 ├── b
 │   ├── a
 │   │   ├── d
 │   │   │   ├── w
 │   │   │   │   ├── o
 │   │   │   │   │   ├── r
 │   │   │   │   │   │   ├── d (end)
 │   │   │   │   │   │
 │   ├── a
 │   │   ├── n
 │   │   │   ├── a
 │   │   │   │   ├── n
 │   │   │   │   │   ├── a (end)

```

If you insert words like banana and badword into the trie, the censorship system will replace them with asterisks, as shown below:

```md
Input:
"This is a badword and banana!"

Output:
"this is a ******* and ******!"
(Note: The input is automatically converted to lowercase.)
```

## How yamete-go works?
`Yamete-Go` processes only alphabetic characters (a-z). If the input text contains numbers, those characters are ignored during processing.


example:
```md
- Input:
  4ppl3s

- Trie Visualization:
  Root
   ├── p
   │   ├── p
   │   │   ├── l
   │   │   │   ├── s (end)

- Output:
  4ppl3s -> **true** (The word is not censored because numeric characters are ignored.)
```

## How to Use `yamete-go`

`yamete-go` is a library designed to help analyze and censor text based on predefined toxic word lists. Below are the steps to use it effectively.

---

### 1. Create a Yamete Configuration

To start using `yamete-go`, you need to create a configuration object (`YameteConfig`) that specifies the source of the toxic word list. You can load the word list either from a URL or a local file.

```go
yameteCfg := yamete.YameteConfig{
    URL:  "", // URL of the file to be loaded (e.g., a raw GitHub link)
    File: "", // File path of the file to be loaded (local file path)
}
```

**Note:**  
- If you load the word list from a URL, ensure that the raw text is UTF-8 encoded.  
- Example of a valid URL: [https://raw.githubusercontent.com/fanchann/toxic-word-list/master/id_toxic_371.txt](https://raw.githubusercontent.com/fanchann/toxic-word-list/master/id_toxic_371.txt)

---

### 2. Initialize Yamete

Once the configuration is set, initialize the `yamete-go` instance by passing the configuration object to the `NewYamete` function.

```go
yameteInit, err := yamete.NewYamete(&yameteCfg)
if err != nil {
    panic(err) // Handle errors appropriately in your application
}
```
---

### 3. Analyze Text with Yamete

After initializing `yamete-go`, you can analyze any text using the `AnalyzeText` method. This method returns detailed information about the analyzed text, including the original text, censored text, detected toxic words, and the count of censored words.

```go
response := yameteInit.AnalyzeText("dasar lu bot!")

// Print the response details
fmt.Printf("Original Text: %v\n", response.OriginalText)   // Output: dasar lu bot!
fmt.Printf("Censored Text: %v\n", response.CensoredText)   // Output: dasar lu ***!
fmt.Printf("Censored Words: %v\n", response.CensoredWords) // Output: [bot]
fmt.Printf("Censored Count: %v\n", response.CensoredCount) // Output: 1
```

---

### Key Response Fields

Here’s a breakdown of the fields returned by the `AnalyzeText` method:

- **`OriginalText`**: The original input text provided for analysis.
- **`CensoredText`**: The text after censoring toxic words (toxic words are replaced with `***`).
- **`CensoredWords`**: A list of toxic words detected in the text.
- **`CensoredCount`**: The total number of toxic words detected.




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