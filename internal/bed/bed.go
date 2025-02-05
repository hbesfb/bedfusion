package bed

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Note that the the user will give the columns with 1-based indexing,
// but that we convert this to zero-based indexing in .VerifyAndHandle()
type Bedfile struct {
	Inputs   []string `arg:"" help:"Bed file path(s). If more than one is provided the files will be joined as if they were one file"`
	Output   string   `env:"OUTPUT_FILE" short:"o" help:"Path to the output file. If unset the output will be written to stdout"`
	FastaIdx string   `env:"FASTA_IDX" help:"Tab separated file containing at least two columns where the first column contains the chromosome and the second it's size. Compatible with fasta index files, but any text file can be used as long as the file conditions are met"`

	StrandCol int `env:"STRAND_COL" group:"input" help:"The column containing the strand information (1-based column index). If this option is set regions on the same strand will not be merged"`
	FeatCol   int `env:"FEAT_COL" group:"input" help:"The column containing the feature (e.g. gene id, transcript id etc.) information (1-based column index). If this option is set regions on the same feature will not be merged"`

	SortType    string   `env:"SORT_TYPE" group:"sorting" enum:"${lexST},${natST},${ccsST},${fidxST}" default:"${lexST}" short:"s" help:"How the bed file should be sorted. ${lexST} = lexicographic sorting (chr: 1 < 10 < 2 < MT < X), ${natST} = natural sorting (chr: 1 < 2 < 10 < MT < X), ${ccsST} = custom chromosome sorting (see --chr-order flag ), ${fidxST} = use ordering from fasta index file (must be used together with --fasta-idx)"`
	ChrOrder    []string `env:"CHR_ORDER" group:"sorting" help:"Comma separated custom chromosome order, to be used with custom chromosome sorting (--sort-type=ccs). Chromosomes not on the list will be sorted naturally after the ones in the list"`
	Deduplicate bool     `env:"DEDUPLICATE" group:"sorting" cmd:"" short:"d" help:"Remove duplicated lines"`

	NoMerge bool `env:"NO_MERGE" group:"merging" cmd:"" help:"Do not merge regions"`
	Overlap int  `env:"OVERLAP" group:"merging" default:"0" help:"Overlap between regions to be merged. Note that touching regions are merged (e.g. if two regions are on the same chr, and the overlap is they will be merged if one ends at 5 and the other starts at 6). If you don't want touching regions to be merged set overlap to -1"`

	PaddingType string `env:"PADDING_TYPE" group:"padding" enum:"${failPT},${warnPT},${forcePT}" default:"${failPT}" help:"Padding type. safe = bedfusion will fail if it encounters a chromosome not in the fasta index file, ${warnPT} = will only pad regions in the fasta index file and give a warning about chromosomes not in the fasta index file, ${forcePT} = will pad regardless, if --fasta-idx is set there will be given a warning about the chromosomes not in the fasta index file, if --fasta-idx is not set no warnings will be given"`
	Padding     int    `env:"PADDING" group:"padding" help:"Padding in bp. Note that padding is done before merging"`
	FirstBase   int    `env:"FIRST_BASE" group:"padding" default:"0" help:"The start coordinate of the first base on each chromosome"`

	SplitSize int `env:"SPLIT_SIZE" group:"splitting" help:"Size of split regions in bp. Will be done after merging."`

	Header       []string `kong:"-"`
	Lines        []Line   `kong:"-"`
	chrOrderMap  map[string]int
	chrLengthMap map[string]int
}

type Line struct {
	Chr    string
	Start  int
	Stop   int
	Strand string
	Feat   string
	Full   []string
}

// Verifies and handles Bedfile input
func (bf *Bedfile) VerifyAndHandle() error {
	if err := bf.verifyAndHandleColumns(); err != nil {
		return err
	}
	if err := bf.verifyFastaIdxCombinations(); err != nil {
		return err
	}
	if err := bf.verifySplitSizeInput(); err != nil {
		return err
	}
	if err := bf.verifyFirstBase(); err != nil {
		return err
	}
	bf.handleCCSSorting()
	bf.cleanPaths()
	return nil
}

// Verifies Strand and Feat columns and subtracts 1 to be able to use zero-based indexing
func (bf *Bedfile) verifyAndHandleColumns() error {
	if bf.StrandCol != 0 {
		if bf.StrandCol < stopIdx+1 {
			return fmt.Errorf("--strand-col is at position less than 3: %d", bf.StrandCol)
		}
		if bf.StrandCol == bf.FeatCol {
			return fmt.Errorf("--strand-col and --feat-col can not be set to the same column: %d == %d", bf.StrandCol, bf.FeatCol)
		}
		bf.StrandCol--
	}
	if bf.FeatCol != 0 {
		if bf.FeatCol < stopIdx+1 {
			return fmt.Errorf("--feat-col is at position less than 3: %d", bf.FeatCol)
		}
		bf.FeatCol--
	}
	return nil
}

// Verify fasta-idx combinations
func (bf Bedfile) verifyFastaIdxCombinations() error {
	// Verify that fasta-idx is set if padding is selected
	if bf.Padding != 0 && bf.PaddingType != "force" && bf.FastaIdx == "" {
		return fmt.Errorf("--padding-type=%s must be used together with --fasta-idx", bf.PaddingType)
	}
	// Verify that fasta-idx is set if sort type is fastaidx
	if bf.SortType == FidxST && bf.FastaIdx == "" {
		return fmt.Errorf("--sort-type=%s must be used together with --fasta-idx", bf.SortType)
	}
	return nil
}

// Verify split size input
func (bf Bedfile) verifySplitSizeInput() error {
	if bf.SplitSize < 0 {
		return fmt.Errorf("--split-size must be > 0: %d", bf.SplitSize)
	}
	return nil
}

// Verify first base input
func (bf *Bedfile) verifyFirstBase() error {
	if bf.FirstBase < 0 || bf.FirstBase > 1 {
		return fmt.Errorf("--first-base must be either 0 or 1: %d", bf.FirstBase)
	}
	return nil
}

// Create chr order map
func (bf *Bedfile) handleCCSSorting() {
	// Creating chromosome order map only if from custom chromosome
	// sorting is chosen
	if bf.SortType == CcsST {
		if len(bf.ChrOrder) == 0 {
			bf.ChrOrder = humanChrOrder
		}
		bf.chrOrderMap = chrOrderToMap(bf.ChrOrder)
	}
}

// Convert provided chromosome order to map
func chrOrderToMap(chrOrder []string) map[string]int {
	chrOrderMap := make(map[string]int)
	for idx, chr := range chrOrder {
		chrOrderMap[strings.ToLower(chr)] = idx + 1
	}
	return chrOrderMap
}

// Clean paths
func (bf *Bedfile) cleanPaths() {
	for i, input := range bf.Inputs {
		bf.Inputs[i] = filepath.Clean(input)
	}
	if bf.Output != "" {
		bf.Output = filepath.Clean(bf.Output)
	}
	if bf.FastaIdx != "" {
		bf.FastaIdx = filepath.Clean(bf.FastaIdx)
	}
}
