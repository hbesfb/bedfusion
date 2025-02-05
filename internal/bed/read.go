package bed

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Bed file constants
const (
	chrIdx   = 0
	startIdx = 1
	stopIdx  = 2
)

// Opening and reading the bed files and optional fasta index file
func (bf *Bedfile) Read() error {
	for _, input := range bf.Inputs {
		bedFile, err := os.Open(input)
		if err != nil {
			return err
		}
		defer bedFile.Close()
		if err := bf.readBed(bedFile); err != nil {
			return fmt.Errorf("can't read bed file %s: %q", input, err)
		}
	}
	if bf.FastaIdx != "" {
		fastaIdxFile, err := os.Open(bf.FastaIdx)
		if err != nil {
			return err
		}
		defer fastaIdxFile.Close()
		if err := bf.readFastaIdx(fastaIdxFile); err != nil {
			return fmt.Errorf("can't read fasta index file %s: %q", bf.FastaIdx, err)
		}
	}
	return nil
}

// Reading the bed file
func (bf *Bedfile) readBed(file io.Reader) error {
	var err error
	var expectedNrOfCols int

	minNrCols := 3
	headerPattern := regexp.MustCompile(`^(browser|track|#)`)
	strandPattern := regexp.MustCompile(`^(\.|\+|-|\+1|-1|1)$`)

	// If there is already content in bf save the expectedNrOfCols
	if len(bf.Lines) != 0 {
		expectedNrOfCols = len(bf.Lines[0].Full)
	}

	lineNr := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l Line
		lineNr++

		lineText := scanner.Text()

		// Handle headers
		if headerPattern.MatchString(lineText) && len(bf.Lines) == 0 {
			bf.Header = append(bf.Header, lineText)
			continue
		}

		// Split line
		l.Full = strings.Split(lineText, "\t")

		// For the first non-header line save the number of columns
		if lineNr == len(bf.Header)+1 && expectedNrOfCols == 0 {
			expectedNrOfCols = len(l.Full)
			if expectedNrOfCols < minNrCols {
				return fmt.Errorf("less than %d columns on line %d: %s", minNrCols, lineNr, lineText)
			}
		}
		if len(l.Full) != expectedNrOfCols {
			return fmt.Errorf("expected %d columns on line %d got %d: %s",
				expectedNrOfCols, lineNr, len(l.Full), lineText)
		}

		// Fill struct
		l.Chr = l.Full[chrIdx]
		l.Start, err = strconv.Atoi(l.Full[startIdx])
		if err != nil {
			return fmt.Errorf("non-int start position on line %d: %s", lineNr, l.Full[startIdx])
		}
		l.Stop, err = strconv.Atoi(l.Full[stopIdx])
		if err != nil {
			return fmt.Errorf("non-int stop position on line %d: %s", lineNr, l.Full[stopIdx])
		}
		// Verify start and stop
		if l.Start > l.Stop {
			return fmt.Errorf("stop is greater than start on line %d: %d > %d\n", lineNr, l.Start, l.Stop)
		}
		if l.Start == l.Stop {
			fmt.Fprintf(os.Stderr, "warning: start and stop is equal on line %d: %d == %d\n", lineNr, l.Start, l.Stop)
		}
		// Set strand and feature if selected
		if bf.StrandCol > stopIdx {
			if bf.StrandCol > len(l.Full)-1 {
				return fmt.Errorf("given strand column, %d, is outside bed file (nr columns=%d)", bf.StrandCol+1, len(l.Full))
			}
			l.Strand = l.Full[bf.StrandCol]
			// Verify strand format
			if !strandPattern.MatchString(l.Strand) {
				return fmt.Errorf("unexpected strand format on line %d: %s", lineNr, l.Strand)
			}
		}
		if bf.FeatCol > stopIdx {
			if bf.FeatCol > len(l.Full)-1 {
				return fmt.Errorf("given strand column, %d, is outside bed file (nr columns=%d)", bf.FeatCol+1, len(l.Full))
			}
			l.Feat = l.Full[bf.FeatCol]
		}
		bf.Lines = append(bf.Lines, l)
	}
	return nil
}

// Reading the fasta index file
func (bf *Bedfile) readFastaIdx(file io.Reader) error {
	var chrOrder []string

	minNrCols := 2
	chrLengthMap := map[string]int{}

	const (
		chrFIdx  = 0
		sizeFIdx = 1
	)

	lineNr := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNr++

		lineText := scanner.Text()

		// Split line
		cols := strings.Split(lineText, "\t")

		// For the first content line set the number of columns if it is empty
		if len(cols) < minNrCols {
			return fmt.Errorf("expected at least %d columns on line %d got %d: %s",
				minNrCols, lineNr, len(cols), lineText)
		}

		// Put chromosome sizes in map and record chromosome order
		size, err := strconv.Atoi(cols[sizeFIdx])
		if err != nil {
			return fmt.Errorf("non-int size for chr %s on line %d: %s", cols[chrFIdx], lineNr, cols[sizeFIdx])
		}
		chrLengthMap[cols[chrFIdx]] = size
		chrOrder = append(chrOrder, cols[chrFIdx])
	}
	// Check that file is not empty
	if lineNr == 0 {
		return fmt.Errorf("fasta index file %s is empty", bf.FastaIdx)
	}

	// Overwrite chr order map if --sorting-type=fidx
	if bf.SortType == FidxST {
		bf.chrOrderMap = chrOrderToMap(chrOrder)
	}
	bf.chrLengthMap = chrLengthMap
	return nil
}
