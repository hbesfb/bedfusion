package bed

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/maruel/natural"
)

// Global sorting function
// Note: mergeSort() is missing from this list as it
// is only intended for internal use
func (bf *Bedfile) Sort() error {
	switch bf.SortType {
	case "lex":
		bf.Lines = lexicographicSort(bf.Lines)
	case "nat":
		bf.Lines = naturalSort(bf.Lines)
	case "ccs", "fidx":
		bf.Lines = customChrSort(bf.Lines, bf.chrOrderMap)
	default:
		return fmt.Errorf("unknown sorting type %s", bf.SortType)
	}
	return nil
}

// Lexicographic sorting
// Sorting hierarchy: chr, start, stop, strand, feat
// Chr sorting: 1 < 10 < 2 < MT < X
func lexicographicSort(lines []Line) []Line {
	slices.SortStableFunc(lines, func(a, b Line) int {
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
// Sorting hierarchy: chr, start, stop, strand, feat
// Chr sorting: 1 < 2 < 10 < MT < X
func naturalSort(lines []Line) []Line {
	slices.SortStableFunc(lines, func(a, b Line) int {
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

// Custom chromosome sorting
// Sorting hierarchy: chr, start, stop, strand, feat
// Sorting chromosomes according to custom order map
// chromosomes not in the map will be put on the bottom
// of the lines in a natural sorting order
func customChrSort(lines []Line, orderMap map[string]int) []Line {
	slices.SortStableFunc(lines, func(a, b Line) int {
		return cmp.Or(
			stringMapCompare(a.Chr, b.Chr, orderMap),
			cmp.Compare(a.Start, b.Start),
			cmp.Compare(a.Stop, b.Stop),
			cmp.Compare(a.Strand, b.Strand),
			naturalStringCompare(a.Feat, b.Feat),
		)
	})
	return lines
}

// Sorting used before merging
// Sorting hierarchy: feat, chr, strand, start, stop
// Chr sorting: 1 < 10 < 2
func mergeSort(lines []Line) []Line {
	slices.SortStableFunc(lines, func(a, b Line) int {
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

// Compares string using a predefined order contained in a map
//
// Strings that are not in the map with be ranked as greater
// than the strings in the map. If neither a or b are in the map
// they will be compared using naturalStringCompare.
//
//	-1 if a is less than b
//	 0 if a equals b
//	+1 if a is greater than b
func stringMapCompare(a, b string, order map[string]int) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if order[a] != 0 && order[b] != 0 {
		return cmp.Compare(order[a], order[b])
	}
	if order[a] != 0 && order[b] == 0 {
		return -1
	}
	if order[b] != 0 && order[a] == 0 {
		return 1
	}
	return naturalStringCompare(a, b)
}
