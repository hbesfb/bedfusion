package bed

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

// Pad regions
func (bf *Bedfile) PadLines() error {
	var paddedLines []Line
	var chrNotInLengthMap []string
	var err error

	// Check padding type (just used for internal checks)
	if !slices.Contains([]string{"err", "warn", "force"}, bf.PaddingType) {
		return fmt.Errorf("unknown padding type %s", bf.PaddingType)
	}
	// Loop over and pad lines
	for _, line := range bf.Lines {
		paddedLines, chrNotInLengthMap, err = bf.padAccordingToPaddingType(line, paddedLines, chrNotInLengthMap)
		if err != nil {
			return err
		}
	}
	// Warn depending on padding type
	bf.paddingWarnings(chrNotInLengthMap)
	bf.Lines = paddedLines
	return nil
}

// Handle missing chromosome in chromosome length map
func (bf Bedfile) padAccordingToPaddingType(line Line, paddedLines []Line, chrNotInLengthMap []string) ([]Line, []string, error) {
	paddedLine, chrInMap := bf.padLine(line)
	if chrInMap {
		paddedLines = append(paddedLines, paddedLine)
	} else {
		switch bf.PaddingType {
		case "err":
			return nil, nil, fmt.Errorf("chromosome %s is not in fasta index file %s", line.Chr, bf.FastaIdx)
		case "warn":
			paddedLines = append(paddedLines, line)
		case "force":
			paddedLines = append(paddedLines, paddedLine)
		}
		chrNotInLengthMap = append(chrNotInLengthMap, line.Chr)
	}
	return paddedLines, chrNotInLengthMap, nil
}

// Handle warnings depending on padding types
func (bf Bedfile) paddingWarnings(chrNotInLengthMap []string) {
	if len(chrNotInLengthMap) > 0 {
		warnMsg := fmt.Sprintf("chromosomes %v not in fasta index file %s",
			sortAndDeduplicateListOfStrings(chrNotInLengthMap), bf.FastaIdx)
		switch bf.PaddingType {
		case "warn":
			fmt.Fprintf(os.Stderr, "warning: %s, no padding was added to regions on these chromosomes\n", warnMsg)
		case "force":
			if bf.FastaIdx != "" {
				fmt.Fprintf(os.Stderr, "warning: %s, regions on these chromosomes were still padded\n", warnMsg)
			}
		}
	}
}

// Pad single line
func (bf Bedfile) padLine(l Line) (Line, bool) {
	// Deep line to make sure we do not overwrite
	fullLineCopy := make([]string, len(l.Full))
	_ = copy(fullLineCopy, l.Full)
	line := Line{
		Chr: l.Chr, Start: l.Start, Stop: l.Stop,
		Strand: l.Strand, Feat: l.Feat,
		Full: fullLineCopy,
	}
	// Line
	line.Start = line.Start - bf.Padding
	line.Stop = line.Stop + bf.Padding
	// Make sure that the padding does not exceed the chromosome limits
	chrLength, ok := bf.chrLengthMap[line.Chr]
	if line.Start < 1 {
		line.Start = 1
	}
	if ok && line.Stop > chrLength {
		line.Stop = chrLength
	}
	line.Full[startIdx] = strconv.Itoa(line.Start)
	line.Full[stopIdx] = strconv.Itoa(line.Stop)
	return line, ok
}
