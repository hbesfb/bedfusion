package bed

import (
	"fmt"
	"path/filepath"
)

// Note that the the user will give the columns with 1-based indexing,
// but that we convert this to zero-based indexing in .VerifyAndHandle()
type Bedfile struct {
	Input     string   `arg:"" help:"Bed file path"`
	Output    string   `env:"OUTPUT_FILE" short:"o" help:"Path to the output file. If unset the output will be written to stdout"`
	StrandCol int      `env:"STRAND_COL" help:"The column containing the strand information (1-based column index)"`
	FeatCol   int      `env:"FEAT_COL" help:"The column containing the feature information (1-based column index)"`
	Header    []string `kong:"-"`
	Lines     []Line   `kong:"-"`
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

// Verifies the user input for Bedfile
// fixes path and subtracts 1 from cols to be able to use zero-based indexing
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
	bf.Input = filepath.Clean(bf.Input)
	if bf.Output != "" {
		bf.Output = filepath.Clean(bf.Output)
	}
	return nil
}
