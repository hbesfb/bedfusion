package bed

import (
	"strings"
)

// Remove duplicates
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
