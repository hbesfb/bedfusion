package sorting

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/maruel/natural"

	"github.com/hbesfb/bedfusion/internal/bed"
)

// Note that the the user will give the columns with 1-based indexing,
// but that we convert this to zero-based indexing in .VerifyAndHandle()
type Config struct {
	SortingType string `env:"SORTING_TYPE" enum:"lex,nat" default:"lex" short:"s" help:"How the bed files should be sorted. lex = lexicographic sorting (chr: 1 < 10 < 2), nat = natural sorting (chr: 1 < 2 < 10)"`
}

// Global sorting function
func (c Config) Sort(lines []bed.Line) ([]bed.Line, error) {
	var sortedLines []bed.Line
	switch c.SortingType {
	case "lex":
		sortedLines = lexicographicSort(lines)
	case "nat":
		sortedLines = naturalSort(lines)
	case "merge":
		sortedLines = mergeSort(lines)
	default:
		return []bed.Line{}, fmt.Errorf("unknown sorting type %s", c.SortingType)
	}
	return sortedLines, nil
}

// Lexicographic sorting
// Sorting order: chr, start, stop, strand, feat
// Chr sorting: 1 < 10 < 2
func lexicographicSort(lines []bed.Line) []bed.Line {
	slices.SortStableFunc(lines, func(a, b bed.Line) int {
		return cmp.Or(
			cmp.Compare(strings.ToLower(a.Chr), strings.ToLower(b.Chr)),
			cmp.Compare(a.Start, b.Start),
			cmp.Compare(a.Stop, b.Stop),
			cmp.Compare(a.Strand, b.Strand),
			cmp.Compare(strings.ToLower(a.Feat), strings.ToLower(b.Feat)),
		)
	})
	return lines
}

// Natural sorting
// Sorting order: chr, start, stop, strand, feat
// Chr sorting: 1 < 2 < 10 (Features are also sorted the same way)
func naturalSort(lines []bed.Line) []bed.Line {
	slices.SortStableFunc(lines, func(a, b bed.Line) int {
		return cmp.Or(
			naturalStringCompare(a.Chr, b.Chr),
			cmp.Compare(a.Start, b.Start),
			cmp.Compare(a.Stop, b.Stop),
			cmp.Compare(a.Strand, b.Strand),
			naturalStringCompare(a.Feat, b.Feat),
		)
	})
	return lines
}

// Sorting used before merging
// Sorting order: feat, chr, strand, start, stop
// Chr sorting: 1 < 10 < 2
func mergeSort(lines []bed.Line) []bed.Line {
	slices.SortStableFunc(lines, func(a, b bed.Line) int {
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

// Natural comparison of strings
//
//	-1 if a is less than b
//	 0 if a equals b
//	+1 if a is greater than b
func naturalStringCompare(a, b string) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if a == b {
		return 0
	}
	if natural.Less(a, b) {
		return -1
	}
	return 1
}
