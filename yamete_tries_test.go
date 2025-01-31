package yametego

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYameteInsert(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc:     "insert word",
			input:    "apple",
			expected: true,
		},
		{
			desc:     "insert word with numeric",
			input:    "b4n4n4",
			expected: true,
		},
		{
			desc:     "insert null string",
			input:    "",
			expected: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			y := newYameteTrie()
			y.insert(tC.input)

			result := y.searchText(tC.input)
			if result != tC.expected {
				t.Errorf("Test '%s' failed: expected %v, got %v", tC.desc, tC.expected, result)
			}
		})
	}
}

func TestYameteSearchWord(t *testing.T) {
	yg := newYameteTrie()

	yg.insert("foo")
	yg.insert("foobar")

	censoredWord, _, _ := yg.censorText("foobarxyz")
	require.Contains(t, censoredWord, "xyz")
}

func TestCensorText(t *testing.T) {
	y := newYameteTrie()
	y.insert("badword")
	censored, count, words := y.censorText("This is a badword example.")
	require.Equal(t, "this is a ******* example.", censored)
	require.Equal(t, 1, count)
	require.Equal(t, []string{"badword"}, words)
}

func TestGetTtlInserted(t *testing.T) {
	y := newYameteTrie()

	ttlTrie := y.getAllTextTtl()
	require.Equal(t, 0, ttlTrie)

	y.insert("grapes")
	y.insert("banana")
	y.insert("apples")
	ttlTrieNew := y.getAllTextTtl()
	require.Equal(t, 3, ttlTrieNew)

}
