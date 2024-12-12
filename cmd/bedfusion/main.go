package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"

	kongyaml "github.com/alecthomas/kong-yaml"

	"github.com/hbesfb/bedfusion/internal/bed"
)

type session struct {
	ConfigFile kong.ConfigFlag `env:"CONFIG_FILE" short:"c" help:"The path to configuration file (must be in key-value yaml format)"`
	Bedfile    bed.Bedfile     `embed:""`
	ctx        *kong.Context
}

func main() {
	var s session
	// Getting variables
	s.ctx = kong.Parse(&s,
		kong.Description("Another tool for sorting and merging bed files.\n\n"+
			"BedFusion follows the bed file standard outlined in: https://github.com/samtools/hts-specs/blob/94500cf76f049e898dec7af23097d877fde5894e/BEDv1.pdf \n\n"+
			"Read priority order: 1. flags 2. configuration file 3. environmental variables"),
		kong.Configuration(kongyaml.Loader),
		kong.UsageOnError(),
	)
	// Verify and handle bed file input
	if err := s.Bedfile.VerifyAndHandle(); err != nil {
		fmt.Fprintf(os.Stderr, "error upon verification: %q\n", err)
		s.ctx.Exit(1)
	}
	// Read bed file
	if err := s.Bedfile.Read(); err != nil {
		fmt.Fprintf(os.Stderr, "error while reading: %q\n", err)
		s.ctx.Exit(1)
	}
	// Merge
	if !s.Bedfile.NoMerge {
		s.Bedfile.MergeLines()
	}
	// Sort
	if err := s.Bedfile.Sort(); err != nil {
		fmt.Fprintf(os.Stderr, "error while sorting: %q\n", err)
		s.ctx.Exit(1)
	}
	// Deduplicate if chosen and we have not merged
	// Must be used after sort as it requires the lines to be sorted
	// before use
	if s.Bedfile.Deduplicate && s.Bedfile.NoMerge {
		s.Bedfile.DeduplicateLines()
	}

	// Write output
	if err := s.Bedfile.Write(); err != nil {
		fmt.Fprintf(os.Stderr, "error while writing: %q\n", err)
		s.ctx.Exit(1)
	}
}
