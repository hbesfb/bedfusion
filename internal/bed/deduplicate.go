package bed

import (
	"sort"
	"strings"

	"github.com/maruel/natural"
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

// Naturally sort and deduplicate items in slice of strings
func sortAndDeduplicateListOfStrings(list natural.StringSlice) []string {
	// Sort list
	sort.Sort(list)
	// Deduplicate
	j := 1
	for i := 1; i < len(list); i++ {
		if list[i] == list[i-1] {
			continue
		}
		list[j] = list[i]
		j++
	}
	// Since we overwrote the beginning of the list
	// we only want to return the parts we have
	// overwritten
	return list[:j]
}
