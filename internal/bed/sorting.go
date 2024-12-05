package bed

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/maruel/natural"
)

// Note that only lowercase is used in this slice
var humanChrOrder = []string{"1", "chr1", "2", "chr2", "3", "chr3", "4", "chr4", "5", "chr5", "6", "chr6", "7", "chr7", "8", "chr8", "9", "chr9", "10", "chr10", "11", "chr11", "12", "chr12", "13", "chr13", "14", "chr14", "15", "chr15", "16", "chr16", "17", "chr17", "18", "chr18", "19", "chr19", "20", "chr20", "21", "chr21", "x", "chrx", "y", "chry", "mt", "chrmt"}

// Global sorting function
// Note: mergeSort() is missing from this list as it
// is only intended for internal use
func (bf *Bedfile) Sort() error {
	switch bf.SortType {
	case "lex":
		bf.Lines = lexicographicSort(bf.Lines)
	case "nat":
		bf.Lines = naturalSort(bf.Lines)
	case "hum":
		bf.Lines = humanSort(bf.Lines)
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
// Sorting order: chr, start, stop, strand, feat
// Chr hierarchy: 1 < 2 < 10 < MT < X
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

// Human chr sorting
// Sorting order: chr, start, stop, strand, feat
// Chr hierarchy: 1 < 2 < 10 < X < MT
func humanSort(lines []Line) []Line {
	slices.SortStableFunc(lines, func(a, b Line) int {
		return cmp.Or(
			stringSliceCompare(a.Chr, b.Chr, humanChrOrder),
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

// Compares string using a predefined order contained in a slice.
//
// Strings that are not in the slice with be ranked as greater
// than the strings in the slice. If neither a or b are in the slice
// they will be compared using naturalStringCompare.
//
//	-1 if a is less than b
//	 0 if a equals b
//	+1 if a is greater than b
func stringSliceCompare(a, b string, order []string) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	aIdx := idxInSlice(order, a)
	bIdx := idxInSlice(order, b)

	if aIdx != -1 && bIdx != -1 {
		return cmp.Compare(aIdx, bIdx)
	}
	if aIdx != -1 && bIdx == -1 {
		return -1
	}
	if bIdx != -1 && aIdx == -1 {
		return 1
	}
	return naturalStringCompare(a, b)
}

// Returns position if item is in slice and
// -1 if item is not in slice
func idxInSlice(slice []string, item string) int {
	for j, i := range slice {
		if item == i {
			return j
		}
	}
	return -1
}
