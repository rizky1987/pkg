package toxic_word_filter

import "strings"

// node merepresentasikan satu state/karakter dalam Trie
type node struct {
	children map[rune]*node
	fail     *node    // Failure link (mirip KMP algorithm)
	output   []string // Menyimpan kata kasar yang ditemukan pada node ini
}

func newNode() *node {
	return &node{
		children: make(map[rune]*node),
		output:   make([]string, 0),
	}
}

type matchResult struct {
	Word     string
	Position int
}

// ahoCorasick Engine
type ahoCorasick struct {
	root *node
}

func NewAhoCorasick(dictionary []string) *ahoCorasick {
	ac := &ahoCorasick{root: newNode()}

	// Step 1: Build Trie
	for _, word := range dictionary {
		ac.insert(strings.ToLower(word))
	}

	// Step 2: Build Failure & Output Links (BFS)
	ac.buildFailureLinks()

	return ac
}
