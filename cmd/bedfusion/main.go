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

// Validate bed input
func (s *session) Validate() error {
	if err := s.Bedfile.VerifyAndHandle(); err != nil {
		return err
	}
	// Give warnings about wrong unused variables if a
	// config file is used
	if s.ConfigFile != "" {
		s.Bedfile.WarnAboutWrongUnusedVariables()
	}
	return nil
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
	// Read bed file
	if err := s.Bedfile.Read(); err != nil {
		fmt.Fprintf(os.Stderr, "error while reading: %q\n", err)
		s.ctx.Exit(1)
	}
	// Merge
	if !s.Bedfile.NoMerge && !s.Bedfile.Fission {
		s.Bedfile.MergeLines()
	}
	// Pad
	if s.Bedfile.Padding != 0 {
		s.Bedfile.PadLines()
	}
	// Fission
	if s.Bedfile.Fission {
		s.Bedfile.SplitLines()
	}
	// Deduplicate if chosen and we have not merged
	if s.Bedfile.Deduplicate && s.Bedfile.NoMerge {
		s.Bedfile.DeduplicateLines()
	}
	// Sort
	if err := s.Bedfile.Sort(); err != nil {
		fmt.Fprintf(os.Stderr, "error while sorting: %q\n", err)
		s.ctx.Exit(1)
	}
	// Write output
	if err := s.Bedfile.Write(); err != nil {
		fmt.Fprintf(os.Stderr, "error while writing: %q\n", err)
		s.ctx.Exit(1)
	}
}
