package bed

import "strconv"

// Pad regions
func (bf *Bedfile) PadLines() {
	var paddedLines []Line
	for _, line := range bf.Lines {
		line.Start = line.Start - bf.Padding
		line.Stop = line.Stop + bf.Padding
		// Make sure that the padding does not exceed the chromosome limits
		// if the chromosome exists in the fasta index file
		chrLength, ok := bf.chrLengthMap[line.Chr]
		if line.Start < 1 {
			line.Start = 1
		}
		if ok && line.Stop > chrLength {
			line.Stop = chrLength
		}
		line.Full[startIdx] = strconv.Itoa(line.Start)
		line.Full[stopIdx] = strconv.Itoa(line.Stop)
		paddedLines = append(paddedLines, line)
	}
	bf.Lines = paddedLines
}
