package bed

import (
	"fmt"
	"os"
	"strconv"
)

// Pad regions
func (bf *Bedfile) PadLines() error {
	var chrNotInLengthMap []string
	var err error

	// Loop over and pad lines
	for i, line := range bf.Lines {
		bf.Lines[i], chrNotInLengthMap, err = bf.padAccordingToPaddingType(line, chrNotInLengthMap)
		if err != nil {
			return err
		}
	}
	// Warn depending on padding type
	bf.paddingWarnings(chrNotInLengthMap)
	return nil
}

// Handle missing chromosome in chromosome length map
func (bf Bedfile) padAccordingToPaddingType(line Line, chrNotInLengthMap []string) (Line, []string, error) {
	// Check padding type (just used for internal checks)
	if !stringInSlice([]string{"err", "warn", "force"}, bf.PaddingType) {
		return Line{}, nil, fmt.Errorf("unknown padding type %s", bf.PaddingType)
	}
	// Pad line
	paddedLine, chrInMap := bf.padLine(line)
	if !chrInMap {
		switch bf.PaddingType {
		case "err":
			return Line{}, nil, fmt.Errorf("chromosome %s is not in fasta index file %s", line.Chr, bf.FastaIdx)
		case "warn":
			paddedLine = line
		}
		chrNotInLengthMap = append(chrNotInLengthMap, line.Chr)
	}
	return paddedLine, chrNotInLengthMap, nil
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
			} else {
				fmt.Fprintf(os.Stderr, "warning: you are now padding without a fasta index file and might pad regions beyond chromosome borders\n")
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
