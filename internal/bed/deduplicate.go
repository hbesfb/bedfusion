package bed

import (
	"strings"
)

// Remove duplicates
// Requires the lines to have been sorted before use
func (bf *Bedfile) DeduplicateLines() {
	var deduplicatedLines []Line
	for i, line := range bf.Lines {
		if i != 0 && strings.Join(line.Full, ",") == strings.Join(bf.Lines[i-1].Full, ",") {
			continue
		}
		deduplicatedLines = append(deduplicatedLines, line)
	}
	bf.Lines = deduplicatedLines
}
