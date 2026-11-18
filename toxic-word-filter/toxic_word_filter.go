package toxic_word_filter

import (
	"context"
	"fmt"
)

type ToxicWordFilterReq struct {
	Text string
}

// BadWordCacheProvider interface untuk mengambil data dari cache/repository
type BadWordCacheProvider interface {
	GetBadWordCollection(ctx context.Context) ([]string, error)
	GetSafePhaseCollection(ctx context.Context) ([]string, error)
}

// ValidateText mengambil data dari cache, lalu menjalankan filter teks
func ToxicWordFilter(ctx context.Context, cache BadWordCacheProvider, text string) (bool, error) {
	badWords, err := cache.GetBadWordCollection(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to load bad words cache: %w", err)
	}

	safePhrases, err := cache.GetSafePhaseCollection(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to load safe phrases cache: %w", err)
	}

	// 1. Bypass Check: Allowlist / Safe Context
	if isSafeContext(text, safePhrases) {
		return true, nil
	}

	// 2. Perform Fast Aho-Corasick Check
	ac := NewAhoCorasick(badWords)
	_, isToxic := ac.Sanitize(text)

	if isToxic {
		return false, nil
	}

	return true, nil
}
