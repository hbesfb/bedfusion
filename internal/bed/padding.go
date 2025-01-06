package bed

import "strconv"

// Pad regions
func (bf *Bedfile) PadLines() {
	var paddedLines []Line
	for _, line := range bf.Lines {
		paddedLines = append(paddedLines, bf.PadLine(line))
	}
	bf.Lines = paddedLines
}

// Pad single line
func (bf Bedfile) PadLine(line Line) Line {
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
	return line
}
