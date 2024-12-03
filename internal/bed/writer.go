package bed

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Writing bed file or standard output
func (bf *Bedfile) Write() error {
	// If output is not set write to Stdout
	if bf.Output == "" {
		return bf.write(os.Stdout)
	}

	// If output is set write to file
	file, err := os.Create(bf.Output)
	if err != nil {
		return fmt.Errorf("cannot create output file: %v", err)
	}
	defer file.Close()
	return bf.write(file)
}

// Write bedfile content as string to writer destination
func (bf *Bedfile) write(writer io.Writer) error {
	reader := strings.NewReader(bf.toString())
	_, err := io.Copy(writer, reader)
	return err
}

// Transform bed file headers and lines into a string for writing
// note that it will use the full lines
func (bf *Bedfile) toString() string {
	var bedAsString string
	// Add header if available
	if len(bf.Header) > 0 {
		bedAsString = fmt.Sprintf("%s\n", strings.Join(bf.Header, "\n"))
	}
	// Add lines
	for _, l := range bf.Lines {
		bedAsString = fmt.Sprintf("%s%s\n", bedAsString, strings.Join(l.Full, "\t"))
	}
	return bedAsString
}
