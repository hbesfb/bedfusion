package bed

import (
	"fmt"
	"path/filepath"
)

// Note that the first column is 1 when user sets column locations,
// but that .VerifyAndHandle() will correct this to standard indexes
type Bedfile struct {
	Input     string   `env:"INPUT_FILE" required:"" short:"i" help:"Bed file path"`
	StrandCol int      `env:"STRAND_COL" help:"The column containing the strand information (first column is 1)"`
	FeatCol   int      `env:"FEAT_COL" help:"The column containing the feature information (first column is 1)"`
	Output    string   `env:"OUTPUT_FILE" short:"o" help:"Path for the output file. If unset the output will be written to stdout"`
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
// fixes path and subtracts 1 from cols to be able to use normal indexing
func (bf *Bedfile) VerifyAndHandle() error {
	if bf.StrandCol != 0 {
		if bf.StrandCol < stopIdx+1 && bf.StrandCol != 0 {
			return fmt.Errorf("strand column is less than 3: %d", bf.StrandCol)
		}
		if bf.StrandCol == bf.FeatCol {
			return fmt.Errorf("overlapping strand and feature columns: %d, %d", bf.StrandCol, bf.FeatCol)
		}
		bf.StrandCol--
	}
	if bf.FeatCol != 0 {
		if bf.FeatCol < stopIdx+1 && bf.FeatCol != 0 {
			return fmt.Errorf("strand column is less than 3: %d", bf.FeatCol)
		}
		bf.FeatCol--
	}
	bf.Input = filepath.Clean(bf.Input)
	return nil
}
