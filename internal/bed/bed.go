package bed

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Note that only lowercase is used in this slice
var humanChrOrder = []string{"1", "chr1", "2", "chr2", "3", "chr3", "4", "chr4", "5", "chr5", "6", "chr6", "7", "chr7", "8", "chr8", "9", "chr9", "10", "chr10", "11", "chr11", "12", "chr12", "13", "chr13", "14", "chr14", "15", "chr15", "16", "chr16", "17", "chr17", "18", "chr18", "19", "chr19", "20", "chr20", "21", "chr21", "X", "chrX", "Y", "chrY", "M", "chrM", "MT", "chrMT"}

// Note that the the user will give the columns with 1-based indexing,
// but that we convert this to zero-based indexing in .VerifyAndHandle()
type Bedfile struct {
	Input     string `arg:"" help:"Bed file path"`
	Output    string `env:"OUTPUT_FILE" short:"o" help:"Path to the output file. If unset the output will be written to stdout"`
	StrandCol int    `env:"STRAND_COL" group:"input" help:"The column containing the strand information (1-based column index)"`
	FeatCol   int    `env:"FEAT_COL" group:"input" help:"The column containing feature (e.g. gene) information (1-based column index)"`

	SortType string   `env:"SORT_TYPE" group:"sorting" enum:"lex,nat,ccs" default:"lex" short:"s" help:"How the bed files should be sorted. lex = lexicographic sorting (chr: 1 < 10 < 2 < MT < X), nat = natural sorting (chr: 1 < 2 < 10 < MT < X), ccs = custom chromosome sorting (see --chr-order flag )"`
	ChrOrder []string `env:"CHR_ORDER" group:"sorting" help:"Comma separated custom chromosome order, to be used with custom chromosome sorting (--sort-type=ccs). If none is provided human chromosome order will be used (1-21, X, Y, MT)"`

	Merge   bool `env:"MERGE" group:"merging" short:"m" cmd:"" negatable:"" help:"Merge bed file. Note that touching regions are merged (e.g. if two regions are on the same chr they will be merged if one ends at 5 and the other starts at 6)"`
	Overlap int  `env:"OVERLAP" group:"merging" default:0 help:"Overlap between regions to be merged. This can be a positive or negative number (e.g. if you don't want touching regions to bed merged set overlap to -1)"`

	Header      []string `kong:"-"`
	Lines       []Line   `kong:"-"`
	chrOrderMap map[string]int
}

type Line struct {
	Chr    string
	Start  int
	Stop   int
	Strand string
	Feat   string
	Full   []string
}

// Bed file constants
const (
	chrIdx   = 0
	startIdx = 1
	stopIdx  = 2
)

// Verifies the user input for Bedfile, adds a chrOrderMap, fixes paths
// and subtracts 1 from cols to be able to use zero-based indexing
func (bf *Bedfile) VerifyAndHandle() error {
	if bf.StrandCol != 0 {
		if bf.StrandCol < stopIdx+1 {
			return fmt.Errorf("strand column is at position less than 3: %d", bf.StrandCol)
		}
		if bf.StrandCol == bf.FeatCol {
			return fmt.Errorf("same column for strand and feature: %d == %d", bf.StrandCol, bf.FeatCol)
		}
		bf.StrandCol--
	}
	if bf.FeatCol != 0 {
		if bf.FeatCol < stopIdx+1 {
			return fmt.Errorf("strand column is at position less than 3: %d", bf.FeatCol)
		}
		bf.FeatCol--
	}
	// Creating chromosome order map only if from custom chromosome
	// sorting is chosen
	if bf.SortType == "ccs" {
		if len(bf.ChrOrder) == 0 {
			bf.ChrOrder = humanChrOrder
		}
		bf.chrOrderMap = chrOrderToMap(bf.ChrOrder)
	}
	bf.Input = filepath.Clean(bf.Input)
	if bf.Output != "" {
		bf.Output = filepath.Clean(bf.Output)
	}
	return nil
}

// Convert provided chromosome order to map
func chrOrderToMap(chrOrder []string) map[string]int {
	chrOrderMap := make(map[string]int)
	for idx, chr := range chrOrder {
		chrOrderMap[strings.ToLower(chr)] = idx + 1
	}
	return chrOrderMap
}
