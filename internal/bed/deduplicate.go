package bed

import (
	"strings"
)

// Remove duplicated lines
func (bf *Bedfile) DeduplicateLines() {
	var deduplicatedLines []Line
	seen := map[string]bool{}
	for _, line := range bf.Lines {
		joinedLine := strings.Join(line.Full, ",")
		if !seen[joinedLine] {
			seen[joinedLine] = true
			deduplicatedLines = append(deduplicatedLines, line)
		}
	}
	bf.Lines = deduplicatedLines
}

// Remove duplicated strings in slice
func deduplicateListOfStrings(list []string) []string {
	var deduplicatedList []string
	seen := map[string]bool{}
	for _, i := range list {
		if !seen[i] {
			seen[i] = true
			deduplicatedList = append(deduplicatedList, i)
		}
	}
	return deduplicatedList
}
