package sort

import (
	"cmp"
	"slices"

	"github.com/hbesfb/bedfusion/internal/bed"
)

// Lexicographic sorting, based on chromosome, start and stop positions
func lexicographicSort(lines []bed.Line) []bed.Line {
	slices.SortFunc(lines, func(a, b bed.Line) int {
		return cmp.Or(
			cmp.Compare(a.Chr, b.Chr),
			cmp.Compare(a.Start, b.Start),
			cmp.Compare(a.Stop, b.Stop),
		)
	})
	return lines
}
