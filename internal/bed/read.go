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

type Bedfile struct {
	Path      string   `env:"INPUT_FILE" required:"" short:"i" help:"The bed file"`
	StrandCol int      `env:"STRAND_COL" help:"The column containing the strand information (first column is 0)"`
	FeatCol   int      `env:"FEAT_COL" help:"The column containing the feature information (first column is 0)"`
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

// TODO: Function for verifying that Feat and Strand cols are not overlapping
// and that they are more than 2

// Opening and reading the
func (bf *Bedfile) Read() error {
	file, err := os.Open(bf.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	return bf.read(file)
}

// Reading the bed file
func (bf *Bedfile) read(file io.Reader) error {
	var err error
	var expectedNrOfCols int

	const (
		chrIdx   = 0
		startIdx = 1
		stopIdx  = 2
	)

	minNrCols := 3
	headerPattern := regexp.MustCompile(`^(browser|track|#)`)
	strandPattern := regexp.MustCompile(`^(\.|\+|-|\+1|-1|1)$`)

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

		// For the first content line save the number of columns
		if lineNr == len(bf.Header)+1 {
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
		// Verify strand and gene column info
		if bf.StrandCol > stopIdx {
			l.Strand = l.Full[bf.StrandCol]
			if !strandPattern.MatchString(l.Strand) {
				return fmt.Errorf("unexpected strand format on line %d: %s", lineNr, l.Strand)
			}
		}
		if bf.FeatCol > stopIdx {
			l.Feat = l.Full[bf.FeatCol]
		}
		bf.Lines = append(bf.Lines, l)
	}
	return nil
}
