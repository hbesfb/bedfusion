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

// Opening and reading the bed files
func (bf *Bedfile) Read() error {
	for _, input := range bf.Inputs {
		file, err := os.Open(input)
		if err != nil {
			return err
		}
		defer file.Close()
		if err := bf.read(file); err != nil {
			return fmt.Errorf("can't read file %s: %q", input, err)
		}
	}
	return nil
}

// Reading the bed file
func (bf *Bedfile) read(file io.Reader) error {
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

		// For the first content line set the number of columns if it is empty
		if lineNr == len(bf.Header)+1 && expectedNrOfCols == 0 {
			expectedNrOfCols = len(l.Full)
			if expectedNrOfCols < minNrCols {
				return fmt.Errorf("less than 3 columns on line %d: %s", lineNr, lineText)
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
