package bed

import (
	"fmt"
	"strconv"
	"strings"
)

// Merge and pad lines in bed file
func (bf *Bedfile) MergeAndPadLines() error {
	var merged Line
	var mergedLines []Line
	var chrNotInLengthMap []string
	for i, l := range mergeSort(bf.Lines) {
		// Pad line
		if bf.Padding != 0 {
			var err error
			l, chrNotInLengthMap, err = bf.padAccordingToPaddingType(l, chrNotInLengthMap)
			if err != nil {
				return err
			}
		}

		// Merge lines
		// If the lines are overlapping or touching merge them
		if i != 0 &&
			merged.Chr == l.Chr &&
			merged.Strand == l.Strand &&
			merged.Feat == l.Feat &&
			merged.Stop+bf.Overlap >= l.Start-1 {
			// Set new stop if it is later than the
			// merged stop
			if l.Stop > merged.Stop {
				merged.Stop = l.Stop
				merged.Full[stopIdx] = strconv.Itoa(l.Stop)
			}
			// Join information in the optional columns
			if len(l.Full) > stopIdx+1 {
				for idx, col := range l.Full[stopIdx+1:] {
					mIdx := idx + stopIdx + 1
					if !stringInSlice(strings.Split(merged.Full[mIdx], ","), col) {
						merged.Full[mIdx] = fmt.Sprintf("%s,%s", merged.Full[mIdx], col)
					}
				}
			}
		} else {
			// If we are not on the first line append merged to MergedLines
			if i != 0 {
				mergedLines = append(mergedLines, merged)
			}
			// Create new merged line
			merged = Line{
				Chr: l.Chr, Start: l.Start, Stop: l.Stop,
				Strand: l.Strand, Feat: l.Feat,
				Full: l.Full,
			}
		}
	}
	// If we have been padding print padding warnings
	if bf.Padding != 0 {
		bf.paddingWarnings(chrNotInLengthMap)
	}
	// Replace lines in Bedfile
	bf.Lines = append(mergedLines, merged)
	return nil
}

// Returns true or false depending on if the string
// is in a slice
func stringInSlice(slice []string, item string) bool {
	for _, i := range slice {
		if item == i {
			return true
		}
	}
	return false
}
