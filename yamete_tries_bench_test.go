package yametego

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var (
	trie        = newYameteTrie()
	hundredWord = []string{
		"badword", "hate", "kill", "stupid", "idiot", "loser", "dumb", "ugly", "horrible", "terrible",
		"awful", "nasty", "rude", "insult", "mock", "bully", "abuse", "violence", "threat", "curse",
		"swear", "offensive", "vulgar", "obscene", "harass", "torment", "torture", "aggression", "hostile",
		"malicious", "spiteful", "vengeful", "cruel", "brutal", "savage", "evil", "wicked", "sinister",
		"corrupt", "deceit", "fraud", "scam", "cheat", "betray", "backstab", "manipulate", "exploit", "oppress",
		"dominate", "control", "dictate", "tyrant", "despot", "fascist", "racist", "bigot", "prejudice", "discriminate",
		"xenophobia", "homophobia", "misogyny", "sexism", "chauvinist", "nazi", "terrorist", "extremist", "radical", "fanatic",
		"zealot", "heretic", "blasphemy", "slander", "libel", "defame", "smear", "gossip", "rumor", "lie",
		"falsehood", "deception", "hypocrisy", "coward", "traitor", "rebel", "riot", "chaos", "anarchy", "revolt",
		"mutiny", "treason", "felony", "misdemeanor", "crime", "illegal", "unlawful", "violate", "infringe", "trespass",
		"plagiarize", "pirate", "hack", "phish", "scam", "fraudulent", "counterfeit", "forgery", "embezzle", "steal",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRandomText(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

/*
benchmark #01
Total inserted words: 121059
Memory allocated: 261284424 bytes (261.284424 MB)
*/
func BenchmarkInsert(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		trie.insert(generateRandomText(15))
	}

	b.StopTimer()

	ttlInsertedWord := trie.getAllTextTtl()
	b.Logf("Total inserted words: %d", ttlInsertedWord)
	memAllocatedLog(b)
}

/*
Memory allocated 447512 bytes (437 KB)
BenchmarkSearchText-8             187056              6123 ns/op
*/
func BenchmarkSearchText(b *testing.B) {
	b.ReportAllocs()
	y := newYameteTrie()
	for _, word := range hundredWord {
		y.insert(word)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, word := range hundredWord {
			y.searchText(word)
		}
	}

	memAllocatedLog(b)
}

/*
Memory allocated: 1430880 bytes (1.36 MB)
BenchmarkCensorText-8             281223              6314 ns/op
*/
func BenchmarkCensorText(b *testing.B) {

	y := newYameteTrie()
	for _, word := range hundredWord {
		y.insert(word)
	}

	text := "You are such an idiot and a complete loser, always making dumb mistakes."
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		y.censorText(text)
	}

	memAllocatedLog(b)
}

func memAllocatedLog(b *testing.B) {
	var memBefore, memAfter runtime.MemStats
	runtime.ReadMemStats(&memBefore)

	b.Helper()

	runtime.ReadMemStats(&memAfter)
	memoryAlloc := memAfter.Alloc - memBefore.Alloc

	b.Logf("Memory allocated: %d bytes", memoryAlloc)
}
