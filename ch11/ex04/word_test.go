package word

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}

	// randomに空白と句読点を入れる
	for _, r := range []rune{' ', ',', '.'} {
		var position int
		if len(runes) == 0 {
			break
		} else if len(runes) == 1 {
			position = 0
		} else {
			position = rng.Intn(len(runes) - 1)
		}
		runes = append(runes[:position+1], runes[position:]...)
		runes[position] = r
	}

	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
