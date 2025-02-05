package bed

import (
	"strconv"
)

// Split lines based on set SplitSize
func (bf *Bedfile) SplitLines() {
	var splitLine Line
	var splitLines []Line
	for _, line := range bf.Lines {
		start := line.Start
		for line.Stop-start > bf.SplitSize {
			stop := start + bf.SplitSize
			splitLine = line.replaceStartAndStop(start, stop)
			splitLines = append(splitLines, splitLine)
			start = stop + 1
		}
		if start != bf.SplitSize {
			splitLine = line.replaceStartAndStop(start, line.Stop)
		}
		splitLines = append(splitLines, splitLine)
	}
	bf.Lines = splitLines
}

// Replace the start and stop position of a line
func (l Line) replaceStartAndStop(start, stop int) Line {
	newLine := Line{
		Chr: l.Chr, Start: start, Stop: stop,
		Strand: l.Strand, Feat: l.Feat,
	}
	newLine.Full = append([]string{l.Chr, strconv.Itoa(start), strconv.Itoa(stop)}, l.Full[stopIdx+1:]...)
	return newLine
}
