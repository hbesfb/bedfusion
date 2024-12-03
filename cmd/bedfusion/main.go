package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"

	kongyaml "github.com/alecthomas/kong-yaml"

	"github.com/hbesfb/bedfusion/internal/bed"
	"github.com/hbesfb/bedfusion/internal/sorting"
)

type session struct {
	ConfigFile kong.ConfigFlag `env:"CONFIG_FILE" short:"c" help:"The path to configuration file (must be in key-value yaml format)"`
	Bedfile    bed.Bedfile     `embed:""`
	Sort       sorting.Config  `embed:""`
	ctx        *kong.Context
}

func main() {
	var s session
	var err error
	// Getting variables
	s.ctx = kong.Parse(&s,
		kong.Description("Another tool for sorting and merging bed files.\n\n"+
			"BedFusion follows the bed file standard outlined in: https://github.com/samtools/hts-specs/blob/94500cf76f049e898dec7af23097d877fde5894e/BEDv1.pdf \n\n"+
			"Read priority order: 1. flags 2. configuration file 3. environmental variables"),
		kong.Configuration(kongyaml.Loader),
		kong.UsageOnError(),
		kong.Vars{},
	)
	// Verify and handle bed file input
	if err := s.Bedfile.VerifyAndHandle(); err != nil {
		fmt.Fprintf(os.Stderr, "error while reading: %q\n", err)
		s.ctx.Exit(1)
	}
	// Read bed file
	if err := s.Bedfile.Read(); err != nil {
		fmt.Fprintf(os.Stderr, "error while reading: %q\n", err)
		s.ctx.Exit(1)
	}
	// TODO: Merge
	// Sort
	s.Bedfile.Lines, err = s.Sort.Sort(s.Bedfile.Lines)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while sorting: %q\n", err)
		s.ctx.Exit(1)
	}
	// Write output
	if err := s.Bedfile.Write(); err != nil {
		fmt.Fprintf(os.Stderr, "error while writing: %q\n", err)
		s.ctx.Exit(1)
	}
}
