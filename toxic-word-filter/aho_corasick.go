package toxic_word_filter

import "strings"

func isSafeContext(text string, safePhrases []string) bool {
	lowerText := strings.ToLower(text)
	for _, safe := range safePhrases {
		if safe != "" && strings.Contains(lowerText, strings.ToLower(safe)) {
			return true
		}
	}
	return false
}

// Memasukkan kata kasar ke dalam Trie
func (ac *ahoCorasick) insert(word string) {
	curr := ac.root
	for _, ch := range word {
		if _, exists := curr.children[ch]; !exists {
			curr.children[ch] = newNode()
		}
		curr = curr.children[ch]
	}
	curr.output = append(curr.output, word)
}

// Membangun Failure Links menggunakan Queue (BFS)
func (ac *ahoCorasick) buildFailureLinks() {
	queue := make([]*node, 0)

	// Level 1 Nodes: Failure link menunjuk ke Root
	for _, child := range ac.root.children {
		child.fail = ac.root
		queue = append(queue, child)
	}

	// BFS Traversing
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for ch, child := range curr.children {
			fallback := curr.fail
			for fallback != nil {
				if nextNode, exists := fallback.children[ch]; exists {
					child.fail = nextNode
					break
				}
				fallback = fallback.fail
			}

			if child.fail == nil {
				child.fail = ac.root
			}

			// Gabungkan output pattern dari failure node
			child.output = append(child.output, child.fail.output...)
			queue = append(queue, child)
		}
	}
}

// Search Function: Scan kalimat dalam O(L)
func (ac *ahoCorasick) Search(text string) []matchResult {
	textClean := strings.ToLower(text)
	results := make([]matchResult, 0)
	curr := ac.root

	for i, ch := range textClean {
		for curr != nil {
			if nextNode, exists := curr.children[ch]; exists {
				curr = nextNode
				break
			}
			curr = curr.fail
		}

		if curr == nil {
			curr = ac.root
			continue
		}

		if len(curr.output) > 0 {
			for _, pattern := range curr.output {
				results = append(results, matchResult{
					Word:     pattern,
					Position: i - len(pattern) + 1,
				})
			}
		}
	}

	return results
}

// Sanitize memeriksa apakah ada kata terlarang dalam teks
func (ac *ahoCorasick) Sanitize(text string) (string, bool) {
	matches := ac.Search(text)
	if len(matches) == 0 {
		return text, false
	}

	runes := []rune(text)
	for _, match := range matches {
		start := match.Position
		end := start + len([]rune(match.Word))

		for i := start; i < end && i < len(runes); i++ {
			runes[i] = '*'
		}
	}

	return string(runes), true
}
