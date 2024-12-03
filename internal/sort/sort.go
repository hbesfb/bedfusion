package sort

import (
	"cmp"
	"slices"

	"github.com/hbesfb/bedfusion/internal/bed"
)

// Lexicographic sorting
// Sorting order: chr, start, stop, strand, feat
func lexicographicSort(lines []bed.Line) []bed.Line {
	slices.SortFunc(lines, func(a, b bed.Line) int {
		return cmp.Or(
			cmp.Compare(a.Chr, b.Chr),
			cmp.Compare(a.Start, b.Start),
			cmp.Compare(a.Stop, b.Stop),
			cmp.Compare(a.Strand, b.Strand),
			cmp.Compare(a.Feat, b.Feat),
		)
	})
	return lines
}

// Sorting used before merging
// Sorting order: feat, chr, strand, start, stop
func mergeSort(lines []bed.Line) []bed.Line {
	slices.SortFunc(lines, func(a, b bed.Line) int {
		return cmp.Or(
			cmp.Compare(a.Feat, b.Feat),
			cmp.Compare(a.Chr, b.Chr),
			cmp.Compare(a.Strand, b.Strand),
			cmp.Compare(a.Start, b.Start),
			cmp.Compare(a.Stop, b.Stop),
		)
	})
	return lines
}
